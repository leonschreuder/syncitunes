package main

type itunesInterface interface {
	newFolder(name string, id int) int
	newPlaylist(name string, parentID int) int
	getPlaylistIDByName(name string) (int, error)
	getParentIDForPlaylist(id int) (int, error)
	addFileToPlaylist(filePath string, playlistID int) error
	deletePlaylistByID(id int) error
}

var iTunes itunesInterface
var root = ""

func fileTreeToItunes(node *node) {
	r := iTunes.newFolder(node.name, 0)
	for _, n := range node.nodes {
		r2 := iTunes.newPlaylist(n.name, r)
		// r2 := iTunes.newFolder(n.name, r)
		// iTunes.newPlaylist(n.nodes[0].name, r2)
		iTunes.addFileToPlaylist(n.nodes[0].path, r2)
	}
}
