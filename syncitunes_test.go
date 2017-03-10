package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var count = 0

func Test__tree_with_single_file_to_itunes(t *testing.T) {
	fileTree = &node{}
	cwd = ""
	mock := mockInterface{}
	iTunes = &mock
	resultNode = &mockNode{}

	addFileToTree("root/some_album/song.mp3")

	fileTreeToItunes(fileTree)

	assert.Equal(t, "root", resultNode.name)
	assert.EqualValues(t, d, resultNode.kind)
	assert.Equal(t, 1, len(resultNode.mockNodes))
	assert.Equal(t, "some_album", resultNode.mockNodes[0].name)
	assert.EqualValues(t, p, resultNode.mockNodes[0].kind)
	assert.Equal(t, 1, len(resultNode.mockNodes[0].mockNodes))
	assert.EqualValues(t, f, resultNode.mockNodes[0].mockNodes[0].kind)
	assert.Equal(t, "root/some_album/song.mp3", resultNode.mockNodes[0].mockNodes[0].name)
}

func Test__tree_with_multiple_files_to_itunes(t *testing.T) {
	fileTree = &node{}
	cwd = ""
	mock := mockInterface{}
	iTunes = &mock
	resultNode = &mockNode{}

	addFileToTree("root/some_album/song.mp3")
	addFileToTree("root/some_other_album/song.mp3")

	fileTreeToItunes(fileTree)

	assert.Equal(t, "root", resultNode.name)
	assert.EqualValues(t, d, resultNode.kind)
	assert.Equal(t, 2, len(resultNode.mockNodes))
	assert.EqualValues(t, p, resultNode.mockNodes[0].kind)
	assert.Equal(t, "some_album", resultNode.mockNodes[0].name)
	assert.EqualValues(t, p, resultNode.mockNodes[1].kind)
	assert.Equal(t, "some_other_album", resultNode.mockNodes[1].name)
	assert.Equal(t, 1, len(resultNode.mockNodes[0].mockNodes))
	assert.EqualValues(t, f, resultNode.mockNodes[0].mockNodes[0].kind)
	assert.Equal(t, "root/some_album/song.mp3", resultNode.mockNodes[0].mockNodes[0].name)
	assert.EqualValues(t, f, resultNode.mockNodes[1].mockNodes[0].kind)
	assert.Equal(t, "root/some_other_album/song.mp3", resultNode.mockNodes[1].mockNodes[0].name)
}

var resultNode *mockNode

type itemType int

type mockNode struct {
	name      string
	id        int
	parent    int
	kind      itemType
	mockNodes []*mockNode
}

const (
	f = iota //file
	d        //dir
	p        //playlist
)

func addNode(name string, t itemType, parent int) int {
	count++
	newNode := &mockNode{name: name, id: count, kind: t}
	if resultNode.name == "" {
		resultNode.name = newNode.name
		resultNode.id = newNode.id
		resultNode.kind = newNode.kind
	} else {
		if resultNode.id == parent {
			resultNode.mockNodes = append(resultNode.mockNodes, newNode)
		} else {
			for _, n := range resultNode.mockNodes {
				if n.id == parent {
					n.mockNodes = append(n.mockNodes, newNode)
				}
			}
		}
	}
	return count
}

type mockInterface struct {
	currentID   int
	pathCreated string
}

func (i *mockInterface) newFolder(name string, id int) int {
	// newNode := &mockNode{name: name, id: id + 1, kind: "f"}
	return addNode(name, d, id)
	// return id + 1
}
func (i *mockInterface) newPlaylist(name string, id int) int {
	// newNode := &mockNode{name: name, id: id + 1}
	// addNode(newNode, id)
	// return newNode.id
	return addNode(name, p, id)
	// return id + 1
}
func (mockInterface) getPlaylistIDByName(name string) (int, error) {
	return -1, nil
}
func (mockInterface) getParentIDForPlaylist(id int) (int, error) {
	return -1, nil
}
func (i *mockInterface) addFileToPlaylist(filePath string, playlistID int) error {
	addNode(filePath, f, playlistID)
	return nil
}

func (mockInterface) deletePlaylistByID(id int) error {
	return nil
}
