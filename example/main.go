package main

import (
	"fmt"
	lgf "listGoFiles"
	"log"
	"os"
)

func main() {
	var root string
	if len(os.Args) == 1 {
		log.Fatal("No path given, Please specify path.")
		return
	}
	if root = os.Args[1]; root == "" {
		log.Fatal("No path given, Please specify path.")
		return
	}

	files, _ := lgf.FilePathWalkDir(root)

	for _, file := range files {
		fmt.Println(file)
	}
}
