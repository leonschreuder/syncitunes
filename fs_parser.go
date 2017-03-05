package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

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

var fileTree node
var cwd string

func scanFolder(root string) {
	cwd = root
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)
}

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() && isSupportedType(f.Name()) {
		addFileToTree(path)
	}
	return nil
}

type node struct {
	name  string
	nodes []*node
}

func addFileToTree(f string) {
	current, rest := popElement(strings.TrimPrefix(f, cwd))
	if fileTree.name != current {
		fileTree = node{name: current}
	}
	currentNode := &fileTree
	for {
		current, rest = popElement(rest)
		var newNode *node
		for _, n := range currentNode.nodes {
			if n.name == current {
				newNode = n
			}
		}

		if newNode == nil {
			newNode = &node{name: current}
			currentNode.nodes = append(currentNode.nodes, newNode)
		}

		currentNode = newNode
		if rest == "" {
			break
		}
	}
}

func popElement(s string) (string, string) {
	splitResult := strings.Split(s, "/")
	current := splitResult[0]
	if len(splitResult) >= 2 {
		var rest string
		if current == "" {
			current = splitResult[1]
			rest = strings.TrimPrefix(s, "/"+current+"/")
		} else {
			rest = strings.TrimPrefix(s, current+"/")
		}
		return current, rest
	}
	return current, ""
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
