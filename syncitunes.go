package main

import (
	"fmt"

	"github.com/meonlol/syncitunes/itunes"
)

type itunesInterface interface {
	NewFolder(name string, id int) (int, error)
	NewPlaylist(name string, parentID int) (int, error)
	GetPlaylistIDByName(name string) (int, error)
	//GetPlaylistIDByNameInParent(name string) (int, error)
	GetParentIDForPlaylist(id int) (int, error)
	AddFileToPlaylist(filePath string, playlistID int) (int, error)
	DeletePlaylistByID(id int) error
}

var iTunes itunesInterface

func main() {
	scanFolder("/Users/leonmoll/leon/@music/")
	printTree(fileTree, 0)
	iTunes = &itunes.ApplescriptInterface{}
	fileTreeToItunes(fileTree, false)
}

func fileTreeToItunes(node *node, includeRoot bool) {
	if !includeRoot {
		processNodes(node.nodes, 0)
	} else {
		recurseFileTreeToItunes(node, 0)
	}
}

func recurseFileTreeToItunes(currentNode *node, parentID int) {
	newFolderID, err := iTunes.NewFolder(currentNode.name, parentID)
	logErr(err)

	processNodes(currentNode.nodes, newFolderID)
}

func processNodes(nodes []*node, parentID int) {
	for _, currentChild := range nodes {
		processNodeUnderParent(currentChild, parentID)
	}
}

func processNodeUnderParent(currentChild *node, parentID int) {
	leefs, nodes := separateLeefsAndNodes(currentChild.nodes)

	if len(leefs) > 0 {
		createPlaylist(currentChild, leefs, parentID)
	}
	if len(nodes) > 0 {
		recurseFileTreeToItunes(currentChild, parentID)
	}
}

func createPlaylist(currentChild *node, leefs []*node, parentID int) {
	playlistID, err := iTunes.NewPlaylist(currentChild.name, parentID)
	logErr(err)

	for _, s := range leefs {
		addLeefToPlaylist(s, playlistID)
	}
}

func addLeefToPlaylist(s *node, playlistID int) {
	_, err := iTunes.AddFileToPlaylist(s.path, playlistID)
	logErr(err)
}

func logErr(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
	}
}

func separateLeefsAndNodes(inputNodes []*node) ([]*node, []*node) {
	var leefs []*node
	var nodes []*node
	for _, n := range inputNodes {
		if n.path != "" {
			leefs = append(leefs, n)
		} else {
			nodes = append(nodes, n)
		}
	}
	return leefs, nodes
}
