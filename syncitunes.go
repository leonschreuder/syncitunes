package main

import (
	"fmt"

	"github.com/meonlol/syncitunes/itunes"
)

type itunesInterface interface {
	NewFolder(name string, id int) (int, error)
	NewPlaylist(name string, parentID int) (int, error)
	GetPlaylistIDByName(name string) (int, error)
	GetParentIDForPlaylist(id int) (int, error)
	AddFileToPlaylist(filePath string, playlistID int) (int, error)
	DeletePlaylistByID(id int) error
}

var iTunes itunesInterface

func main() {
	scanFolder("/Users/leonmoll/leon/@music/")
	printTree(fileTree, 0)
	iTunes = &itunes.ApplescriptInterface{}
	fileTreeToItunes(fileTree)
}

func fileTreeToItunes(node *node) {
	recurseFileTreeToItunes(node, 0)
}

func recurseFileTreeToItunes(currentNode *node, lastParentID int) {
	parentID, err := iTunes.NewFolder(currentNode.name, lastParentID)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
	}
	for _, n := range currentNode.nodes {
		leefs, nodes := splitInLeefsAndNodes(n.nodes)

		if len(leefs) > 0 {
			r2, err := iTunes.NewPlaylist(n.name, parentID)
			if err != nil {
				fmt.Println("ERROR: ", err.Error())
			}
			for _, s := range leefs {
				_, err := iTunes.AddFileToPlaylist(s.path, r2)
				if err != nil {
					fmt.Println("ERROR: ", err.Error())
				}
			}
		}
		if len(nodes) > 0 {
			recurseFileTreeToItunes(n, parentID)
		}
	}
}

func splitInLeefsAndNodes(inputNodes []*node) ([]*node, []*node) {
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
