// Surace COMPUTES AN svg rendering of a 3-D surface function
package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

const ( //实际上这个地方是declare 但是写成了assign
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 //number of grid cells
	xyrange       = 30.0                //axis ranges (-xrange .. +xrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4
	angle         = math.Pi / 6 // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type zFunc func(x, y float64) float64

func svg(w io.Writer, f zFunc) {
	fmt.Printf("svg xmlns = 'http://www.w3.org/2000/svg' "+"style=' stroke: grey; fill: white; stroke-width: 0.7' "+
		"width = '%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			// Add here to do the problem control.
			if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
				//fmt.Errorf("svg: corner() return a NaN value, loop %d - %d\n", i, j)
				continue
			}
			fmt.Fprintf(w, "<polygon style='stroke: %s; fill: #222222' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				"#666666", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int, f zFunc) (float64, float64) {
	//Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx, sy)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) //Hypot is sqrt(p*p + q*q) function which takes care of overflow and underflow
	return math.Sin(r)    //r
}

func eggbox(x, y float64) float64 {
	// f(x,y) = x**2 + y**2 + 25(sin^2(x) + sin^2(y))
	return (math.Pow(x, 2.0) + math.Pow(y, 2) + 25*(math.Pow(math.Sin(x), 2)+math.Pow(math.Sin(y), 2)))
}

func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}

func main() {
	usage := "usage: ex3.2 saddle|eggbox"
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	var f zFunc
	switch os.Args[1] {
	case "saddle":
		f = saddle
	case "eggbox":
		f = eggbox
	default:
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	svg(os.Stdout, f)
}
