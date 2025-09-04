package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
	"github.com/mlange-42/ark/ecs"
)

type Engine struct {
	World   ecs.World
	Systems []System
}

type System interface {
	Initialize(*ecs.World)
	Update(*ecs.World, float32)
}

func (engine *Engine) AddSystem(system System) {
	engine.Systems = append(engine.Systems, system)
}

func NewEngine() *Engine {
	engine := Engine{
		World: ecs.NewWorld(),
	}

	engine.AddSystem(&PlayerMovementInputSystem{})
	engine.AddSystem(&MovementUpdateSystem{})
	engine.AddSystem(&RenderableRectangleSystem{})

	for _, system := range engine.Systems {
		system.Initialize(&engine.World)
	}

	mapper := ecs.NewMap4[Position, Velocity, RenderableRectangle, PlayerPawn](&engine.World)
	mapper.NewEntity(
		&Position{X: 0, Y: 0},
		&Velocity{X: 0, Y: 0},
		&RenderableRectangle{Width: 1, Height: 1, Color: raylib.Red},
		&PlayerPawn{},
	)

	return &engine
}

func (e *Engine) Update(dt float32) {
	for _, system := range e.Systems {
		system.Update(&e.World, dt)
	}
}
