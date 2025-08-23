package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"
)

type TileSetReferenceXML struct {
	FirstGID int    `xml:"firstgid,attr"`
	Source   string `xml:"source,attr"`
}

type LayerXML struct {
	ID     int    `xml:"id,attr"`
	Name   string `xml:"name,attr"`
	Width  int    `xml:"width,attr"`
	Height int    `xml:"height,attr"`
	Data   struct {
		Encoding string `xml:"encoding,attr"`
		Content  string `xml:",chardata"`
	} `xml:"data"`
}

type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type TileMapXML struct {
	// <map version="1.10" tiledversion="1.11.2" orientation="orthogonal" renderorder="right-down" width="30" height="20" tilewidth="16" tileheight="16" infinite="0" nextlayerid="3" nextobjectid="1">
	Version           string                `xml:"version,attr"`
	TiledVersion      string                `xml:"tiledversion,attr"`
	Orientation       string                `xml:"orientation,attr"`
	RenderOrder       string                `xml:"renderorder,attr"`
	Width             int                   `xml:"width,attr"`
	Height            int                   `xml:"height,attr"`
	TileWidth         int                   `xml:"tilewidth,attr"`
	TileHeight        int                   `xml:"tileheight,attr"`
	Infinite          int                   `xml:"infinite,attr"`
	NextLayerID       int                   `xml:"nextlayerid,attr"`
	NextObjectID      int                   `xml:"nextobjectid,attr"`
	TileSetReferences []TileSetReferenceXML `xml:"tileset"`
	Properties        []Property            `xml:"properties>property"`
	Layers            []LayerXML            `xml:"layer"`
}

func (tile_map_xml *TileMapXML) GetPropertyValue(name string) string {
	for _, property := range tile_map_xml.Properties {
		if property.Name == name {
			return property.Value
		}
	}
	return ""
}

func (tile_map_xml *TileMapXML) GetPropertyValueAsInt(name string) (int, error) {
	raw_value := tile_map_xml.GetPropertyValue(name)
	value, err := strconv.Atoi(raw_value)
	return value, err
}

type TileSetDefinitionXML struct {
	Name         string `xml:"name,attr"`
	Version      string `xml:"version,attr"`
	TiledVersion string `xml:"tiledversion,attr"`
	TileWidth    int    `xml:"tilewidth,attr"`
	TileHeight   int    `xml:"tileheight,attr"`
	TileCount    int    `xml:"tilecount,attr"`
	Columns      int    `xml:"columns,attr"`
	Image        struct {
		ImagePath   string `xml:"source,attr"`
		ImageWidth  int    `xml:"width,attr"`
		ImageHeight int    `xml:"height,attr"`
	} `xml:"image"`
}

func LoadXMLFromFile[XMLType any](path string) *XMLType {
	xmlFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer xmlFile.Close()
	xmlStruct := new(XMLType)

	byteValue, _ := io.ReadAll(xmlFile)

	err = xml.Unmarshal(byteValue, &xmlStruct)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return xmlStruct
}
