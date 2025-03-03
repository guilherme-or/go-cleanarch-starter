package util

import (
	"os"
	"path/filepath"
)

var (
	BasePath = GetBasePath()
)

func GetBasePath() string {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join("/", path)
}