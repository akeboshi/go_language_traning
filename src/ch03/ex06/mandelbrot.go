//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, superSampling(img))
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			r := 255 - contrast*n
			g := 255 - r
			b := 0
			return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
		}
	}
	return color.RGBA{0, 0, 255, 255}
}

// FIXME: 何が正しいのか不明
func superSampling(source image.Image) image.Image {
	img := image.NewRGBA(source.Bounds())
	for py := 0; py < img.Rect.Dy(); py++ {
		for px := 0; px < img.Rect.Dx(); px++ {
			img.Set(px, py, average(source, px, py))
		}
	}
	return img
}

func average(source image.Image, px, py int) color.Color {
	var colors [5]color.Color
	colors[0] = source.At(px-1, py-1)
	colors[1] = source.At(px-1, py)
	colors[2] = source.At(px, py)
	colors[3] = source.At(px+1, py)
	colors[4] = source.At(px+1, py+1)

	var resultR, resultG, resultB uint32
	for _, color := range colors {
		r, g, b, _ := color.RGBA()
		resultR += r
		resultG += g
		resultB += b
	}
	resultR /= 5
	resultG /= 5
	resultB /= 5
	return color.RGBA{uint8(resultR), uint8(resultG), uint8(resultB), 255}
}
