package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
	cairo "github.com/bit101/go-cairo"
)

const (
	height = 800.0
	width  = 800.0
	res    = 50.0
)

var filename = util.ParentDir() + ".png"

func main() {
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(0.25)
	surface.Translate(0.5, 0.5)
	surface.SetLineJoin(cairo.LineJoinRound)

	hexGrid(surface, 0, 0, width, height, 40, drawHex)

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}

func hexGrid(surface *blgo.Surface, x, y, w, h, r float64, render func(surface *blgo.Surface, r float64)) {
	sin60 := math.Sin(math.Pi / 3)
	xinc := r * 2 * sin60
	yinc := r * 1.5

	offset := 0.0
	for yy := y; yy < y+h+r; yy += yinc {
		for xx := x; xx < x+w+r; xx += xinc {
			surface.Save()
			surface.Translate(xx+offset, yy)
			surface.Rotate(math.Pi / 2)
			drawHex(surface, r)
			surface.Restore()
		}
		if offset == 0 {
			offset = r * sin60
		} else {
			offset = 0
		}
	}
}

func drawHex(surface *blgo.Surface, r float64) {
	r2 := r * math.Sin(math.Pi/3)
	for i := 0.0; i < 6.0; i++ {
		a := math.Pi / 3 * i
		surface.LineTo(math.Cos(a)*r, math.Sin(a)*r)
	}
	surface.ClosePath()
	surface.SetSourceRGB(1, 1, 1)
	surface.FillPreserve()
	surface.SetSourceRGB(0, 0, 0)
	surface.Stroke()

	for i := 0.0; i < 6.0; i++ {
		a := math.Pi/3*i + math.Pi/6
		surface.FillCircle(math.Cos(a)*r2, math.Sin(a)*r2, 4)
	}

	i := random.IntRange(0, 5)
	surface.Rotate(float64(i) * math.Pi / 3)
	surface.SetLineWidth(8)
	surface.MoveTo(math.Cos(math.Pi/6)*r2, math.Sin(math.Pi/6)*r2)
	surface.LineTo(math.Cos(math.Pi/6)*r2/2, math.Sin(math.Pi/6)*r2/2)
	surface.LineTo(math.Cos(math.Pi*5/6)*r2/2, math.Sin(math.Pi*5/6)*r2/2)
	surface.LineTo(math.Cos(math.Pi*5/6)*r2, math.Sin(math.Pi*5/6)*r2)
	surface.Stroke()

	surface.Rotate(math.Pi)
	surface.MoveTo(math.Cos(math.Pi/6)*r2, math.Sin(math.Pi/6)*r2)
	surface.LineTo(math.Cos(math.Pi/6)*r2/2, math.Sin(math.Pi/6)*r2/2)
	surface.LineTo(math.Cos(math.Pi*5/6)*r2/2, math.Sin(math.Pi*5/6)*r2/2)
	surface.LineTo(math.Cos(math.Pi*5/6)*r2, math.Sin(math.Pi*5/6)*r2)
	surface.Stroke()
}
