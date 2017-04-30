package filescanner

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/meonlol/syncitunes/tree"
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

var FileTree *tree.Node

var cwd string

func ScanFolder(root string) (*tree.Node, error) {
	dir, _ := filepath.Abs(root)
	cwd = filepath.Dir(dir) + "/"
	FileTree = &tree.Node{}
	err := filepath.Walk(root, visit)
	if err != nil {
		return nil, err
	}
	return FileTree, nil
}

func visit(path string, f os.FileInfo, err error) error {
	if !f.IsDir() && isSupportedType(f.Name()) {
		AddFileToTree(path)
	}
	return nil
}

func AddFileToTree(path string) {
	nodeName, remainingNodes := shiftNode(strings.TrimPrefix(path, cwd))
	FileTree.NewRoot(nodeName)
	currentNode := FileTree
	for {
		nodeName, remainingNodes = shiftNode(remainingNodes)
		currentNode = currentNode.GetOrMakeChildWithName(nodeName)

		if remainingNodes == "" {
			//this is a leaf, add the absolute path to the Node
			currentNode.Path = path
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
