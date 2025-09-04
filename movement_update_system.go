package main

import (
	"github.com/mlange-42/ark/ecs"
)

type MovementUpdateSystem struct {
	filter *ecs.Filter2[Position, Velocity]
}

func (s *MovementUpdateSystem) Initialize(w *ecs.World) {
	s.filter = ecs.NewFilter2[Position, Velocity](w)
}

func (s *MovementUpdateSystem) Update(w *ecs.World, dt float32) {
	query := s.filter.Query()

	for query.Next() {
		position, velocity := query.Get()

		position.X += velocity.X * dt
		position.Y += velocity.Y * dt
	}
}
