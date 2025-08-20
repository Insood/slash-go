package main

import (
	"fmt"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	AxisLength = 100
)

func DrawMainScreen(Game *Game) {
	Game.Camera.Zoom += float32(raylib.GetMouseWheelMove()) * ZoomSpeed

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

	for _, tile_map := range Game.AssetManager.Tilemaps {
		for layer_ix, layer := range tile_map.Layers {
			// Only drawing the first layer for now
			if layer_ix > 1 {
				break
			}
			for row := range layer.Rows {
				for col := range layer.Columns {

					tile := layer.GetTile(col, row)

					if tile.TileSetTile == nil {
						continue
					}

					fmt.Println(row, col)
					tex := tile.TileSetTile.TileSet.Texture

					raylib.DrawTexturePro(
						tex,
						tile.TileSetTile.Rect,
						raylib.Rectangle{X: float32(col), Y: float32(row), Width: 1, Height: 1},
						raylib.Vector2{X: 0, Y: 0},
						0,
						raylib.White)
				}
			}
		}
	}

	raylib.DrawRectangleRec(Game.Player, Game.PlayerColor)

	raylib.EndMode2D()

	raylib.EndDrawing()
}
