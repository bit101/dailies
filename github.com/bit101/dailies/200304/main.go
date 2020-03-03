package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/color"
	"github.com/bit101/blgo/noise"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

const (
	height  = 800.0
	width   = height * 3 / 2
	res     = 50.0
	maxIter = 40.0
)

var filename = util.ParentDir() + ".png"

func main() {
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(0.25)
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
	n := random.FloatRange(5, 50)
	f := -1.0
	if random.Boolean() {
		f = 1.0
	}
	surface.SetLineWidth(random.FloatRange(0.2, 0.4))
	s := 0.003
	offset := 50.0

	surface.SetSourceColor(color.RandomRGB())
	for j := 0.0; j <= n; j++ {
		aa := j * f
		r2 := r * (1 - j/n)
		for i := 0.0; i < 6.0; i++ {
			a := math.Pi/2 + i*math.Pi/3 + aa
			xx := x + math.Cos(a)*r2
			yy := y + math.Sin(a)*r2 - j*3
			nn := blmath.Map(noise.Perlin2(xx*s, yy*s), -1, 1, 0, math.Pi*2)
			surface.LineTo(xx+math.Cos(nn)*offset, yy+math.Sin(nn)*offset)
		}
		surface.ClosePath()
		surface.SetSourceRGB(1, 1, 1)
		surface.FillPreserve()
		surface.SetSourceRGB(0, 0, 0)
		surface.Stroke()
		// surface.Restore()
	}

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
