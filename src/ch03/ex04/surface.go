package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

var width, height float64
var xyscale float64
var zscale float64
var topColor, downColor string

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		width = getFloat64Query(r.FormValue("width"), 600)
		height = getFloat64Query(r.FormValue("height"), 320)
		topColor = getColorQuery(r.FormValue("top"), "r")
		downColor = getColorQuery(r.FormValue("down"), "g")
		xyscale = width / 2 / xyrange
		zscale = height * 0.4

		fmt.Fprint(w, svg())
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func getColorQuery(word string, defaultVal string) string {
	if word == "r" || word == "g" || word == "b" {
		return word
	}
	return defaultVal
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

func svg() string {
	var functions []func(float64, float64) float64
	functions = append(functions, hypot)
	min, max := getMinAndMaxHeight(functions)
	result := ""

	result += fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg'"+
		" style='stroke: grey; stroke-width: 0.7'"+
		" width='%d' height ='%d'>\n", int(width), int(height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, bx, by, cx, cy, dx, dy := getPolygonPoints(i, j, functions)
			if valid(ax, ay, bx, by, cx, cy, dx, dy) {
				result += fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' style='fill: %s' />\n",
					ax, ay, bx, by, cx, cy, dx, dy, getColor(i, j, min, max, functions))
			}
		}
	}

	result += "</svg>\n"
	return result
}

func getColor(i, j int, min, max float64, functions []func(float64, float64) float64) string {
	z := getHeight(i, j, functions)
	colorVal := int(((z - min) / (max - min)) * 255)
	r := setAssignedColor(colorVal, "r")
	g := setAssignedColor(colorVal, "g")
	b := setAssignedColor(colorVal, "b")
	return fmt.Sprintf("#%02x%02x%02x", int(r), int(g), int(b))
}

func setAssignedColor(colorVal int, assignedColor string) int {
	if topColor == assignedColor {
		return colorVal
	} else if downColor == assignedColor {
		return 255 - colorVal
	} else {
		return 0
	}
}

func getPolygonPoints(i, j int, functions []func(float64, float64) float64) (float64, float64, float64, float64, float64, float64, float64, float64) {
	ax, ay := corner(i+1, j, functions)
	bx, by := corner(i, j, functions)
	cx, cy := corner(i, j+1, functions)
	dx, dy := corner(i+1, j+1, functions)
	return ax, ay, bx, by, cx, cy, dx, dy
}

func getMinAndMaxHeight(functions []func(float64, float64) float64) (float64, float64) {
	min := math.Inf(1)
	max := math.Inf(-1)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			z := getHeight(i, j, functions)
			if z < min {
				min = z
			}
			if z > max {
				max = z
			}
		}
	}
	return min, max
}

func getHeight(i, j int, functions []func(float64, float64) float64) float64 {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := applyMethods(x, y, functions)
	return z
}

func corner(i, j int, functions []func(float64, float64) float64) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := applyMethods(x, y, functions)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func valid(val ...float64) bool {
	for _, v := range val {
		if math.IsInf(v, 0) || math.IsNaN(v) {
			return false
		}
	}
	return true
}

func applyMethods(x, y float64, functions []func(float64, float64) float64) float64 {
	var sums float64
	for _, f := range functions {
		sums += f(x, y)
	}
	functionLen := len(functions)
	if functionLen == 0 {
		functionLen = 1
	}
	return sums / float64(functionLen)
}

func mogle(x, y float64) float64 {
	return math.Sin(-x) * math.Pow(1.5, -math.Hypot(x, y))
}

func egg(x, y float64) float64 {
	return math.Pow(2, math.Sin(y)) * math.Pow(2, math.Sin(x)) / 12
}

func saddle(x, y float64) float64 {
	return math.Sin(x*y/10) / 10
}

func hypot(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
