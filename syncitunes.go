package main

import "strings"

type itunesInterface interface {
	newFolder(name string) string
	newPlaylist(name, parentID string) string
	getPlaylistIDByName(name string) (string, error)
	getParentIDForPlaylist(id string) (string, error)
	addFileToPlaylist(filePath, playlistID string) error
	deletePlaylistByID(id string) error
}

var iTunes itunesInterface
var root = ""

func createPlaylist(file string) string {
	parentID := iTunes.newFolder("someFolder")
	id := iTunes.newPlaylist("itunes-boundry", parentID)
	iTunes.addFileToPlaylist(file, id)
	return id
}

var treeFound neoNode

// var currentNode *neoNode

type neoNode struct {
	name  string
	nodes []*neoNode
}

// func visit(path string, f os.FileInfo, err error) error {
// 	if currentNode == nil {
// 		treeFound = neoNode{name: f.Name()}
// 		currentNode = &treeFound
// 	} else {
// 		newNode := &neoNode{name: f.Name()}
// 		currentNode.nodes = append(currentNode.nodes, newNode)
// 		currentNode = newNode
// 	}
// 	return nil
// }

type node struct {
	name  string
	nodes []node
}

func addToTree(f string) {
	current, rest := popElement(f)
	if treeFound.name != current {
		treeFound = neoNode{name: current}
	}
	currentNode := &treeFound
	for {
		current, rest = popElement(rest)
		var newNode *neoNode
		for _, n := range currentNode.nodes {
			if n.name == current {
				newNode = n
			}
		}

		if newNode == nil {
			newNode = &neoNode{name: current}
			currentNode.nodes = append(currentNode.nodes, newNode)
		}

		currentNode = newNode
		if rest == "" {
			break
		}
	}
}

func convertToTree(files ...string) node {
	var finalTree node
	for _, singleFile := range files {
		currentTree := growTree(singleFile)
		finalTree = currentTree
	}
	return finalTree
}

func growTree(path string) node {
	current, rest := popElement(path)
	currentNode := node{name: current}
	if rest != "" {
		currentNode.nodes = []node{growTree(rest)}
	}
	return currentNode
}

func joinTree(tree1, tree2 node) node {
	tree1.nodes = append(tree1.nodes, tree2.nodes...)
	return tree1
}

func popElement(s string) (string, string) {
	// filepath.SplitList(s)
	splitResult := strings.Split(s, "/")
	if len(splitResult) >= 2 {
		rest := strings.TrimPrefix(s, splitResult[0]+"/")
		return splitResult[0], rest
	}
	return splitResult[0], ""
}
