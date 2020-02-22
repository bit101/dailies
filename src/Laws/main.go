package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

const width = 800.0
const height = 800.0

func main() {
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)

	random.RandSeed()
	r0 := random.FloatRange(10, 200)
	r1 := random.FloatRange(10, 200)
	r2 := random.FloatRange(10, 200)
	a := r0 + r1
	b := r0 + r2
	c := r1 + r2

	x0, y0 := width/2, height/2

	angle1 := random.FloatRange(0, math.Pi*2)
	x1 := x0 + math.Cos(angle1)*a
	y1 := y0 + math.Sin(angle1)*a

	angle2 := angle1 + math.Acos((c*c-a*a-b*b)/(-2*a*b))
	x2 := x0 + math.Cos(angle2)*b
	y2 := y0 + math.Sin(angle2)*b

	circle(surface, x0, y0, r0)
	circle(surface, x1, y1, r1)
	circle(surface, x2, y2, r2)

	surface.WriteToPNG("out.png")
	util.ViewImage("out.png")
}

func circle(surface *blgo.Surface, x, y, r float64) {
	surface.StrokeCircle(x, y, r)
}
