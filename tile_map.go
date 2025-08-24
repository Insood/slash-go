package main

import raylib "github.com/gen2brain/raylib-go/raylib"

type GameTile struct {
	Rect        raylib.Rectangle
	Row         int
	Column      int
	TileSetTile *TileSetTile
}

type TileMap struct {
	Columns int
	Rows    int
	OffsetX int
	OffsetY int
	Layers  []DrawingLayer
}

type DrawingLayer struct {
	Columns int
	Rows    int
	Tiles   []GameTile
}

func (drawing_layer *DrawingLayer) GetTile(column int, row int) *GameTile {
	return &drawing_layer.Tiles[row*drawing_layer.Columns+column]
}

func (drawing_layer *DrawingLayer) SetTile(column int, row int, tile *GameTile) {
	drawing_layer.Tiles[row*drawing_layer.Columns+column] = *tile
}

type TileSetTile struct {
	Rect          raylib.Rectangle
	TileSet       *TileSet
	CollisionRect raylib.Rectangle
}

type TileSet struct {
	Width   int
	Height  int
	Columns int
	Tiles   []TileSetTile
	Texture raylib.Texture2D
}
