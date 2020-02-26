package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/util"
)

const (
	mp4FileName = "200226.mp4"
	xres        = 2.0
	yres        = 2.0

	width     = 640.0
	height    = 360.0
	frames    = 300
	framesDir = "frames"
	fps       = 30
)

var (
	surface *blgo.Surface
	maxIter = 100.0
)

func main() {
	surface = blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetSourceRGB(0, 0, 0)
	surface.SetLineWidth(0.1)

	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(render)
	util.ConvertToYoutube(framesDir, mp4FileName, fps)

}

func render(percent float64) {
	surface.ClearRGB(1, 1, 1)
	offset := percent * math.Pi * 2.0
	ll := 20.0
	// ll := percent*20.0 + 10.0
	// maxIter = 80.0*percent + 20.0

	for x := 0.0; x < width+xres; x += xres {
		for y := 0.0; y < height; y += yres {
			r := getAngle(x, y)
			if r < math.Pi*2 {
				surface.MoveTo(x, y)
				surface.LineTo(x+math.Cos(r+offset)*ll, y+math.Sin(r+offset)*ll)
			}
		}
	}
	surface.Stroke()
}

func getAngle(x, y float64) float64 {
	rrange := 0.001
	irange := rrange * height / width
	rc := -1.2624
	ic := -0.4082
	rmin := rc - rrange/2
	rmax := rmin + rrange
	imin := ic - irange/2
	imax := imin + irange
	r := blmath.Map(x, 0, width, rmin, rmax)
	i := blmath.Map(y, 0, height, imin, imax)

	c := complex(r, i)
	z := complex(0, 0)
	n := 0.0
	for n = 0.0; n < maxIter && real(z) < 4.0; n++ {
		z = z*z + c
	}
	return n / maxIter * math.Pi * 2

}
