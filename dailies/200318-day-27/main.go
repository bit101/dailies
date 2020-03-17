package main

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
	"github.com/bit101/blgrids"
)

func main() {
	surface := blgo.NewSurface(800, 800)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(0.75)
	// blgrids.Hex(surface, 0, 0, 800, 800, 40, 35)
	blgrids.HexWithRenderer(surface, 0, 0, 800, 800, 40, func(surface *blgo.Surface, r float64) {
		a := 0.0
		rot := -0.12
		if random.Boolean() {
			rot = 0.12
		}
		for r1 := r - 1; r1 > 0; r1 -= 2 {
			surface.StrokePolygon(0, 0, r1, 6, a)
			a += rot
		}
	})
	surface.WriteToPNG("out.png")
	util.ViewImage("out.png")
}
