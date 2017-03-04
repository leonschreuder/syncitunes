package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test__sould_be_able_to_create_playlist_with_file(t *testing.T) {
	mock := mockInterface{}
	iTunes = &mock

	wd, _ := os.Getwd()
	root = wd
	absFile := wd + "/itunes-boundry/empty.mp3"
	id := createPlaylist(absFile)

	assert.Equal(t, "1337", id)
	assert.Equal(t, "itunes-boundry", mock.playlistName)
	assert.Equal(t, absFile, mock.file)
}

func Test__sould_be_able_to_create_playlist_in_folder(t *testing.T) {
	mock := mockInterface{}
	iTunes = &mock

	wd, _ := os.Getwd()
	root = wd
	absFile := wd + "/someFolder/itunes-boundry/empty.mp3"
	createPlaylist(absFile)

	assert.Equal(t, "itunes-boundry", mock.playlistName)
	assert.Equal(t, "someFolder", mock.folderCreated)
	assert.Equal(t, absFile, mock.file)
}

func Test__should_add_single_node(t *testing.T) {
	treeFound = neoNode{}

	addToTree("1.mp3")

	assert.Equal(t, "1.mp3", treeFound.name)
}

func Test__should_add_sub_node(t *testing.T) {
	treeFound = neoNode{}

	addToTree("root/1.mp3")

	assert.Equal(t, "root", treeFound.name)
	assert.Equal(t, 1, len(treeFound.nodes))
	assert.Equal(t, "1.mp3", treeFound.nodes[0].name)
}

func Test__should_add_second_sub_node(t *testing.T) {
	treeFound = neoNode{}

	addToTree("root/1.mp3")
	addToTree("root/2.mp3")

	assert.Equal(t, "root", treeFound.name)
	assert.Equal(t, 2, len(treeFound.nodes))
	assert.Equal(t, "1.mp3", treeFound.nodes[0].name)
	assert.Equal(t, "2.mp3", treeFound.nodes[1].name)
}

func Test__should_add_second_level_node(t *testing.T) {
	treeFound = neoNode{}

	addToTree("root/subFolder/1.mp3")

	assert.Equal(t, "root", treeFound.name)
	assert.Len(t, treeFound.nodes, 1)
	assert.Equal(t, "subFolder", treeFound.nodes[0].name)
	assert.Len(t, treeFound.nodes[0].nodes, 1)
}

func Test__should_add_multiple_second_level_nodes(t *testing.T) {
	treeFound = neoNode{}

	addToTree("root/subFolder/1.mp3")
	addToTree("root/subFolder/2.mp3")

	assert.Equal(t, "root", treeFound.name)
	assert.Equal(t, 1, len(treeFound.nodes))
	assert.Equal(t, "subFolder", treeFound.nodes[0].name)
	assert.Equal(t, 2, len(treeFound.nodes[0].nodes))
}

func Test__should_add_multiple_multilevel_nodes(t *testing.T) {
	treeFound = neoNode{}

	addToTree("root/subFolder/1.mp3")
	addToTree("root/subFolder/2.mp3")
	addToTree("root/subFolder2/3.mp3")

	assert.Equal(t, "root", treeFound.name)
	assert.Equal(t, 2, len(treeFound.nodes))
	assert.Equal(t, "subFolder", treeFound.nodes[0].name)
	assert.Equal(t, "subFolder2", treeFound.nodes[1].name)
	assert.Equal(t, "1.mp3", treeFound.nodes[0].nodes[0].name)
	assert.Equal(t, "2.mp3", treeFound.nodes[0].nodes[1].name)
	assert.Equal(t, "3.mp3", treeFound.nodes[1].nodes[0].name)
}

