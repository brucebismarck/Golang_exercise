// mandelbrot emits a PNG image of the Mandelbrot fractal
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
		epsX                   = (xmax - xmin) / width
		epsY                   = (ymax - ymin) / height
	)

	offX := []float64{-epsX, epsX}
	offY := []float64{-epsY, epsY}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			// create subpixel
			subPixels := make([]color.Color, 0) //length is 0
			for i := 0; i < 2; i++ {            // split to four so 2*2
				for j := 0; j < 2; j++ {
					z := complex(x+offX[i], y+offY[j])
					subPixels = append(subPixels, mandelbrot(z))
				}
			}
			//Image point (px,py) represents complex value z.
			img.Set(px, py, avg(subPixels))
		}
	}
	png.Encode(os.Stdout, img) //Note: ignoring errors
}

func avg(colors []color.Color) color.Color {
	var r, g, b, a uint16
	n := len(colors)
	for _, c := range colors {
		r1, g1, b1, a1 := c.RGBA()
		r += uint16(r1 / uint32(n))
		g += uint16(g1 / uint32(n))
		b += uint16(b1 / uint32(n))
		a += uint16(a1 / uint32(n))
	}
	return color.RGBA64{r, g, b, a}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	//const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.RGBA{255 - 15*n, 255 - n, 255 - 8*n, 54}
		}
	}
	return color.Black
}
