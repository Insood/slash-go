package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	AxisLength = 100
)

func DrawMainScreen(Game *Game) {
	Game.Camera.Zoom += float32(raylib.GetMouseWheelMove()) * 2

	if Game.Camera.Zoom > 128 {
		Game.Camera.Zoom = 128
	} else if Game.Camera.Zoom < 16 {
		Game.Camera.Zoom = 16
	}

	if raylib.IsKeyDown(raylib.KeyLeft) {
		Game.Camera.Offset.X -= CameraPanSpeed
	} else if raylib.IsKeyDown(raylib.KeyRight) {
		Game.Camera.Offset.X += CameraPanSpeed
	}

	if raylib.IsKeyDown(raylib.KeyUp) {
		Game.Camera.Offset.Y -= CameraPanSpeed
	} else if raylib.IsKeyDown(raylib.KeyDown) {
		Game.Camera.Offset.Y += CameraPanSpeed
	}

	if raylib.IsKeyDown(raylib.KeyA) {
		Game.Player.X -= PlayerMoveSpeed
	} else if raylib.IsKeyDown(raylib.KeyD) {
		Game.Player.X += PlayerMoveSpeed
	}

	if raylib.IsKeyDown(raylib.KeyW) {
		Game.Player.Y -= PlayerMoveSpeed
	} else if raylib.IsKeyDown(raylib.KeyS) {
		Game.Player.Y += PlayerMoveSpeed
	}

	raylib.BeginDrawing()
	raylib.ClearBackground(raylib.RayWhite)
	raylib.BeginMode2D(*Game.Camera)

	raylib.DrawLine(0, AxisLength, 0, -AxisLength, raylib.LightGray)
	raylib.DrawLine(AxisLength, 0, -AxisLength, 0, raylib.LightGray)

	raylib.DrawRectangleRec(Game.Player, Game.PlayerColor)
	raylib.DrawTexturePro(
		*Game.Texture,
		// raylib.Rectangle{X: 0, Y: 0, Width: float32(Game.Texture.Width), Height: float32(Game.Texture.Height)},
		raylib.Rectangle{X: 16, Y: 16, Width: 16, Height: 16},
		raylib.Rectangle{X: 0, Y: 0, Width: 1, Height: 1},
		raylib.Vector2{X: 0, Y: 0},
		0,
		raylib.White)

	raylib.EndMode2D()

	// raylib.DrawTexture(*Game.Texture, 0, 0, raylib.White)

	raylib.EndDrawing()
}