func Test__should_add_realistic_folder_structure(t *testing.T) {
	treeFound = neoNode{}

	addToTree("root/Aphrodite/Aphrodite - Urban Jungle/1.mp3")
	addToTree("root/Aphrodite/Aphrodite - Urban Jungle/2.mp3")
	addToTree("root/Arctic Monkeys/Arctic Monkeys - Artic Monkeys (Album)/1.mp3")
	addToTree("root/Arctic Monkeys/Arctic Monkeys - Artic Monkeys (Album)/2.mp3")
	addToTree("root/Arctic Monkeys/Arctic Monkeys - Artic Monkeys (Album)/3.mp3")
	addToTree("root/Arctic Monkeys/Arctic Monkeys - Favourite Worst Nightmare/1.mp3")
	addToTree("root/Arctic Monkeys/Arctic Monkeys - Favourite Worst Nightmare/2.mp3")

	assert.Equal(t, "root", treeFound.name)
	assert.Equal(t, 2, len(treeFound.nodes))
	assert.Equal(t, "Aphrodite", treeFound.nodes[0].name)
	assert.Equal(t, "Arctic Monkeys", treeFound.nodes[1].name)
	assert.Equal(t, 1, len(treeFound.nodes[0].nodes))
	assert.Equal(t, "Aphrodite - Urban Jungle", treeFound.nodes[0].nodes[0].name)
	assert.Equal(t, 2, len(treeFound.nodes[0].nodes[0].nodes))
	assert.Equal(t, 2, len(treeFound.nodes[1].nodes))
	assert.Equal(t, "Arctic Monkeys - Artic Monkeys (Album)", treeFound.nodes[1].nodes[0].name)
	assert.Equal(t, 3, len(treeFound.nodes[1].nodes[0].nodes))
}

func Test__should_convert_single_file_into_tree(t *testing.T) {

	// tree := convertToTree("root/one/first/1.mp3")
	tree := growTree("root/one/first/1.mp3")

	assert.Equal(t, "root", tree.name)
	assert.Equal(t, 1, len(tree.nodes))
	assert.Equal(t, "one", tree.nodes[0].name)
	assert.Equal(t, 1, len(tree.nodes[0].nodes))
	assert.Equal(t, "first", tree.nodes[0].nodes[0].name)
	assert.Equal(t, 1, len(tree.nodes[0].nodes[0].nodes))
	assert.Equal(t, "1.mp3", tree.nodes[0].nodes[0].nodes[0].name)
}

// func Test__should_convert_multiple_files_into_tree(t *testing.T) {

// 	tree := convertToTree("root/one/first/1.mp3", "root/second/2.mp3")

// 	fmt.Println(tree)
// 	assert.Equal(t, "root", tree.name)
// 	assert.Equal(t, 2, len(tree.nodes))
// 	assert.Equal(t, "one", tree.nodes[0].name)
// 	assert.Equal(t, "second", tree.nodes[1].name)
// }

func Test__should_join_on_first_level(t *testing.T) {
	tree1 := node{"root", []node{node{"one", []node{node{name: "1"}}}}}
	tree2 := node{"root", []node{node{name: "two"}}}

	finalTree := joinTree(tree1, tree2)

	assert.Equal(t, "root", finalTree.name)
	assert.Equal(t, 2, len(finalTree.nodes))
	assert.Equal(t, 1, len(finalTree.nodes[0].nodes))
}

// func Test__should_join_on_second_level(t *testing.T) {
// 	tree1 := node{"root", []node{node{"one", []node{node{name: "1"}}}}}
// 	tree2 := node{"root", []node{node{"one", []node{node{name: "2"}}}}}

// 	finalTree := joinTree(tree1, tree2)

// 	assert.Equal(t, "root", finalTree.name)
// 	assert.Equal(t, 1, len(finalTree.nodes))
// 	// assert.Equal(t, 2, len(finalTree.nodes[0].nodes))
// }

type mockInterface struct {
	file          string
	playlistName  string
	folderCreated string
}

func (i *mockInterface) newFolder(name string) string {
	i.folderCreated = name
	return ""
}
func (i *mockInterface) newPlaylist(name, parentID string) string {
	i.playlistName = name
	return "1337"
}
func (mockInterface) getPlaylistIDByName(name string) (string, error) {
	return "", nil
}
func (mockInterface) getParentIDForPlaylist(id string) (string, error) {
	return "", nil
}
func (i *mockInterface) addFileToPlaylist(filePath, playlistID string) error {
	i.file = filePath
	return nil
}

func (mockInterface) deletePlaylistByID(id string) error {
	return nil
}
