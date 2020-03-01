package main

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/util"
)

const filename = "200221.png"
const xres = 1.0
const yres = 2.0
const maxIter = 200.0

var width float64
var height float64

func main() {
	width = 800.0
	height = 800.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(0.25)

	for y := -100.0; y < height; y += yres {
		for x := 0.0; x < width+xres; x += xres {
			surface.LineTo(x, y+getHeight(x, y)+0.5)
		}
		surface.LineTo(width, height)
		surface.LineTo(0, height)
		surface.SetSourceRGB(1, 1, 1)
		surface.FillPreserve()

		surface.SetSourceRGB(0, 0, 0)
		surface.Stroke()
	}

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}

func getHeight(x, y float64) float64 {
	rrange := 0.05
	irange := rrange * height / width * 1.5
	rmin := -0.16
	rmax := rmin + rrange
	imin := -0.9
	imax := imin + irange

	r := blmath.Map(x, 0, width, rmin, rmax)
	i := blmath.Map(y, 0, height, imin, imax)

	c := complex(r, i)
	z := complex(0, 0)
	n := 0.0
	for n = 0.0; n < maxIter && real(z) < 4.0; n++ {
		z = z*z + c
	}
	return n / maxIter * 60

}
