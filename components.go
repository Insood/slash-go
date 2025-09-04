package main

import raylib "github.com/gen2brain/raylib-go/raylib"

type Position struct {
	X float32
	Y float32
}

type Velocity struct {
	X float32
	Y float32
}

type RenderableRectangle struct {
	Width  float32
	Height float32
	Color  raylib.Color
}

type PlayerPawn struct{}
