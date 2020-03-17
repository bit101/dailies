package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

const (
	height = 800.0
	width  = 800.0
	res    = 40.0
	xres   = 2.0
	yres   = 3.0
)

var methods []func(surface *blgo.Surface)

var filename = util.ParentDir() + ".png"

func main() {
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.Translate(0.5, 0.5)
	surface.SetLineWidth(2)

	buffer := blgo.NewSurface(width, height)
	buffer.ClearRGB(1, 1, 1)

	methods = []func(surface *blgo.Surface){
		tile0,
		tile1,
		tile2,
		tile3,
		tile4,
	}

	for y := 0.0; y < height+res; y += res {
		for x := 0.0; x < width+res; x += res {
			drawTile(buffer, x, y)
		}
	}

	for y := 0; y < height; y += yres {
		for x := 0; x < width+xres; x += xres {
			yy := float64(y)
			r, _, _, _ := buffer.GetPixel(x, y)
			if float64(r) < 128 {
				yy -= 3
			}
			yy += random.FloatRange(-0.5, 0.5)
			surface.LineTo(float64(x), yy)
		}
		surface.SetSourceRGB(0, 0, 0)
		surface.StrokePreserve()
		surface.SetSourceRGB(1, 1, 1)
		surface.LineTo(width, height)
		surface.LineTo(0, height)
		surface.Fill()
	}
	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}

func drawTile(surface *blgo.Surface, x, y float64) {
	surface.Save()
	surface.Rectangle(x, y, res, res)
	surface.Clip()

	surface.Translate(x+res/2, y+res/2)
	i := random.IntRange(0, 4)
	r := random.IntRange(0, 3)
	surface.Rotate(float64(r) * math.Pi / 2)
	methods[i](surface)
	surface.Restore()
}

func tile0(surface *blgo.Surface) {
	surface.FillCircle(0, -res/2, res/8)
	surface.FillCircle(0, res/2, res/8)
	surface.FillRectangle(-res/2, -res/8, res, res/4)
}

func tile1(surface *blgo.Surface) {
	surface.FillCircle(-res/2, 0, res/8)
	surface.FillCircle(res/2, 0, res/8)
	surface.FillCircle(0, -res/2, res/8)
	surface.FillCircle(0, res/2, res/8)
}

func tile2(surface *blgo.Surface) {
	surface.FillRectangle(-res/2, -res/8, res, res/4)
	surface.FillRectangle(-res/8, -res/2, res/4, res)
}

func tile3(surface *blgo.Surface) {
	surface.SetLineWidth(res / 4)
	surface.StrokeCircle(res/2, -res/2, res/2)
	surface.FillCircle(0, res/2, res/8)
	surface.FillCircle(-res/2, 0, res/8)
}

func tile4(surface *blgo.Surface) {
	surface.SetLineWidth(res / 4)
	surface.StrokeCircle(res/2, -res/2, res/2)
	surface.StrokeCircle(-res/2, res/2, res/2)
}
