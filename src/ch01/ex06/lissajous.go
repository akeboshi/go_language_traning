//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{color.Gray{123}}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var i uint8
	for i = 0x00; i < 0xff; i++ {
		palette = append(palette, color.RGBA{0x00, i + 0x01, 0x00, 0xff})
	}
	lissajouse(os.Stdout)
}

func lissajouse(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	isColorIncrease := true
	var colorIndex uint8 = 1
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
			if isColorIncrease {
				colorIndex++
				if colorIndex == 255 {
					isColorIncrease = false
				}
			} else {
				colorIndex--
				if colorIndex == 1 {
					isColorIncrease = true
				}
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
