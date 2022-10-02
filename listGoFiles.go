package listGoFiles

import (
	"io/fs"
	"path/filepath"
)

var files []string

func FilePathWalkDir(root string) ([]string, error) {
	err := filepath.WalkDir(root, visit)
	return files, err
}

func visit(path string, di fs.DirEntry, err error) error {
	if !di.IsDir() {
		if filepath.Ext(path) == ".go" {
			// fmt.Println(path)
			files = append(files, path)
		}
	}
	return nil
}
