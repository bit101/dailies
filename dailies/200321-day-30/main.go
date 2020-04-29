package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

const (
	width    = 800.0
	height   = 800.0
	filename = "out.png"
	res      = 16.0
)

func main() {
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(0.25, 0.25, 0.25)
	surface.SetLineWidth(0.75)

	for y := 0.0; y < height; y += res {
		for x := 0.0; x < width; x += res {
			drawTile(surface, x, y, res)
		}
	}

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}

func drawTile(surface *blgo.Surface, x, y, res float64) {
	surface.Save()
	surface.Translate(x+res/2, y+res/2)

	if random.Boolean() {
		surface.Rotate(math.Pi / 2)
	}

	surface.SetSourceRGB(1, 1, 1)
	surface.FillRectangle(-res/4, -res/2, res/2, res)

	surface.SetSourceRGB(0, 0, 0)
	surface.MoveTo(-res/4, -res/2)
	surface.LineTo(-res/4, res/2)
	surface.MoveTo(res/4, -res/2)
	surface.LineTo(res/4, res/2)
	surface.Stroke()

	surface.SetSourceRGB(1, 1, 1)
	surface.FillRectangle(-res/2, -res/4, res, res/2)

	surface.SetSourceRGB(0, 0, 0)
	surface.MoveTo(-res/2, -res/4)
	surface.LineTo(res/2, -res/4)
	surface.MoveTo(-res/2, res/4)
	surface.LineTo(res/2, res/4)
	surface.Stroke()

	surface.Restore()
}
