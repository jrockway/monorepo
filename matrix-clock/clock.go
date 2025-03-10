package main

import (
	"context"
	"image"
	"image/color"
	"time"

	"github.com/goiot/devices/dotstar"
	"github.com/jrockway/monorepo/internal/log"
	"go.uber.org/zap"
	"golang.org/x/exp/io/spi"
	"golang.org/x/exp/io/spi/driver"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"golang.org/x/net/trace"
)

// n is the number of LEDs on the strip.
const (
	n = 8 * 8 * 6
)

var myFont = &basicfont.Face{
	Advance: 5,
	Width:   5,
	Height:  8,
	Ascent:  8,
	Descent: 0,
	Mask:    mask5x8,
	Ranges: []basicfont.Range{
		{Low: '\u0020', High: '\u007f', Offset: 0},
		{Low: '\ufffd', High: '\ufffe', Offset: 95},
	},
}

type fakeSPI struct{}

func (x fakeSPI) Open() (driver.Conn, error) { return x, nil }
func (fakeSPI) Configure(k, v int) error     { return nil }
func (fakeSPI) Tx(w, r []byte) error         { return nil }
func (fakeSPI) Close() error                 { return nil }

func getClockImage(t time.Time) *image.RGBA {
	now := t.Format("15:04:05")
	img := image.NewRGBA(image.Rect(0, 0, 48, 8))
	(&font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.RGBA{R: 0x20, G: 0xa0, B: 0xff, A: 0xff}),
		Face: myFont,
		Dot:  fixed.Point26_6{X: fixed.Int26_6(0), Y: fixed.Int26_6(540)},
	}).DrawString(now)
	return img
}

func drawClock(ctx context.Context) {
	dev := "/dev/spidev0.0"
	l := trace.NewEventLog("peripheral", "display")
	l.Printf("open " + dev)
	d, err := dotstar.Open(&spi.Devfs{Dev: dev, Mode: spi.Mode3, MaxSpeed: 100}, n)
	if err != nil {
		l.Errorf("open dotstar: %v", err)
		log.Error(ctx, "failed to open dotstar matrix; using dummy SPI driver", zap.Error(err))
		d, _ = dotstar.Open(fakeSPI{}, n)
	}

	// Blank the display.
	for i := 0; i < 6*8*8; i++ {
		d.SetRGBA(i, dotstar.RGBA{R: 0, G: 0, B: 0, A: 0})
	}
	if err := d.Draw(); err != nil {
		l.Errorf("blank display: %v", err)
	}

	log.Info(ctx, "starting clock update loop")
	for {
		// Render the current time.
		img := getClockImage(time.Now())
		for _, matrix := range []int{0, 2, 4} {
			i := matrix * 64
			for x := matrix * 8; x < (matrix+1)*8; x++ {
				for y := 0; y < 8; y++ {
					r, g, b, _ := img.At(x, y).RGBA()
					scale := byte(4)
					if matrix > 3 {
						scale = 50
					}
					d.SetRGBA(i, dotstar.RGBA{R: byte(r) / scale, G: byte(g) / scale, B: byte(b) / scale, A: 5})
					i++
				}
			}
		}
		for _, matrix := range []int{1, 3, 5} {
			i := matrix * 64
			for x := (matrix+1)*8 - 1; x >= matrix*8; x-- {
				for y := 7; y >= 0; y-- {
					r, g, b, _ := img.At(x, y).RGBA()
					scale := byte(4)
					if matrix > 3 {
						scale = 50
					}
					d.SetRGBA(i, dotstar.RGBA{R: byte(r) / scale, G: byte(g) / scale, B: byte(b) / scale, A: 5})
					i++
				}
			}
		}
		if err := d.Draw(); err != nil {
			l.Errorf("draw clock: %v", err)
		}
		UpdateStatus(Status{ClockFace: img})
		l.Printf("sleeping for %s", time.Until(time.Now().Add(time.Second).Truncate(time.Second)).String())
		time.Sleep(time.Until(time.Now().Add(time.Second).Truncate(time.Second)))
	}
}
