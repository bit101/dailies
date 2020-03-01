package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/util"
)

const filename = "200224.png"
const xres = 1.3
const yres = 1.3
const maxIter = 100.0

var width float64
var height float64

func main() {
	height = 800.0
	width = height * 3 / 2
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetSourceRGB(0, 0, 0)
	surface.Translate(0.5, 0.5)
	surface.SetLineWidth(0.1)
	ll := 20.0

	for x := 0.0; x < width+xres; x += xres {
		for y := 0.0; y < height; y += yres {
			r := getAngle(x, y)
			if r < math.Pi*2 {
				surface.MoveTo(x, y)
				surface.LineTo(x+math.Cos(r)*ll, y+math.Sin(r)*ll)
			}
		}
	}
	surface.Stroke()

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}

func getAngle(x, y float64) float64 {
	rrange := 0.001
	irange := rrange * height / width
	rc := -1.2624
	ic := -0.4082
	rmin := rc - rrange/2
	rmax := rmin + rrange
	imin := ic - irange/2
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
