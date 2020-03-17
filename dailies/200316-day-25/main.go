package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
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
	p := blmath.LerpSin(percent, 0, 5)
	surface.ClearRGB(1, 1, 1)
	surface.Save()
	surface.Translate(width/2, height/2)
	for y := -height / 2; y < height/2; y += res {
		for x := -width / 2; x < width/2; x += res {
			x1 := x / width
			y1 := y / height
			w := (x + res - 3*p) / width
			h := (y + res - p) / height
			surface.MoveTo(swirl(x1, y1, p))
			surface.LineTo(swirl(w, y1, p))
			surface.LineTo(swirl(w, h, p))
			surface.LineTo(swirl(x1, h, p))
			surface.ClosePath()
			surface.Stroke()

		}
	}
	surface.Restore()
}

func swirl(x, y, s float64) (float64, float64) {
	r := math.Sqrt(x*x+y*y) * s
	x1 := x*math.Sin(r*r) - y*math.Cos(r*r)
	y1 := x*math.Cos(r*r) + y*math.Sin(r*r)
	return x1 * width, y1 * height
}
