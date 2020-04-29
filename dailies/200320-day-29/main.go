package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/util"
)

const (
	width         = 400.0
	height        = 400.0
	res           = 10.0
	timeInSeconds = 10
	fps           = 30
	frames        = timeInSeconds * fps
	framesDir     = "frames"
	outFileName   = "out.gif"
)

var surface *blgo.Surface

func main() {
	surface = blgo.NewSurface(width, height)
	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(render)
	util.ConvertToGIF(framesDir, outFileName, fps)
	util.ViewImage(outFileName)
}

func render(percent float64) {
	surface.ClearRGB(1, 1, 1)
	offset := percent * math.Pi * 2
	for y := 0.0; y < height; y += res {
		for x := 0.0; x < width; x += res {
			surface.Save()
			surface.Translate(x+res/2, y+res/2)
			dx := x + res/2 - width/2
			dy := y + res/2 - height/2
			dist := math.Sqrt(dx*dx + dy*dy)
			b := blmath.Map(math.Sin(dist*0.05+offset), -1, 1, 0, 1)
			if b < 0.5 {
				surface.Rotate(math.Pi / 2)
			}

			for i := 0.0; i < res; i += 5 {
				surface.MoveTo(-res/2+i, -res/2)
				surface.LineTo(res/2, res/2-i)
				surface.MoveTo(-res/2, -res/2+i)
				surface.LineTo(res/2-i, res/2)
			}
			surface.Stroke()

			surface.Restore()
		}
	}
}
