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

func fileTreeToItunes(node *node) {
	recurseFileTreeToItunes(node, 0)
}

func recurseFileTreeToItunes(currentNode *node, lastParentID int) {
	parentID := iTunes.newFolder(currentNode.name, lastParentID)
	for _, n := range currentNode.nodes {
		leefs, nodes := splitInLeefsAndNodes(n.nodes)

		if len(leefs) > 0 {
			r2 := iTunes.newPlaylist(n.name, parentID)
			for _, s := range leefs {
				iTunes.addFileToPlaylist(s.path, r2)
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
