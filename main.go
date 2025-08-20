package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth     = 800
	ScreenHeight    = 600
	CameraPanSpeed  = 5
	PlayerMoveSpeed = 0.05
	ZoomSpeed       = 4
)

var (
	Gray = raylib.NewColor(128, 128, 128, 255)
)

type Game struct {
	Camera       *raylib.Camera2D
	Player       raylib.Rectangle
	PlayerColor  raylib.Color
	AssetManager *AssetManager
}

func main() {
	raylib.InitWindow(ScreenWidth, ScreenHeight, "Slash")
	defer raylib.CloseWindow()

	Game := Game{
		Camera:       &raylib.Camera2D{},
		Player:       raylib.Rectangle{X: -1, Y: -1, Width: 1, Height: 1},
		PlayerColor:  raylib.Red,
		AssetManager: LoadAssets(),
	}

	Game.Camera.Target = raylib.Vector2{X: 0, Y: 0}
	Game.Camera.Offset = raylib.Vector2{X: ScreenWidth / 2, Y: ScreenHeight / 2}
	Game.Camera.Rotation = 0
	Game.Camera.Zoom = 128

	raylib.SetTargetFPS(60)

	for !raylib.WindowShouldClose() {
		DrawMainScreen(&Game)
	}
}
