package tiles

import (
	urfutils "github.com/SimNine/go-urfutils/src"
	"github.com/hajimehoshi/ebiten/v2"
)

type TileType int

const (
	GRASS TileType = iota
	WATER
	SAND
)

var tileImages map[TileType]*ebiten.Image

func initSingleTileImage(tileType TileType, filePath string) {
	img, err := urfutils.LoadEbitenImageFromFile(filePath)
	if err != nil {
		panic(err)
	}
	tileImages[tileType] = img
}

func InitTileImages() {
	tileImages = make(map[TileType]*ebiten.Image)

	initSingleTileImage(GRASS, "assets/tiles/grass.png")
	initSingleTileImage(WATER, "assets/tiles/water.png")
	initSingleTileImage(SAND, "assets/tiles/sand.png")
}
