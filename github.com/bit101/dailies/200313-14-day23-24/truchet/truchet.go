package truchet

import (
	"math"

	"github.com/bit101/blgo"
	"github.com/bit101/blgo/blmath"
)

// Truchet draw a truchet tile
func Truchet(surface *blgo.Surface, pattern Pattern, x, y, tileSize, brightness float64) {
	quarterTile := tileSize * 0.25
	halfTile := tileSize * 0.5
	tile := pattern.GetTile(int(y/tileSize), int(x/tileSize))

	t := blmath.Clamp(2*brightness-0.5, 0, 1)
	midx := blmath.Map(t, 0, 1, -quarterTile, quarterTile)
	midy := blmath.Map(t, 0, 1, quarterTile, -quarterTile)

	surface.Save()
	surface.Translate(x+tileSize/2, y+tileSize/2)
	surface.Rotate(float64(tile) * math.Pi / 2)

	surface.MoveTo(-halfTile, -halfTile)
	surface.LineTo(midx, midy)
	surface.LineTo(halfTile, halfTile)
	surface.LineTo(-halfTile, halfTile)
	surface.Fill()

	surface.Restore()
}
