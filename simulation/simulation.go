package simulation

import (
	"math/rand"
	"strconv"

	"github.com/SimNine/go-city/world"
	"github.com/SimNine/go-urfutils/src/geom"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var TICKS_PER_GENERATION = 1000

func NewSimulation(dims geom.Dims[int]) *Simulation {
	random := rand.New(rand.NewSource(int64(rand.Int())))
	return &Simulation{
		random: random,
		world: world.NewWorld(
			random,
			dims,
		),
	}
}

type Simulation struct {
	random *rand.Rand

	world *world.World

	tickNum int
	Paused  bool
}

func (s *Simulation) Draw(
	screen *ebiten.Image,
	viewport geom.Viewport[int],
) {
	s.world.Draw(screen, viewport)

	// Print out debug info
	printRoot := 10
	ebitenutil.DebugPrintAt(screen, "FPS: "+strconv.FormatFloat(ebiten.ActualFPS(), 'f', 3, 64), 10, printRoot)
	ebitenutil.DebugPrintAt(screen, "TPS: "+strconv.FormatFloat(ebiten.ActualTPS(), 'f', 3, 64), 10, printRoot+15)
}

func (s *Simulation) Update() {
	if !s.Paused {
		s.world.Update()

		s.tickNum++
		if s.tickNum >= TICKS_PER_GENERATION {
			s.tickNum = 0
		}
	}
}
