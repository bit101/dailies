package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/noise"
	"github.com/bit101/blgo/util"
)

const (
	filename = "200302.png"
	width    = 800.0
	height   = 800.0
	res      = 20.0
)

func main() {
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(0.5)
	surface.Translate(0.5, 0.5)
	even := true

	xinc := res * 2 * math.Sin(math.Pi/3)
	yinc := res * 1.5
	evenOffset := math.Sin(math.Pi/3) * res

	for y := -100.0; y < height+res+100; y += yinc {
		for x := -100.0; x < width+res+100; x += xinc {
			if even {
				drawHex(surface, x+evenOffset, y, res)
			} else {
				drawHex(surface, x, y, res)
			}
		}
		even = !even
	}

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}

func drawHex(surface *blgo.Surface, x, y, r float64) {
	surface.SetSourceRGB(0.4, 0.4, 0.4)
	surface.MoveTo(perlinize(x, y))
	surface.LineTo(getPoint(x, y, r, 0))
	surface.LineTo(getPoint(x, y, r, 1))
	surface.LineTo(getPoint(x, y, r, 2))
	surface.Fill()

	surface.SetSourceRGB(0.8, 0.8, 0.8)
	surface.MoveTo(perlinize(x, y))
	surface.LineTo(getPoint(x, y, r, 2))
	surface.LineTo(getPoint(x, y, r, 3))
	surface.LineTo(getPoint(x, y, r, 4))
	surface.Fill()

	surface.SetSourceRGB(0.6, 0.6, 0.6)
	surface.MoveTo(perlinize(x, y))
	surface.LineTo(getPoint(x, y, r, 4))
	surface.LineTo(getPoint(x, y, r, 5))
	surface.LineTo(getPoint(x, y, r, 0))
	surface.Fill()
}

func getPoint(x, y, r float64, i int) (float64, float64) {
	a := math.Pi/2 + float64(i)*math.Pi/3
	return perlinize(x+math.Cos(a)*r, y+math.Sin(a)*r)
}

func perlinize(x, y float64) (float64, float64) {
	offset := 30.0
	scale := 0.0035
	n := noise.Perlin2(x*scale, y*scale+50) * math.Pi * 2
	return x + math.Cos(n)*offset, y + math.Sin(n)*offset
}
