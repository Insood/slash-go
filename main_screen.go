package main

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

const (
	AxisLength = 100
)

func DrawLayer(layer *DrawingLayer, offset_x int, offset_y int) {
	for row := range layer.Rows {
		for col := range layer.Columns {

			tile := layer.GetTile(col, row)

			if tile.TileSetTile == nil {
				continue
			}

			tex := tile.TileSetTile.TileSet.Texture

			raylib.DrawTexturePro(
				tex,
				tile.TileSetTile.Rect,
				raylib.Rectangle{X: float32(col + offset_x), Y: float32(row + offset_y), Width: 1, Height: 1},
				raylib.Vector2{X: 0, Y: 0},
				0,
				raylib.White)
		}
	}
}

func DrawLayers(asset_manager *AssetManager) {
	for _, tile_map := range asset_manager.Tilemaps {
		for _, layer := range tile_map.Layers {
			DrawLayer(&layer, tile_map.OffsetX, tile_map.OffsetY)
		}
	}
}

func DrawTileMapOutline(asset_manager *AssetManager) {
	for _, tile_map := range asset_manager.Tilemaps {
		raylib.DrawRectangleLinesEx(
			raylib.Rectangle{
				X:      float32(tile_map.OffsetX),
				Y:      float32(tile_map.OffsetY),
				Width:  float32(tile_map.Columns),
				Height: float32(tile_map.Rows),
			},
			1.0/16,
			raylib.Green)
	}
}

func DrawCollisions(asset_manager *AssetManager) {
	for _, tile_map := range asset_manager.Tilemaps {
		for _, layer := range tile_map.Layers {
			for row := range layer.Rows {
				for col := range layer.Columns {
					tile := layer.GetTile(col, row)
					if tile.TileSetTile == nil {
						continue
					}

					collision_rect := tile.TileSetTile.CollisionRect

					if collision_rect.Width > 0 && collision_rect.Height > 0 {
						rect := raylib.Rectangle{
							X:      float32(col+tile_map.OffsetX) + collision_rect.X,
							Y:      float32(row+tile_map.OffsetY) + collision_rect.Y,
							Width:  collision_rect.Width,
							Height: collision_rect.Height,
						}

						raylib.DrawRectangleLinesEx(rect, 1.0/16, raylib.Red)
					}
				}
			}
		}
	}
}

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
	raylib.ClearBackground(raylib.Black)
	raylib.BeginMode2D(*Game.Camera)

	raylib.DrawLine(0, AxisLength, 0, -AxisLength, raylib.LightGray)
	raylib.DrawLine(AxisLength, 0, -AxisLength, 0, raylib.LightGray)

	DrawLayers(Game.AssetManager)
	DrawCollisions(Game.AssetManager)
	DrawTileMapOutline(Game.AssetManager)

	raylib.DrawRectangleRec(Game.Player, Game.PlayerColor)

	raylib.EndMode2D()

	raylib.EndDrawing()
}
