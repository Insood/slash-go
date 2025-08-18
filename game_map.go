package main

type GameTile struct {
	GameMap *GameMap
}
type GameMap struct {
	Width      int
	Height     int
	Background DrawingLayer
	Foreground DrawingLayer
}

type DrawingLayer struct {
	Width  int
	Height int
	Tiles  []GameTile
}

func LoadGameMap(GameMapXML *GameMapXML) *GameMap {
	game_map := GameMap{
		Width:  GameMapXML.Width,
		Height: GameMapXML.Height,
	}

	return &game_map
}
