package main

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/blmath"
	"github.com/bit101/blgo/noise"
	"github.com/bit101/blgo/util"
)

const (
	tileSize = 5.0
	filename = "200314.png"
	width    = 800.0
	height   = 800.0
)

func main() {
	surface := blgo.NewSurface(width, height)
	surface.ClearRGB(1, 1, 1)
	surface.SetLineWidth(0.5)
	pf := NewPerlinFlow(0.04, 5)

	a := 0.0
	for r := 10.0; r < width/2; {
		x, y := width/2+math.Cos(a)*r, height/2+math.Sin(a)*r
		surface.LineTo(pf.TransformPoint(x, y))
		a += 0.01
		r += 0.005
	}
	surface.Stroke()

	surface.WriteToPNG(filename)
	util.ViewImage(filename)
}

// PerlinFlow represents a Perlin noise flow field
type PerlinFlow struct {
	scale    float64
	strength float64
	angle    float64
	z        float64
}

// NewPerlinFlow creates a new PerlinFlow
func NewPerlinFlow(scale, strength float64) *PerlinFlow {
	return &PerlinFlow{scale, strength, -math.Pi / 2, 0.0}
}

// SetScale sets the scale of the field
func (p *PerlinFlow) SetScale(scale float64) {
	p.scale = scale
}

// SetStrength sets the strength of the transformation of the field
func (p *PerlinFlow) SetStrength(strength float64) {
	p.strength = strength
}

// SetAngle sets the base angle of the flow vectors. Default = -math.Pi / 2
func (p *PerlinFlow) SetAngle(angle float64) {
	p.angle = angle
}

// SetZ sets the z axis value used for the perlin noise
func (p *PerlinFlow) SetZ(z float64) {
	p.z = z
}

// TransformPoint transforms an x, y point along the Perlin noise field
func (p *PerlinFlow) TransformPoint(x, y float64) (float64, float64) {
	angle := blmath.Map(noise.Perlin(x*p.scale, y*p.scale, p.z), -1, 1, 0, math.Pi*2) + p.angle
	return x + math.Cos(angle)*p.strength, y + math.Sin(angle)*p.strength
}

// GetHeight returns the height of the flow field mapped to a range
func (p *PerlinFlow) GetHeight(x, y, min, max float64) float64 {
	return blmath.Map(noise.Perlin(x*p.scale, y*p.scale, p.z), -1, 1, min, max)
}
