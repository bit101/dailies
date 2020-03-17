package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/noise"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
	cairo "github.com/bit101/go-cairo"
)

const (
	height      = 400.0
	width       = 400.0
	res         = 10
	strokeWidth = res / 3
	drawOutline = false
	frames      = 150
	framesDir   = "frames"
	fps         = 30
)

var filename = util.ParentDir() + ".gif"
var surface *blgo.Surface
var perlinOffset = 15.0
var yoffset = 0.0

// removed transforms to make it easier to do something I plan for later
func main() {
	surface = blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.Translate(0.5, 0.5)
	surface.SetLineJoin(cairo.LineJoinRound)
	surface.SetLineCap(cairo.LineCapRound)

	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(render)
	util.ConvertToGIF(framesDir, filename, fps)

	util.ViewImage(filename)
}

func render(percent float64) {
	surface.ClearRGB(1, 1, 1)
	yoffset = blmath.LerpSin(percent, 0, -height/2)
	hexGrid(surface, 0, 0, width, height, res, drawHex)
}

func hexGrid(surface *blgo.Surface, x, y, w, h, r float64, render func(surface *blgo.Surface, x, y, r float64)) {
	sin60 := math.Sin(math.Pi / 3)
	xinc := r * 2 * sin60
	yinc := r * 1.5

	random.Seed(0)
	offset := 0.0

	for yy := y; yy < y+h*2+r; yy += yinc {
		for xx := x; xx < x+w+r; xx += xinc {
			drawHex(surface, xx+offset, yy, r)
		}
		if offset == 0 {
			offset = r * sin60
		} else {
			offset = 0
		}
	}
}

func drawHex(surface *blgo.Surface, x, y, r float64) {
	if drawOutline {
		surface.SetLineWidth(0.15)
		for i := 0.0; i < 6.0; i++ {
			a := math.Pi/3*i + math.Pi/2
			surface.LineTo(x+math.Cos(a)*r, y+math.Sin(a)*r)
		}
		surface.ClosePath()
		surface.Stroke()
	}

	drawTile(surface, x, y, r)
}

func drawTile(surface *blgo.Surface, x, y, r float64) {
	surface.SetLineWidth(strokeWidth)
	i := float64(random.IntRange(0, 5))
	drawLine(surface, i, x, y, r)
	drawLine(surface, i+2, x, y, r)
	drawLine(surface, i+4, x, y, r)
}

func drawLine(surface *blgo.Surface, i, x, y, r float64) {
	angle := i*math.Pi/3 + math.Pi/6
	xx := x + math.Cos(angle)*r
	yy := y + math.Sin(angle)*r + yoffset
	o := math.Pi / 3 * i
	s := 0.007
	// offset := r * 2
	for i := o + math.Pi*5/6; i <= o+math.Pi*9/6; i += 0.01 {
		xxx := xx + math.Cos(i)*r/2
		yyy := yy + math.Sin(i)*r/2
		n := noise.Perlin2(xxx*s, yyy*s) * math.Pi * 2
		surface.LineTo(xxx+math.Cos(n)*perlinOffset, yyy+math.Sin(n)*perlinOffset)
	}
	surface.Stroke()
}
