package main

import (
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

var fileTree *node

type node struct {
	name  string
	nodes []*node
}

func (n *node) NewRoot(rootName string) {
	if n.name == "" {
		n.name = rootName
	}
}

func (n *node) getOrMakeChildWithName(nodeName string) *node {
	for _, currentNode := range n.nodes {
		if currentNode.name == nodeName {
			return currentNode
		}
	}
	newNode := &node{name: nodeName}
	n.nodes = append(n.nodes, newNode)
	return newNode
}

var cwd string

func scanFolder(root string) error {
	cwd = root
	return filepath.Walk(root, visit)
}

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() && isSupportedType(f.Name()) {
		addFileToTree(path)
	}
	return nil
}

func addFileToTree(f string) {
	nodeName, remainingNodes := shiftNode(strings.TrimPrefix(f, cwd))
	fileTree.NewRoot(nodeName)
	currentNode := fileTree
	for {
		nodeName, remainingNodes = shiftNode(remainingNodes)
		currentNode = currentNode.getOrMakeChildWithName(nodeName)

		if remainingNodes == "" {
			break
		}
	}
}

func shiftNode(filePath string) (string, string) {
	for i, rune := range filePath {
		if os.IsPathSeparator(uint8(rune)) && i > 0 {
			return filePath[:i], filePath[i+1:]
		}
	}
	return filePath, ""
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
