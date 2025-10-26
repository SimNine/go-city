package world

import (
	"image/color"
	"math/rand"

	"github.com/SimNine/go-city/world/tiles"
	"github.com/SimNine/go-urfutils/src/geom"
	"github.com/hajimehoshi/ebiten/v2"
)

var COLOR_SKYBLUE = color.RGBA{
	R: 100,
	G: 181,
	B: 246,
	A: 255,
}

func NewWorld(
	random *rand.Rand,
	dims geom.Dims[int],
) *World {

	// Create the world struct
	env := &World{
		random: random,

		tiles: make([][]tiles.Tile, dims.X, dims.Y),
	}

	return env
}

type World struct {
	random *rand.Rand

	tiles [][]tiles.Tile
}

func (e *World) Draw(
	screen *ebiten.Image,
	viewport geom.Viewport[int],
) {
	// Fill the background with blue
	screen.Fill(COLOR_SKYBLUE)
}

func (e *World) Update() {
	// do nothing for now
}
