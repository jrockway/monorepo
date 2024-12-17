package pctx

import (
	"time"

	"github.com/jrockway/monorepo/internal/log"
	"github.com/jrockway/monorepo/internal/meters"
)

func ExampleChild() {
	log.InitLogger()
	log.SetLevel(log.DebugLevel)
	ctx := Background("")
	meters.Inc(ctx, "counter", 1)
	meters.Set(ctx, "gauge", 42)
	meters.Sample(ctx, "sampler", "hi")

	ctx, c := WithCancel(ctx)
	ctx = Child(ctx, "aggregated", WithCounter("counter", 0, meters.WithFlushInterval(time.Second)))
	ctx = Child(ctx, "", WithGauge("gauge", 0, meters.WithFlushInterval(time.Second)))
	for i := 0; i < 1<<24; i++ {
		meters.Inc(ctx, "counter", 1)
		meters.Set(ctx, "gauge", i)
	}
	c()
	// Output:
}
