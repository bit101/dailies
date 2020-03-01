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
	filename  = "200227.gif"
	width     = 300.0
	height    = 300.0
	radius    = 100.0
	res       = 0.01
	scale     = 0.0017
	offset    = 50.0
	framesDir = "frames"
	frames    = 70
)

var (
	surface *blgo.Surface
)

func main() {
	surface = blgo.NewSurface(width, height)
	surface.SetSourceRGB(1, 0.5, 0)
	surface.SetLineWidth(0.5)
	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(render)
	util.ConvertToGIF(framesDir, filename, frames)

	util.ViewImage(filename)
}

func render(percent float64) {
	yy := blmath.Lerp(percent, -radius*2, height+radius*2)
	surface.ClearRGB(0, 0, 0.6)

	for i := 0.0; i < math.Pi*2; i += res {
		x, y := math.Cos(i)*radius+width/2, math.Sin(i)*radius*2+yy
		a := noise.Perlin2(x*scale, y*scale) * offset
		surface.LineTo(x+math.Cos(a)*offset, y+math.Sin(a))
	}
	surface.ClosePath()
	surface.Fill()
}
