package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth     = 800
	ScreenHeight    = 600
	CameraPanSpeed  = 0.5
	PlayerMoveSpeed = 0.01
)

var (
	Gray = raylib.NewColor(128, 128, 128, 255)
)

type Game struct {
	Camera      *raylib.Camera2D
	Player      raylib.Rectangle
	PlayerColor raylib.Color
	Texture     *raylib.Texture2D
}

func main() {
	// tileMap := LoadXMLFromFile[TileSetDefinitionXML]("assets/tilesets/floor.tsx")
	// fmt.Println(tileMap)

	// tileMap = LoadXMLFromFile[TileSetDefinitionXML]("assets/tilesets/walls.tsx")
	// fmt.Println(tileMap)

	// gameMap := LoadXMLFromFile[GameMapXML]("assets/maps/main.tmx")
	// fmt.Println(gameMap)

	// LoadAssets()

	raylib.InitWindow(ScreenWidth, ScreenHeight, "Slash")
	defer raylib.CloseWindow()

	texture := raylib.LoadTexture("assets/tiles/walls.png")

	Game := Game{
		Camera:      &raylib.Camera2D{},
		Player:      raylib.Rectangle{X: -1, Y: -1, Width: 1, Height: 1},
		PlayerColor: raylib.Red,
		Texture:     &texture,
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
