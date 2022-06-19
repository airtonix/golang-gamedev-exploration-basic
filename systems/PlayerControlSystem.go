package systems

import (
	"fmt"

	"github.com/airtonix/rpg/components"
	"github.com/sedyh/mizu/pkg/engine"
)

type PlayerControlSystem struct {
}

func NewPlayerControlSystem() *PlayerControlSystem {
	return &PlayerControlSystem{}
}

func (system *PlayerControlSystem) Update(world engine.World) {

	// Get controlled objects
	player, ok := world.View(
		components.Position{},
		components.Velocity{},
		components.Control{},
	).Get()

	if !ok {
		fmt.Println("No entity with Position, Control and Velocity")
		return
	}

	var pos *components.Position
	var vel *components.Velocity
	var control *components.Control

	player.Get(&pos, &vel, &control)

	vel.X = 0.0
	vel.Y = 0.0

	if control.Direction["Right"] {
		vel.X = 1.0
	}

	if control.Direction["Left"] {
		vel.X = -1.0
	}

	if control.Direction["Up"] {
		vel.Y = -1.0
	}

	if control.Direction["Down"] {
		vel.Y = 1.0
	}

	// vel.X = num.Lerp(vel.X, control.MoveSpeed*moveDirectionX, control.MoveSpeed)
	// vel.Y = num.Lerp(vel.Y, control.MoveSpeed*moveDirectionY, control.MoveSpeed)

}
