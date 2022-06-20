package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/airtonix/rpg/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/sedyh/mizu/pkg/engine"
)

func main() {

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowSizeLimits(300, 200, -1, -1)
	ebiten.SetFPSMode(ebiten.FPSModeVsyncOffMaximum)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	rand.Seed(time.Now().UTC().UnixNano())

	g := engine.NewGame(&game.Game{})
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
