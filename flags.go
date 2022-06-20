//go:build !js || !wasm
// +build !js !wasm

package main

import (
	"flag"

	"github.com/hajimehoshi/ebiten/v2"
)

func ParseFlags() {
	var (
		fullscreen bool
	)
	flag.BoolVar(&fullscreen, "fullscreen", false, "run in fullscreen mode")
	flag.Parse()

	if fullscreen {
		ebiten.SetFullscreen(true)
	}

}
