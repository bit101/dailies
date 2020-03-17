package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/util"
	"github.com/bit101/blgrids"
)

func main() {
	surface := blgo.NewSurface(800, 800)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(0.75)
	blgrids.HexLayout(0, 0, 800, 800, 40, func(x, y float64) {
		a := math.Pi / 2
		for r := 39.0; r > 0; r -= 2 {
			surface.StrokePolygon(x, y, r, 6, a)
			a += 0.12
		}
	})
	surface.WriteToPNG("out.png")
	util.ViewImage("out.png")
}
