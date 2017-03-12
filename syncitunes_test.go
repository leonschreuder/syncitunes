package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var count = 0

func setup() {
	fileTree = &node{}
	cwd = ""
	mock := mockInterface{}
	iTunes = &mock
	resultNode = &mockNode{}
	count = 0
}

func Test__tree_with_single_file_to_itunes(t *testing.T) {
	setup()

	addFileToTree("root/some_album/song.mp3")

	fileTreeToItunes(fileTree)

	assertNodeNameAndType(t, "root", d, resultNode)
	assertNodeChildCountEquals(t, 1, resultNode)
	assertNodeNameAndType(t, "some_album", p, resultNode.mockNodes[0])
	assertNodeChildCountEquals(t, 1, resultNode.mockNodes[0])
	assertNodeNameAndType(t, "root/some_album/song.mp3", f, resultNode.mockNodes[0].mockNodes[0])
}

func Test__tree_with_multiple_files_to_itunes(t *testing.T) {
	setup()

	addFileToTree("root/some_album/song.mp3")
	addFileToTree("root/some_other_album/song.mp3")

	fileTreeToItunes(fileTree)

	assertNodeNameAndType(t, "root", d, resultNode)
	assertNodeChildCountEquals(t, 2, resultNode)
	assertNodeNameAndType(t, "some_album", p, resultNode.mockNodes[0])
	assertNodeNameAndType(t, "some_other_album", p, resultNode.mockNodes[1])
	assertNodeChildCountEquals(t, 1, resultNode.mockNodes[0])
	assertNodeNameAndType(t, "root/some_album/song.mp3", f, resultNode.mockNodes[0].mockNodes[0])
	assertNodeNameAndType(t, "root/some_other_album/song.mp3", f, resultNode.mockNodes[1].mockNodes[0])
}

func Test__tree_with_multiple_audio_files(t *testing.T) {
	setup()

	addFileToTree("root/some_album/song.mp3")
	addFileToTree("root/some_album/song2.mp3")
	addFileToTree("root/some_album/song3.mp3")

	fileTreeToItunes(fileTree)

	assertNodeNameAndType(t, "root", d, resultNode)
	assertNodeChildCountEquals(t, 1, resultNode)
	assertNodeNameAndType(t, "some_album", p, resultNode.mockNodes[0])
	assertNodeChildCountEquals(t, 3, resultNode.mockNodes[0])
	assertNodeNameAndType(t, "root/some_album/song.mp3", f, resultNode.mockNodes[0].mockNodes[0])
	assertNodeNameAndType(t, "root/some_album/song2.mp3", f, resultNode.mockNodes[0].mockNodes[1])
	assertNodeNameAndType(t, "root/some_album/song3.mp3", f, resultNode.mockNodes[0].mockNodes[2])
}

func Test__tree_with_single_deep_file_to_itunes(t *testing.T) {
	setup()

	addFileToTree("root/some_style/some_artist/some_album/song.mp3")

	fileTreeToItunes(fileTree)

	assertNodeNameAndType(t, "root", d, resultNode)
	assertNodeChildCountEquals(t, 1, resultNode)
	assertNodeNameAndType(t, "some_style", d, resultNode.mockNodes[0])
	assertNodeChildCountEquals(t, 1, resultNode.mockNodes[0])
	assertNodeNameAndType(t, "some_artist", d, resultNode.mockNodes[0].mockNodes[0])
	assertNodeChildCountEquals(t, 1, resultNode.mockNodes[0].mockNodes[0])
	assertNodeNameAndType(t, "some_album", p, resultNode.mockNodes[0].mockNodes[0].mockNodes[0])
	assertNodeChildCountEquals(t, 1, resultNode.mockNodes[0].mockNodes[0].mockNodes[0])
	assertNodeNameAndType(t, "root/some_style/some_artist/some_album/song.mp3", f, resultNode.mockNodes[0].mockNodes[0].mockNodes[0].mockNodes[0])
}

