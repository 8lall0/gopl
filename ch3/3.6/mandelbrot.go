package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
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

	var offX = []float64{-epsX, epsX}
	var offY = []float64{-epsY, epsY}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < height; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin

			subPix := make([]color.Color, 0)
			for _, i := range offX {
				for _, j := range offY {
					z := complex(x+i, y+j)
					subPix = append(subPix, mandelbrot(z))
				}
			}

			img.Set(px, py, avgColor(subPix))
		}
	}

	png.Encode(os.Stdout, img)
}

func avgColor(colors []color.Color) color.Color {
	var r, g, b, a uint16
	n := uint32(len(colors))

	for _, color := range colors {
		_r, _g, _b, _a := color.RGBA()
		r += uint16(_r / n)
		g += uint16(_g / n)
		b += uint16(_b / n)
		a += uint16(_a / n)
	}

	return color.RGBA64{r, g, b, a}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 100
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			switch {
			case n > 50:
				return color.RGBA{100, 0, 0, 255}
			default:
				logScale := math.Log(float64(n)) / math.Log(float64(iterations))
				return color.RGBA{0, 0, 255 - uint8(logScale*255), 255}
			}
		}
	}

	return color.Black
}
