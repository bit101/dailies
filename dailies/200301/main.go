package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/noise"
	"github.com/bit101/blgo/util"
)

const (
	filename  = "200301.gif"
	width     = 360.0
	height    = 360.0
	xres      = 2.0
	yres      = 5.0
	fl        = 600.0
	pscale    = 0.02
	a         = -math.Pi / 4
	b         = math.Pi / 6
	frames    = 55
	framesDir = "frames"
	size      = 140.0
)

var (
	surface *blgo.Surface
)

func main() {
	surface = blgo.NewSurface(width, height)
	surface.SetLineWidth(0.5)

	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(render)

	util.ConvertToGIF(framesDir, filename, frames)

	util.ViewImage(filename)
}

func render(percent float64) {
	offset := blmath.LerpSin(percent, 0.0, 20.0)
	surface.Save()
	surface.ClearRGB(1, 1, 1)
	surface.Translate(width*0.5, height*0.4)

	for y := -size; y <= size; y += yres {
		for x := -size; x <= size; x += xres {
			z := -size + noise.Perlin2(x*pscale, y*pscale)*offset

			xx, yy := rotateAndProject(a, x, y, z)
			surface.LineTo(xx, yy)
		}
		surface.Stroke()
	}

	for y := -size; y <= size; y += yres {
		for x := -size; x <= size; x += xres {
			z := -size + noise.Perlin2(x*pscale, y*pscale)*(20.0-offset)

			xx, yy := rotateAndProject(-a, x, y, z)
			surface.LineTo(xx, yy)
		}
		surface.Stroke()
	}
	surface.Restore()

}

func rotateAndProject(a, x, y, z float64) (float64, float64) {
	c := math.Cos(a)
	s := math.Sin(a)
	c2 := math.Cos(math.Pi / 6)
	s2 := math.Sin(math.Pi / 6)
	xx := x*c - z*s
	zz := z*c + x*s
	yy := y*c2 - zz*s2
	zzz := zz*c2 + y*s2
	scale := fl / (fl + zzz + 150.0)
	return xx * scale, yy * scale
}