func Test__tree_with_multiple_files_at_different_depths(t *testing.T) {
	setup()

	addFileToTree("root/some_album/song.mp3")
	addFileToTree("root/some_other_album/song.mp3")
	addFileToTree("root/some_style/some_artist/some_album/song.mp3")

	fileTreeToItunes(fileTree)

	assertNodeNameAndType(t, "root", d, resultNode)
	assertNodeChildCountEquals(t, 3, resultNode)
	assertNodeNameAndType(t, "some_album", p, resultNode.mockNodes[0])
	assertNodeNameAndType(t, "some_other_album", p, resultNode.mockNodes[1])
	assertNodeNameAndType(t, "some_style", d, resultNode.mockNodes[2])
}

func Test__should_handle_songs_on_different_depths_correctly(t *testing.T) {
	setup()

	addFileToTree("root/some_artist/some_album/song.mp3")
	addFileToTree("root/some_artist/some_album/song1.mp3")
	addFileToTree("root/some_artist/some_album/cd1/song.mp3")

	fileTreeToItunes(fileTree)

	assertNodeNameAndType(t, "some_album", p, resultNode.mockNodes[0].mockNodes[0])
	assertNodeNameAndType(t, "root/some_artist/some_album/song.mp3", f, resultNode.mockNodes[0].mockNodes[0].mockNodes[0])
	assertNodeNameAndType(t, "root/some_artist/some_album/song1.mp3", f, resultNode.mockNodes[0].mockNodes[0].mockNodes[1])
	assertNodeNameAndType(t, "some_album", d, resultNode.mockNodes[0].mockNodes[1])
	assertNodeNameAndType(t, "cd1", p, resultNode.mockNodes[0].mockNodes[1].mockNodes[0])
	assertNodeNameAndType(t, "root/some_artist/some_album/cd1/song.mp3", f, resultNode.mockNodes[0].mockNodes[1].mockNodes[0].mockNodes[0])
}

func assertNodeNameAndType(t *testing.T, name string, typ itemType, n *mockNode) {
	assert.Equal(t, name, n.name)
	assert.EqualValues(t, typ, n.kind, "expected different itunes item type.")
}

func assertNodeChildCountEquals(t *testing.T, count int, n *mockNode) {
	assert.Equal(t, count, len(n.mockNodes))
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

func (m *mockNode) String() string {
	var children []string
	if len(m.mockNodes) > 0 {
		for _, child := range m.mockNodes {
			children = append(children, child.String())
		}
	}
	// return fmt.Sprintf("[n:%q id:%d ch:%d]", m.name, m.id, len(m.mockNodes))
	return fmt.Sprintf("[n:%q id:%d ch:%d]", m.name, m.id, children)
}

const (
	f = iota //file
	d        //dir
	p        //playlist
)

func addNode(name string, t itemType, parent int) int {
	count++
	newNode := &mockNode{name: name, id: count, kind: t}
	// fmt.Printf("==add %q==\n", name)
	// fmt.Printf("looking for Id:%d\n", parent)
	parentNode := findParent(resultNode, parent)
	if parentNode != nil && parentNode.name != "" {
		// fmt.Printf("\n    found:%q(%d)\n", parentNode.name, parentNode.id)
		parentNode.mockNodes = append(parentNode.mockNodes, newNode)
	} else {
		resultNode = newNode
	}
	// fmt.Println("complete node:", resultNode)
	return count
}

func findParent(currentNode *mockNode, parentID int) *mockNode {
	// fmt.Printf("   cn:%q", currentNode)
	if currentNode.id == parentID {
		return currentNode
	}
	for _, n := range currentNode.mockNodes {
		result := findParent(n, parentID)
		if result != nil {
			return result
		}
	}
	return nil
}

type mockInterface struct {
	currentID   int
	pathCreated string
}

func (i *mockInterface) newFolder(name string, id int) int {
	return addNode(name, d, id)
}
func (i *mockInterface) newPlaylist(name string, id int) int {
	return addNode(name, p, id)
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
