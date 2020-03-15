package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

const (
	timeInSeconds = 5
	fps           = 30
	frames        = timeInSeconds * fps
	framesDir     = "frames"
	outFileName   = "out.gif"
	width         = 400.0
	height        = 400.0
	res           = 20.0
)

var surface *blgo.Surface

func main() {
	surface = blgo.NewSurface(width, height)
	surface.SetLineWidth(0.5)
	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(render)
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}

func render(percent float64) {
	p := blmath.LerpSin(percent, -5, 5)
	random.Seed(0)
	surface.ClearRGB(1, 1, 1)
	surface.Save()
	surface.Translate(width/2, height/2)

	for i := 0; i < 100000; i++ {
		x := random.FloatRange(-0.5, 0.5)
		y := random.FloatRange(-0.5, 0.5)
		x1, y1 := polar(x, y, p)
		surface.FillRectangle(x1, y1, 1, 1)
	}

	surface.Restore()
}

func polar(x, y, s float64) (float64, float64) {
	r := math.Sqrt(x*x+y*y) * s
	t := math.Atan2(y, x)
	x1 := 1 / r * (math.Cos(t) + math.Sin(r))
	y1 := 1 / r * (math.Sin(t) - math.Cos(r))
	return x1 * width / 2, y1 * height / 2
}
