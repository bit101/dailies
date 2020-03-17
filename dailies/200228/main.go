package main

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/noise"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

const (
	filename = "200228.png"
	width    = 800.0
	height   = 800.0
	res      = 40.0
)

var ()

func main() {
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(0.25)
	surface.Translate(0.5, 0.5)

	for y := 0.0; y < height+res; y += res {
		for x := 0.0; x < width+res; x += res {
			drawRect(surface, x, y)
		}
	}

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}

func drawRect(surface *blgo.Surface, x, y float64) {

	surface.Save()
	surface.Translate(x, y)
	surface.Rectangle(0, 0, res, res)
	surface.StrokePreserve()
	surface.Clip()

	scale := 0.005
	n := noise.Perlin2(x*scale, y*scale)
	count := blmath.Map(n, -0.75, 1, 10, 800)
	grey := blmath.Map(n, -0.75, 0.75, 0, 1)
	surface.SetSourceRGB(grey, grey, grey)
	for i := 0.0; i < count; i++ {
		surface.MoveTo(
			random.FloatRange(-res, res*2),
			random.FloatRange(-res, res*2),
		)
		surface.LineTo(
			random.FloatRange(-res, res*2),
			random.FloatRange(-res, res*2),
		)
	}
	surface.Stroke()
	surface.Restore()
}
