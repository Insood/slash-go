package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/mlange-42/ark/ecs"
)

type PlayerMovementInputSystem struct {
	filter *ecs.Filter2[Velocity, PlayerPawn]
}

func (s *PlayerMovementInputSystem) Initialize(w *ecs.World) {
	s.filter = ecs.NewFilter2[Velocity, PlayerPawn](w)
}

func (s *PlayerMovementInputSystem) Update(w *ecs.World, dt float32) {
	query := s.filter.Query()

	for query.Next() {
		velocity, _ := query.Get()

		velocity.X = 0
		velocity.Y = 0

		if raylib.IsKeyDown(raylib.KeyA) {
			velocity.X = -PlayerMoveSpeed
		}
		if raylib.IsKeyDown(raylib.KeyD) {
			velocity.X = PlayerMoveSpeed
		}

		if raylib.IsKeyDown(raylib.KeyW) {
			velocity.Y = -PlayerMoveSpeed
		}
		if raylib.IsKeyDown(raylib.KeyS) {
			velocity.Y = PlayerMoveSpeed
		}
	}
}
