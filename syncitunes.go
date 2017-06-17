package main

import (
	"fmt"

	"github.com/meonlol/syncitunes/filescanner"
	"github.com/meonlol/syncitunes/itunes"
	"github.com/meonlol/syncitunes/itunes/applescript_cli"
	"github.com/meonlol/syncitunes/tree"
)

var iTunes itunes.Interface

var nodeTree *tree.Node

// Build GUI with this? https://github.com/andlabs/ui

func main() {
	nodeTree, _ = filescanner.ScanFolder("/Users/leonmoll/leon/@music/")
	tree.Print(nodeTree, 0)
	iTunes = &applescript_cli.Adapter{}
	iTunes.UpdateTreeWithExisting(nodeTree)
	fileTreeToItunes(nodeTree, false)
}

func fileTreeToItunes(node *tree.Node, includeRoot bool) {
	if !includeRoot {
		processNodes(node.Nodes, 0)
	} else {
		recurseFileTreeToItunes(node, 0)
	}
}

func recurseFileTreeToItunes(currentNode *tree.Node, parentID int) {
	newFolderID, err := iTunes.NewFolder(currentNode.Name, parentID)
	logErr(err)

	processNodes(currentNode.Nodes, newFolderID)
}

func processNodes(nodes []*tree.Node, parentID int) {
	for _, currentChild := range nodes {
		processNodeUnderParent(currentChild, parentID)
	}
}

func processNodeUnderParent(currentChild *tree.Node, parentID int) {
	leefs, nodes := separateLeefsAndNodes(currentChild.Nodes)

	if len(leefs) > 0 {
		createPlaylist(currentChild, leefs, parentID)
	}
	if len(nodes) > 0 {
		recurseFileTreeToItunes(currentChild, parentID)
	}
}

func createPlaylist(currentChild *tree.Node, leefs []*tree.Node, parentID int) {
	playlistID, err := iTunes.NewPlaylist(currentChild.Name, parentID)
	logErr(err)

	for _, s := range leefs {
		addLeefToPlaylist(s, playlistID)
	}
}

func addLeefToPlaylist(s *tree.Node, playlistID int) {
	_, err := iTunes.AddFileToPlaylist(s.Path, playlistID)
	logErr(err)
}

func logErr(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
	}
}

func separateLeefsAndNodes(inputNodes []*tree.Node) ([]*tree.Node, []*tree.Node) {
	var leefs []*tree.Node
	var nodes []*tree.Node
	for _, n := range inputNodes {
		if n.Path != "" {
			leefs = append(leefs, n)
		} else {
			nodes = append(nodes, n)
		}
	}
	return leefs, nodes
}
