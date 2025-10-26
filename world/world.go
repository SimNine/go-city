package world

import (
	"image/color"
	"math/rand"

	"github.com/SimNine/go-city/world/tiles"
	"github.com/SimNine/go-urfutils/src/geom"
	"github.com/SimNine/go-urfutils/src/gfx"
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
	tileMap := make([][]tiles.Tile, dims.Y)
	for i := range dims.Y {
		tileMap[i] = make([]tiles.Tile, dims.X)
	}
	env := &World{
		random: random,

		tiles:          tileMap,
		defaultTileImg: gfx.EbitenCreateHollowRectangleImage(geom.Dims[int]{5, 5}, color.RGBA{255, 0, 0, 255}),
	}

	return env
}

type World struct {
	random *rand.Rand

	tiles          [][]tiles.Tile
	defaultTileImg *ebiten.Image
}

func (w *World) Draw(
	screen *ebiten.Image,
	viewport geom.Viewport[int],
) {
	// Fill the background with blue
	screen.Fill(COLOR_SKYBLUE)

	// For each tile, draw it
	for row := range w.tiles {
		for col := range w.tiles[row] {
			imgOpts := &ebiten.DrawImageOptions{}
			imgOpts.GeoM.Translate(float64(col*6-viewport.Pos.X), float64(row*6-viewport.Pos.Y))
			screen.DrawImage(w.defaultTileImg, imgOpts)
		}
	}
}

func (w *World) Update() {
	// do nothing for now
}
