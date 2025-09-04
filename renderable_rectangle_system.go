package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/mlange-42/ark/ecs"
)

type RenderableRectangleSystem struct {
	filter *ecs.Filter2[Position, RenderableRectangle]
}

func (s *RenderableRectangleSystem) Initialize(w *ecs.World) {
	s.filter = ecs.NewFilter2[Position, RenderableRectangle](w)
}

func (s *RenderableRectangleSystem) Update(w *ecs.World, _ float32) {
	query := s.filter.Query()

	for query.Next() {
		position, renderable_rect := query.Get()
		rect := raylib.Rectangle{X: position.X, Y: position.Y, Width: renderable_rect.Width, Height: renderable_rect.Height}
		raylib.DrawRectangleRec(rect, renderable_rect.Color)
	}
}
