package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var filesList []string
var supportedFileTypes = []string{
	".mp3",  //mp3 types
	".aiff", //aiff types
	".aif",
	".aifc",
	".wav", //wave types
	".wave",
	".m4a", //mpeg-4 types
	".mp4",
	".3gp",
	".m4b", //aac types
	".m4p",
	".m4r",
	".m4v",
	".aac",
	".caf", //apple lossless
}

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() && isSupportedType(f.Name()) {
		filesList = append(filesList, path)
	}
	return nil
}

func isSupportedType(fileName string) bool {
	extension := filepath.Ext(fileName)
	for _, ft := range supportedFileTypes {
		if extension == ft {
			return true
		}
	}
	return false
}

func findMusic(root string) {
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}
