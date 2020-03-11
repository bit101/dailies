package main

import (
	"fmt"
	"math"
	"os"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/util"
	"github.com/bit101/dailies/200313-day-22/truchet"
	"github.com/bit101/go-cairo"
)

const (
	tileSize = 5.0
	filename = "200314.png"
)

var image *cairo.Surface
var width float64
var height float64

// var pattern = truchet.PatternDouat72

var pattern = truchet.PatternDouat72

func main() {
	image, _ = cairo.NewSurfaceFromPNG("keith-peters.png")
	// image, _ = cairo.NewSurfaceFromPNG("boy-howdy.png")
	if image == nil {
		fmt.Println("Error loading png")
		os.Exit(1)
	}
	data := image.GetData()

	width = float64(image.GetWidth())
	height = float64(image.GetHeight())
	fmt.Printf("width = %+v\n", width)
	fmt.Printf("height = %+v\n", height)
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)

	for y := 0.0; y < height; y += tileSize {
		for x := 0.0; x < width; x += tileSize {
			brightness := getBrightnessFromPNG(x, y, data)
			truchet.Truchet(surface, pattern, x, y, tileSize, brightness)
		}
	}

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}

func getBrightness(x, y float64) float64 {
	dx := x - width/2
	dy := y - height/2
	dist := math.Sqrt(dx*dx + dy*dy)
	n := math.Sin(dist * 0.05)
	return blmath.Map(n, -1, 1, 0, 1)
}

func getBrightnessFromPNG(x, y float64, data []byte) float64 {
	val := 0
	for yy := int(y); yy < int(y+tileSize); yy++ {
		for xx := int(x); xx < int(x+tileSize); xx++ {
			xxx := int(math.Min(float64(xx), width-1))
			yyy := int(math.Min(float64(yy), height-1))
			index := (yyy*image.GetWidth() + xxx) * 4
			val += int(data[index])
			val += int(data[index+1])
			val += int(data[index+2])
		}
	}

	return 1.0 - float64(val)/(3*tileSize*tileSize)/255.0

}
