package main

import raylib "github.com/gen2brain/raylib-go/raylib"

func DrawSplashScreen() {
	raylib.BeginDrawing()
	raylib.ClearBackground(raylib.RayWhite)

	startText := "Press Space to Start"
	raylib.DrawText(startText, ScreenWidth/2, 255, 20, Gray)

	raylib.EndDrawing()
}
