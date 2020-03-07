package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
	cairo "github.com/bit101/go-cairo"
)

// todo: remove transforms, draw lines by short segments, offset by pnoise.
const (
	height      = 800.0
	width       = 800.0
	res         = 40.0
	strokeWidth = res / 5
	drawOutline = true
)

var funcs = []func(x, y float64, surface *blgo.Surface, r2 float64){
	drawTile1,
	drawTile2,
	drawTile3,
	drawTile4,
	drawTile5,
	drawTile6,
}

var filename = util.ParentDir() + ".png"

func main() {
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.Translate(0.5, 0.5)
	surface.SetLineJoin(cairo.LineJoinRound)
	surface.SetLineCap(cairo.LineCapRound)

	hexGrid(surface, 0, 0, width, height, res, drawHex)

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}

func hexGrid(surface *blgo.Surface, x, y, w, h, r float64, render func(x, y float64, surface *blgo.Surface, r float64)) {
	sin60 := math.Sin(math.Pi / 3)
	xinc := r * 2 * sin60
	yinc := r * 1.5

	offset := 0.0
	for yy := y; yy < y+h+r; yy += yinc {
		for xx := x; xx < x+w+r; xx += xinc {
			// surface.Save()
			// surface.Translate(xx+offset, yy)
			// surface.Rotate(math.Pi / 2)
			drawHex(xx+offset, yy, surface, r)
			// surface.Restore()
		}
		if offset == 0 {
			offset = r * sin60
		} else {
			offset = 0
		}
	}
}

func drawHex(x, y float64, surface *blgo.Surface, r float64) {
	r2 := r * math.Sin(math.Pi/3)
	if drawOutline {
		surface.SetLineWidth(0.15)
		for i := 0.0; i < 6.0; i++ {
			a := math.Pi / 3 * i
			surface.LineTo(x+math.Cos(a)*r, y+math.Sin(a)*r)
		}
		surface.ClosePath()
		// surface.SetSourceRGB(1, 1, 1)
		// surface.FillPreserve()
		surface.SetSourceRGB(0, 0, 0)
		surface.Stroke()
	}

	// drawConnections(surface, r2)

	// i := random.IntRange(0, 5)
	// surface.Rotate(float64(i) * math.Pi / 3)

	surface.SetLineWidth(strokeWidth)
	j := random.IntRange(0, len(funcs)-1)
	funcs[j](x, y, surface, r2)
}

func drawTile1(x, y float64, surface *blgo.Surface, r2 float64) {
	drawLine1(surface, r2)
	surface.Rotate(math.Pi)
	drawLine1(surface, r2)
}

func drawTile2(x, y float64, surface *blgo.Surface, r2 float64) {
	drawLine2(surface, r2)
	surface.Rotate(math.Pi)
	drawLine2(surface, r2)
}

func drawTile3(x, y float64, surface *blgo.Surface, r2 float64) {
	drawLine3(surface, r2)
	surface.Rotate(math.Pi / 3)
	drawLine1(surface, r2)
	surface.Rotate(math.Pi)
	drawLine1(surface, r2)
}

func drawTile4(x, y float64, surface *blgo.Surface, r2 float64) {
	drawLine1(surface, r2)
	surface.Rotate(math.Pi * 2 / 3)
	drawLine1(surface, r2)
	surface.Rotate(math.Pi * 2 / 3)
	drawLine1(surface, r2)
}

func drawTile5(x, y float64, surface *blgo.Surface, r2 float64) {
	drawLine3(surface, r2)
	surface.Rotate(math.Pi * 2 / 3)
	drawLine2(surface, r2)
	surface.Rotate(math.Pi)
	drawLine2(surface, r2)
}

func drawTile6(x, y float64, surface *blgo.Surface, r2 float64) {
	drawLine3(surface, r2)
}

func drawLine1(surface *blgo.Surface, r2 float64) {
	surface.MoveTo(math.Cos(math.Pi/6)*r2, math.Sin(math.Pi/6)*r2)
	surface.LineTo(math.Cos(math.Pi/6)*r2/2, math.Sin(math.Pi/6)*r2/2)
	surface.LineTo(math.Cos(math.Pi/2)*r2/2, math.Sin(math.Pi/2)*r2/2)
	surface.LineTo(math.Cos(math.Pi/2)*r2, math.Sin(math.Pi/2)*r2)
	surface.Stroke()
}

func drawLine2(surface *blgo.Surface, r2 float64) {
	surface.MoveTo(math.Cos(math.Pi/6)*r2, math.Sin(math.Pi/6)*r2)
	surface.LineTo(math.Cos(math.Pi/6)*r2/2, math.Sin(math.Pi/6)*r2/2)
	surface.LineTo(math.Cos(math.Pi*5/6)*r2/2, math.Sin(math.Pi*5/6)*r2/2)
	surface.LineTo(math.Cos(math.Pi*5/6)*r2, math.Sin(math.Pi*5/6)*r2)
	surface.Stroke()
}

func drawLine3(surface *blgo.Surface, r2 float64) {
	surface.MoveTo(math.Cos(math.Pi/6)*r2, math.Sin(math.Pi/6)*r2)
	surface.LineTo(math.Cos(math.Pi*7/6)*r2, math.Sin(math.Pi*7/6)*r2)
	surface.Stroke()
}

func drawConnections(surface *blgo.Surface, r2 float64) {
	for i := 0.0; i < 6.0; i++ {
		a := math.Pi/3*i + math.Pi/6
		surface.FillCircle(math.Cos(a)*r2, math.Sin(a)*r2, strokeWidth/2)
	}
}
