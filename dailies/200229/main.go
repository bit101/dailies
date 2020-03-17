package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/noise"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

const (
	filename = "200229.png"
	width    = 800.0
	height   = 800.0
	res      = 20.0
)

func main() {
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(0.5)
	surface.Translate(0.5, 0.5)
	even := true

	xinc := res * 2 * math.Sin(math.Pi/3)
	yinc := res * 1.5
	evenOffset := math.Sin(math.Pi/3) * res

	for y := 0.0; y < height+res; y += yinc {
		for x := 0.0; x < width+res; x += xinc {
			if even {
				drawHex(surface, x+evenOffset, y, res-2)
			} else {
				drawHex(surface, x, y, res-2)
			}
		}
		even = !even
	}

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}

func drawHex(surface *blgo.Surface, x, y, r float64) {
	offset := 40.0
	scale := 0.003
	surface.Save()
	for i := 0.0; i < 6.0; i++ {
		a := math.Pi/2 + i*math.Pi/3
		xx := x + math.Cos(a)*r
		yy := y + math.Sin(a)*r
		n := noise.Perlin2(xx*scale, yy*scale) * math.Pi * 2
		surface.LineTo(xx+math.Cos(n)*offset, yy+math.Sin(n)*offset)
	}
	surface.ClosePath()
	surface.SetSourceRGB(1, 1, 1)
	surface.FillPreserve()
	surface.SetSourceRGB(0, 0, 0)
	surface.StrokePreserve()
	surface.Clip()

	size := res + offset
	surface.Translate(x, y)
	surface.Rotate(random.FloatRange(0, math.Pi*2))
	for i := -size; i < size; i += 3 {
		surface.MoveTo(-size, i)
		surface.LineTo(size, i)
		// surface.MoveTo(x+random.FloatRange(-size, size), y+random.FloatRange(-size, size))
		// surface.LineTo(x+random.FloatRange(-size, size), y+random.FloatRange(-size, size))
	}
	surface.Stroke()
	surface.Restore()

}
