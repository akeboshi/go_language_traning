package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		x := getFloat64Query(r.FormValue("x"), 0.0)
		y := getFloat64Query(r.FormValue("y"), 0.0)
		scale := getFloat64Query(r.FormValue("scale"), 2.0)
		png.Encode(w, render(x, y, scale))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getFloat64Query(word string, defaultVal float64) float64 {
	if word == "" {
		return defaultVal
	}
	val, err := strconv.ParseFloat(word, 64)
	if err != nil {
		return defaultVal
	}
	return val
}

func render(deltaX, deltaY, scale float64) image.Image {
	xmin, ymin, xmax, ymax := -scale, -scale, scale, scale
	const (
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x+deltaX, y-deltaY)
			img.Set(px, py, newton(z))
		}
	}
	return img
}

func newton(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	for n := uint(0); n < iterations; n++ {
		z = z - ((z*z*z*z - 1) / (4 * z * z * z))
		if cmplx.Abs(z*z*z*z-1) < 1e-3 {
			r := 255 - contrast*n
			g := 255 - r
			b := 0
			return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		}
	}
	return color.RGBA{0, 0, 255, 255}
}
