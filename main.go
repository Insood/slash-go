package main

import (
	"fmt"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

var (
	Background = raylib.NewColor(255, 255, 255, 255)
	Gray       = raylib.NewColor(128, 128, 128, 255)
)

func main() {
	tileMap := LoadXMLFromFile[TileSetDefinition]("assets/tilesets/floor.tsx")
	fmt.Println(tileMap)

	tileMap = LoadXMLFromFile[TileSetDefinition]("assets/tilesets/walls.tsx")
	fmt.Println(tileMap)

	gameMap := LoadXMLFromFile[GameMap]("assets/maps/main.tmx")
	fmt.Println(gameMap)

	// raylib.InitWindow(ScreenWidth, ScreenHeight, "Slash")
	// defer raylib.CloseWindow()

	// raylib.SetTargetFPS(60)

	// for !raylib.WindowShouldClose() {
	// 	DrawSplashScreen()
	// }
}
