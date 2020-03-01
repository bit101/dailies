package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/random"
	"github.com/bit101/blgo/util"
)

const (
	outFileName = "200225.gif"
	mp4FileName = "200225.mp4"
	xres        = 1.0
	yres        = 2.0

	width     = 1280.0
	height    = 720.0
	frames    = 300
	framesDir = "frames"
	fps       = 30
)

var (
	surface *blgo.Surface
	maxIter float64
)

func main() {
	surface = blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(0.25)
	animation := anim.NewAnimation(surface, frames, framesDir)
	animation.Render(render)
	// util.ConvertToGIF(framesDir, outFileName, fps)
	convertToYoutube(framesDir, mp4FileName, fps)

	util.VLC(mp4FileName)
	// util.ViewImage(outFileName)
}

func render(percent float64) {
	random.Seed(0)
	maxIter = 10.0 + percent*300.0
	for y := -100.0; y < height; y += yres {
		for x := 0.0; x < width+xres; x += xres {
			surface.LineTo(x, y+getHeight(x, y)+random.FloatRange(-1, 0.25))
		}
		surface.LineTo(width, height)
		surface.LineTo(0, height)
		surface.SetSourceRGB(1, 1, 1)
		surface.FillPreserve()

		surface.SetSourceRGB(0, 0, 0)
		surface.Stroke()
	}
}

func getHeight(x, y float64) float64 {
	rrange := 0.05
	irange := rrange * height / width * 1.5
	rmin := -0.16
	rmax := rmin + rrange
	imin := -0.9
	imax := imin + irange

	r := blmath.Map(x, 0, width, rmin, rmax)
	i := blmath.Map(y, 0, height, imin, imax)

	c := complex(r, i)
	z := complex(0, 0)
	n := 0.0
	for n = 0.0; n < maxIter && real(z) < 4.0; n++ {
		z = z*z + c
	}
	return n / maxIter * 60

}

func convertToYoutube(folder, outFileName string, fps int) {
	path := folder + "/frame_%04d.png"
	fpsArg := fmt.Sprintf("%d", fps)

	cmd := exec.Command("ffmpeg", "-framerate", fpsArg, "-i", path, "-s:v", "1280x720",
		"-c:v", "libx264", "-profile:v", "high", "-crf", "20",
		"-pix_fmt", "yuv420p", outFileName)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
