// generated by go generate; DO NOT EDIT.

package main

// This data is derived from files in the font/fixed directory of the Plan 9
// Port source code (https://github.com/9fans/plan9port) which were originally
// based on the public domain X11 misc-fixed font files.

import "image"

// mask5x8 contains 96 4×8 glyphs in 3072 Pix bytes.
var mask5x8 = &image.Alpha{
	Stride: 4,
	Rect:   image.Rectangle{Max: image.Point{4, 96 * 8}},
	Pix: []byte{
		// 0x20 ' '
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x21 '!'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x22 '"'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x23 '#'
		0x00, 0xff, 0x00, 0xff,
		0x00, 0xff, 0x00, 0xff,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0xff, 0x00, 0xff,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x24 '$'
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0xff, 0xff,
		0xff, 0x00, 0xff, 0x00,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x25 '%'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x26 '&'
		0x00, 0xff, 0x00, 0x00,
		0xff, 0x00, 0xff, 0x00,
		0xff, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0xff, 0x00, 0xff, 0x00,
		0xff, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x27 '\''
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x28 '('
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x29 ')'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x2a '*'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x2b '+'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x2c ','
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0x00,

		// 0x2d '-'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x2e '.'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0xff, 0x00,

		// 0x2f '/'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x30 '0'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x31 '1'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x32 '2'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x33 '3'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x34 '4'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0xff, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x35 '5'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x36 '6'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x37 '7'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x38 '8'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x39 '9'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x3a ':'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x3b ';'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0xff,
		0x00, 0x00, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0x00,

		// 0x3c '<'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x3d '='
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x3e '>'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x3f '?'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x40 '@'
		0x00, 0x00, 0xff, 0xff,
		0x00, 0xff, 0x00, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0xff, 0x00,
		0xff, 0x00, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0xff, 0xff,

		// 0x41 'A'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0xff, 0xff, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x42 'B'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x43 'C'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x44 'D'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x45 'E'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x46 'F'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x47 'G'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0xff, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x48 'H'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0xff, 0xff, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x49 'I'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x4a 'J'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0xff, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x4b 'K'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0xff, 0x00,
		0xff, 0xff, 0x00, 0x00,
		0xff, 0x00, 0xff, 0x00,
		0xff, 0x00, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x4c 'L'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x4d 'M'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x4e 'N'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0xff, 0x00, 0xff,
		0xff, 0xff, 0xff, 0xff,
		0xff, 0x00, 0xff, 0xff,
		0xff, 0x00, 0xff, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x4f 'O'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x50 'P'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x51 'Q'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0xff, 0x00, 0xff,
		0xff, 0x00, 0xff, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0xff,

		// 0x52 'R'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x53 'S'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x54 'T'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x55 'U'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x56 'V'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x57 'W'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0xff, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x58 'X'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x59 'Y'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x5a 'Z'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x5b '['
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x5c '\\'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x5d ']'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x5e '^'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x5f '_'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff,

		// 0x60 '`'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x61 'a'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x62 'b'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x63 'c'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0xff,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x64 'd'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x65 'e'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0xff, 0xff,
		0xff, 0xff, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x66 'f'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0xff, 0x00, 0x00,
		0xff, 0xff, 0xff, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x67 'g'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,

		// 0x68 'h'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x69 'i'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x6a 'j'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0x00, 0xff, 0x00,

		// 0x6b 'k'
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x6c 'l'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x6d 'm'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0x00, 0xff,
		0xff, 0x00, 0xff, 0x00,
		0xff, 0x00, 0xff, 0x00,
		0xff, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x6e 'n'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x6f 'o'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x70 'p'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,

		// 0x71 'q'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0xff, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0xff,

		// 0x72 'r'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0xff, 0x00,
		0xff, 0xff, 0x00, 0xff,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x73 's'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x74 't'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0xff, 0xff, 0xff, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x75 'u'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x76 'v'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x77 'w'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0x00,
		0xff, 0x00, 0xff, 0x00,
		0xff, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x78 'x'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,
		0x00, 0xff, 0xff, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x79 'y'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0xff, 0x00, 0x00, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0xff,
		0xff, 0x00, 0x00, 0xff,
		0x00, 0xff, 0xff, 0x00,

		// 0x7a 'z'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x7b '{'
		0x00, 0x00, 0xff, 0xff,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0xff, 0xff, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0xff, 0xff,
		0x00, 0x00, 0x00, 0x00,

		// 0x7c '|'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x7d '}'
		0xff, 0xff, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0xff, 0xff,
		0x00, 0xff, 0x00, 0x00,
		0x00, 0x00, 0xff, 0x00,
		0xff, 0xff, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// 0x7e '~'
		0x00, 0x00, 0x00, 0x00,
		0x00, 0xff, 0x00, 0xff,
		0xff, 0x00, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,

		// U+FFFD REPLACEMENT CHARACTER
		0x00, 0xff, 0xff, 0xff,
		0xff, 0xff, 0x00, 0xff,
		0xff, 0x00, 0xff, 0x00,
		0xff, 0xff, 0xff, 0x00,
		0xff, 0xff, 0x00, 0xff,
		0xff, 0xff, 0xff, 0xff,
		0xff, 0xff, 0x00, 0xff,
		0x00, 0xff, 0xff, 0xff,
	},
}
