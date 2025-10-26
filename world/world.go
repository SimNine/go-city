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

var COLOR_BLACK = color.RGBA{
	R: 0,
	G: 0,
	B: 0,
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
		for j := range dims.X {
			tileMap[i][j] = tiles.Tile{T: tiles.WATER}
		}
	}
	env := &World{
		random: random,

		tiles: tileMap,
	}

	return env
}

type World struct {
	random *rand.Rand

	tiles [][]tiles.Tile
}

func (w *World) Draw(
	screen *ebiten.Image,
	viewport geom.Viewport[int],
) {
	// Fill the background with blue
	screen.Fill(COLOR_BLACK)

	// For each tile, draw it
	for row := range w.tiles {
		for col := range w.tiles[row] {
			image := w.tiles[row][col].GetImage()
			imgOpts := &ebiten.DrawImageOptions{}
			imgOpts.GeoM.Translate(float64(col*(image.Bounds().Dx()+2)-viewport.Pos.X), float64(row*(image.Bounds().Dy()+2)-viewport.Pos.Y))
			screen.DrawImage(image, imgOpts)
		}
	}
}

func (w *World) Update() {
	// do nothing for now
}
