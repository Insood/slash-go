package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

/*
Tiled files reference images, tilesets, etc using relative path (ie: ../tiles/image.png)
We know the absolute path of the file (relative to our executable) that we are reading data
from so can calculate the absolute path of the referenced resource
*/
func AbsolutePathFromRelative(referencing_file_path string, relative_path string) string {
	base := filepath.Dir(referencing_file_path)
	abs := filepath.Join(base, relative_path)
	return abs
}

func LoadTilesetFromXML(TileSetDefinitionXML *TileSetDefinitionXML) *TileSet {
	tile_set := TileSet{}

	tile_width := TileSetDefinitionXML.TileWidth
	tile_height := TileSetDefinitionXML.TileHeight

	columns := TileSetDefinitionXML.Image.ImageWidth / TileSetDefinitionXML.TileWidth
	rows := TileSetDefinitionXML.Image.ImageHeight / TileSetDefinitionXML.TileHeight

	for row := range rows {
		for column := range columns {
			rect := raylib.Rectangle{
				X:      float32(column * tile_width),
				Y:      float32(row * tile_height),
				Width:  float32(tile_width),
				Height: float32(tile_height),
			}
			tile_set.Tiles = append(tile_set.Tiles, TileSetTile{Rect: rect, TileSet: &tile_set})
		}
	}

	return &tile_set
}

type AssetManager struct {
	Textures map[string]raylib.Texture2D
	Tilesets map[string]TileSet
	Tilemaps map[string]TileMap
}

func (asset_manager *AssetManager) LoadTexture(path string) raylib.Texture2D {
	texture := raylib.LoadTexture(path)
	asset_manager.Textures[path] = texture
	return asset_manager.Textures[path]
}

func (asset_manager *AssetManager) LoadTileset(tileset_path string) {
	tile_set_definition := LoadXMLFromFile[TileSetDefinitionXML](tileset_path)
	relative_image_path := tile_set_definition.Image.ImagePath
	tile_set := LoadTilesetFromXML(tile_set_definition)
	absolute_image_path := AbsolutePathFromRelative(tileset_path, relative_image_path)
	texture, ok := asset_manager.Textures[absolute_image_path]
	if !ok {
		panic(fmt.Errorf("could not find texture %s", absolute_image_path))
	}
	tile_set.Texture = texture
	asset_manager.Tilesets[tileset_path] = *tile_set
}

func (asset_manager *AssetManager) GetTileSetTile(tile_map_path string, tile_map_xml *TileMapXML, tile_id int) *TileSetTile {
	// TileSetReferences are in a array with ascending FirstGID values. If each TileSetReference is for
	// an 8x8 tileset (64 tiles), then the FirstGIDs are [1, 65, 129, ...]

	tile_set_tile := &TileSetTile{}
	for _, tile_set_reference := range tile_map_xml.TileSetReferences {
		relative_path := tile_set_reference.Source
		tile_set_path := AbsolutePathFromRelative(tile_map_path, relative_path)
		tile_set := asset_manager.Tilesets[tile_set_path]
		if tile_id < tile_set_reference.FirstGID+len(tile_set.Tiles) {
			tile_set_tile = &tile_set.Tiles[tile_id-tile_set_reference.FirstGID]
			break
		}
	}

	if tile_set_tile == nil {
		panic(fmt.Errorf("could not find tileset tile for tile ID %d", tile_id))
	}
	return tile_set_tile
}

func (asset_manager *AssetManager) LoadTilemap(tile_map_path string) {
	tile_map_xml := LoadXMLFromFile[TileMapXML](tile_map_path)

	offset_x, err_x := tile_map_xml.GetPropertyValueAsInt("offset_x")
	offset_y, err_y := tile_map_xml.GetPropertyValueAsInt("offset_y")
	if (err_x != nil) || (err_y != nil) {
		panic("offset_x and offset_y must be present and be integers")
	}

	tile_map := TileMap{
		Columns: tile_map_xml.Width,
		Rows:    tile_map_xml.Height,
		OffsetX: offset_x,
		OffsetY: offset_y,
		Layers:  make([]DrawingLayer, len(tile_map_xml.Layers)),
	}

	fmt.Println(tile_map_xml.Properties)

	for i, layer := range tile_map_xml.Layers {
		drawing_layer := DrawingLayer{
			Columns: layer.Width,
			Rows:    layer.Height,
			Tiles:   make([]GameTile, layer.Width*layer.Height),
		}
		tile_map.Layers[i] = drawing_layer
		tile_data := strings.TrimSpace(layer.Data.Content)
		tile_data = strings.ReplaceAll(tile_data, "\r\n", "\n")

		for row, row_string := range strings.Split(tile_data, "\n") {
			for column, tile_string := range strings.Split(row_string, ",") {
				tile_string = strings.TrimSpace(tile_string)

				// Otherside of the last comma in the line
				if tile_string == "" {
					continue
				}

				tile_id, err := strconv.Atoi(tile_string)
				if err != nil {
					panic(err)
				}

				var tile_set_tile *TileSetTile
				if tile_id != 0 {
					tile_set_tile = asset_manager.GetTileSetTile(tile_map_path, tile_map_xml, tile_id)
				}

				game_tile := GameTile{
					Row:         row,
					Column:      column,
					Rect:        raylib.Rectangle{X: float32(column), Y: float32(row), Width: 1, Height: 1},
					TileSetTile: tile_set_tile,
				}

				tile_map.Layers[i].SetTile(column, row, &game_tile)
			}
		}
	}

	asset_manager.Tilemaps[tile_map_path] = tile_map
}

func LoadAssets() *AssetManager {
	paths := make(map[string][]string)

	err := filepath.WalkDir("assets", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		extension := filepath.Ext(path)
		switch extension {
		case ".png", ".jpg":
			paths["textures"] = append(paths["textures"], path)
		case ".tsx":
			paths["tilesets"] = append(paths["tilesets"], path)
		case ".tmx":
			paths["tilemaps"] = append(paths["tilemaps"], path)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	asset_manager := AssetManager{
		Textures: make(map[string]raylib.Texture2D),
		Tilesets: make(map[string]TileSet),
		Tilemaps: make(map[string]TileMap),
	}

	for _, path := range paths["textures"] {
		asset_manager.LoadTexture(path)
	}

	for _, path := range paths["tilesets"] {
		asset_manager.LoadTileset(path)
	}

	for _, path := range paths["tilemaps"] {
		asset_manager.LoadTilemap(path)
	}

	return &asset_manager
}
