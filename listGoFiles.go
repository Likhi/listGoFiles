package listGoFiles

import (
	"io/fs"
	"path/filepath"
)

var files []string

func FilePathWalkDir(root string) ([]string, error) {
	err := filepath.WalkDir(root, fnOnVisit)
	return files, err
}

func fnOnVisit(path string, di fs.DirEntry, err error) error {
	if !di.IsDir() {
		if filepath.Ext(path) == ".go" {
			files = append(files, path)
		}
	}
	return nil
}
