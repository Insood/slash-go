package main

import (
	"fmt"
	"io/fs"
	"path/filepath"

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

type TileSetTile struct {
	TileMapRect raylib.Rectangle
}

type TileSet struct {
	Width   int
	Height  int
	Columns int
	Tiles   []TileSetTile
	Texture raylib.Texture2D
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
			tile_set.Tiles = append(tile_set.Tiles, TileSetTile{TileMapRect: rect})
		}
	}

	return &tile_set
}

type AssetManager struct {
	Textures map[string]raylib.Texture2D
	Tilesets map[string]TileSet
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

func (asset_manager *AssetManager) LoadTilemap(tilemap_path string) {

}

func LoadAssets() AssetManager {
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

	fmt.Println(paths)

	if err != nil {
		panic(err)
	}

	asset_manager := AssetManager{
		Textures: make(map[string]raylib.Texture2D),
		Tilesets: make(map[string]TileSet),
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

	return asset_manager
}
