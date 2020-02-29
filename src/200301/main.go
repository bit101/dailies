package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/geom"
	"github.com/bit101/blgo/util"
)

type point3d struct {
	x float64
	y float64
	z float64
}

const (
	filename = "200229.png"
	width    = 800.0
	height   = 800.0
	xres     = 2.0
	yres     = 5.0
	fl       = 600.0
	pscale   = 0.02
	a        = -math.Pi / 4
	b        = math.Pi / 6
)

var (
	points []point3d
	path   []*geom.Point
)

func main() {
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(0.5)
	surface.Translate(width*0.5, height*0.5)

	points = append(points, point3d{-200.0, -200.0, -200.0})
	points = append(points, point3d{200.0, -200.0, -200.0})
	points = append(points, point3d{200.0, 200.0, -200.0})
	points = append(points, point3d{-200.0, 200.0, -200.0})

	points = append(points, point3d{-200.0, -200.0, 200.0})
	points = append(points, point3d{200.0, -200.0, 200.0})
	points = append(points, point3d{200.0, 200.0, 200.0})

	project(surface, points[:4])
	project(surface, append(points[:2], points[5], points[4]))
	project(surface, append(points[1:3], points[6], points[5]))
	project(surface, append(points[5:6], points[5]))

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}

func project(surface *blgo.Surface, points []point3d) {
	for _, p := range points {
		x := p.x*math.Cos(a) - p.z*math.Sin(a)
		z := p.z*math.Cos(a) + p.x*math.Sin(a)
		y := p.y*math.Cos(b) - z*math.Sin(b)
		z1 := z*math.Cos(b) + p.y*math.Sin(b)
		scale := fl / (fl + z1 + 200)
		path = append(path, geom.NewPoint(x*scale, y*scale))
	}
	surface.StrokePath(path, false)
}
