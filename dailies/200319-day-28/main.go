package main

import (
	"fmt"
	"math"
	"os"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/util"
	"github.com/bit101/go-cairo"
)

const (
	filename = "out.png"
	res      = 10.0
)

var (
	width  float64
	height float64
)

func main() {
	image, _ := cairo.NewSurfaceFromPNG("panda.png")
	if image == nil {
		fmt.Println("Error loading png")
		os.Exit(1)
	}
	data := image.GetData()
	width = float64(image.GetWidth())
	height = float64(image.GetHeight())

	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	// surface.SetLineWidth(0.5)

	for y := 0.0; y < height; y += res {
		for x := 0.0; x < width; x += res {
			surface.Save()
			surface.Translate(x+res/2, y+res/2)
			b := getBrightnessFromPNG(x, y, data)
			// surface.SetSourceRGB(b, b, b)
			// surface.FillRectangle(x+1, y+1, res-2, res-2)
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

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}

func getBrightnessFromPNG(x, y float64, data []byte) float64 {
	val := 0
	for yy := int(y); yy < int(math.Min(y+res, height-1)); yy++ {
		for xx := int(x); xx < int(math.Min(x+res, width-1)); xx++ {
			index := (yy*int(width) + xx) * 4
			val += int(data[index])
			val += int(data[index+1])
			val += int(data[index+2])
		}
	}

	return 1.0 - float64(val)/(3*res*res)/255.0

}
