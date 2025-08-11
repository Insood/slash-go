package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

type AssetManager struct {
}

func LoadAssets() AssetManager {
	err := filepath.WalkDir("assets", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		fmt.Printf("Path: %s, IsDir: %t\n", path, d.IsDir())
		return nil
	})

	if err != nil {
		panic(err)
	}

	return AssetManager{}
}
