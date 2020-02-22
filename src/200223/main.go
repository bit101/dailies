package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/util"
)

const filename = "200223.png"
const xres = 1.0
const yres = 1.0
const maxIter = 400.0

var width float64
var height float64

func main() {
	width = 800.0
	height = 800.0
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetSourceRGB(0, 0, 0)
	surface.Translate(0.5, 0.5)
	surface.SetLineWidth(0.1)
	ll := 12.0

	for x := 0.0; x < width+xres; x += xres {
		for y := 0.0; y < height; y += yres {
			r := getAngle(x, y)
			surface.MoveTo(x, y)
			surface.LineTo(x+math.Cos(r)*ll, y+math.Sin(r)*ll)
		}
	}
	surface.Stroke()

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}

func getAngle(x, y float64) float64 {
	rrange := 0.007
	irange := rrange * height / width
	rmin := -0.255
	rmax := rmin + rrange
	imin := 0.637
	imax := imin + irange

	r := blmath.Map(x, 0, width, rmin, rmax)
	i := blmath.Map(y, 0, height, imin, imax)

	c := complex(r, i)
	z := complex(0, 0)
	n := 0.0
	for n = 0.0; n < maxIter && real(z) < 4.0; n++ {
		z = z*z + c
	}
	return n / maxIter * math.Pi * 2

}
