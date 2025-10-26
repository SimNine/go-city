package tiles

import "github.com/hajimehoshi/ebiten/v2"

type Tile struct {
	T TileType
}

func (t *Tile) GetImage() *ebiten.Image {
	return tileImages[t.T]
}
