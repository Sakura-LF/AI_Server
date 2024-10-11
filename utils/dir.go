package utils

import (
	"os"
	"path"
)

func GetRootPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for i := 0; i < 2; i++ {
		dir = path.Dir(dir)
	}
	return dir
}
