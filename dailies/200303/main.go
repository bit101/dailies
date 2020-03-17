package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/util"
)

const (
	height  = 800.0
	width   = height * 3 / 2
	res     = 5.0
	maxIter = 40.0
)

var filename = util.ParentDir() + ".png"

func main() {
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(0.5)
	surface.Translate(0.5, 0.5)
	even := true

	xinc := res * 2 * math.Sin(math.Pi/3)
	yinc := res * 1.5
	evenOffset := math.Sin(math.Pi/3) * res

	for y := 0.0; y < height+res; y += yinc {
		for x := 0.0; x < width+res; x += xinc {
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
	offset := 100.0
	surface.Save()
	for i := 0.0; i < 6.0; i++ {
		a := math.Pi/2 + i*math.Pi/3
		xx := x + math.Cos(a)*r
		yy := y + math.Sin(a)*r
		// n := noise.Perlin2(xx*scale, yy*scale) * math.Pi * 2
		n := getHeight(xx, yy)
		surface.LineTo(xx, yy+n*offset)
	}
	surface.ClosePath()
	g := blmath.Map(getHeight(x, y), 0, 1, 1, 0.25)
	surface.SetSourceRGB(g, g, g)
	surface.FillPreserve()
	surface.SetSourceRGB(0, 0, 0)
	surface.Stroke()
	surface.Restore()

}

func getHeight(x, y float64) float64 {
	rrange := 3.0
	irange := rrange * height / width * 1.2
	rmin := -2.0
	rmax := rmin + rrange
	imin := -1.0
	imax := imin + irange

	r := blmath.Map(x, 0, width, rmin, rmax)
	i := blmath.Map(y, 0, height, imin, imax)

	c := complex(r, i)
	z := complex(0, 0)
	n := 0.0
	for n = 0.0; n < maxIter && real(z) < 4.0; n++ {
		z = z*z + c
	}
	return n / maxIter

}
