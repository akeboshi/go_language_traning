//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var mu sync.Mutex
var count int
var palette = []color.Color{color.Gray{123}}

func main() {
	initPallet()
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	var cycles = 5
	strCycles := r.FormValue("cycles")
	i, err := strconv.Atoi(strCycles)
	if err == nil && i > 0 {
		cycles = i
	}
	lissajouse(w, cycles)
}

func initPallet() {
	rand.Seed(time.Now().UTC().UnixNano())
	var i uint8
	for i = 0x00; i < 0xff; i++ {
		palette = append(palette, color.RGBA{0x00, i + 0x01, 0x00, 0xff})
	}
}

func lissajouse(out io.Writer, cycles int) {
	const (
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
		for t := 0.0; t < math.Pi*float64(cycles); t += res {
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

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
