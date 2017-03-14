package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var count = 0

func setup() {
	fileTree = &node{}
	mock := mockInterface{}
	iTunes = &mock
	resultNode = &mockNode{}
	count = 0
}

func Test__should_support_folder_playlist_and_song(t *testing.T) {
	setup()

	addFileToTree("root/some_album/song.mp3")

	fileTreeToItunes(fileTree)

	assertTreeMapHasNameAndType(t, []int{}, "root", d)
	assertTreeMapHasNameAndType(t, []int{0}, "some_album", p)
	assertTreeMapHasNameAndType(t, []int{0, 0}, "root/some_album/song.mp3", f)
}

func Test__should_support_files_in_differnt_folders(t *testing.T) {
	setup()

	addFileToTree("root/some_album/song.mp3")
	addFileToTree("root/some_other_album/song.mp3")

	fileTreeToItunes(fileTree)

	assertTreeMapHasNameAndType(t, []int{0}, "some_album", p)
	assertTreeMapHasNameAndType(t, []int{1}, "some_other_album", p)
	assertTreeMapHasNameAndType(t, []int{0, 0}, "root/some_album/song.mp3", f)
	assertTreeMapHasNameAndType(t, []int{1, 0}, "root/some_other_album/song.mp3", f)
}

func Test__should_support_multiple_audio_files_in_same_playlist(t *testing.T) {
	setup()

	addFileToTree("root/some_album/song.mp3")
	addFileToTree("root/some_album/song2.mp3")
	addFileToTree("root/some_album/song3.mp3")

	fileTreeToItunes(fileTree)

	assertTreeMapHasNameAndType(t, []int{0}, "some_album", p)
	assertTreeMapHasNameAndType(t, []int{0, 0}, "root/some_album/song.mp3", f)
	assertTreeMapHasNameAndType(t, []int{0, 1}, "root/some_album/song2.mp3", f)
	assertTreeMapHasNameAndType(t, []int{0, 2}, "root/some_album/song3.mp3", f)
}

func Test__should_support_recursive_nesting_of_nodes(t *testing.T) {
	setup()

	addFileToTree("root/some_style/some_artist/some_album/song.mp3")

	fileTreeToItunes(fileTree)

	assertTreeMapHasNameAndType(t, []int{}, "root", d)
	assertTreeMapHasNameAndType(t, []int{0}, "some_style", d)
	assertTreeMapHasNameAndType(t, []int{0, 0}, "some_artist", d)
	assertTreeMapHasNameAndType(t, []int{0, 0, 0}, "some_album", p)
	assertTreeMapHasNameAndType(t, []int{0, 0, 0, 0}, "root/some_style/some_artist/some_album/song.mp3", f)
}

func Test__should_support_mixed_folder_and_playlists(t *testing.T) {
	setup()

	addFileToTree("root/some_artist/some_album/song.mp3")
	addFileToTree("root/some_artist/some_album/song1.mp3")
	addFileToTree("root/some_artist/some_album/cd1/song.mp3")

	fileTreeToItunes(fileTree)

	assertTreeMapHasNameAndType(t, []int{0, 0}, "some_album", p)
	assertTreeMapHasNameAndType(t, []int{0, 0, 0}, "root/some_artist/some_album/song.mp3", f)
	assertTreeMapHasNameAndType(t, []int{0, 0, 1}, "root/some_artist/some_album/song1.mp3", f)
	assertTreeMapHasNameAndType(t, []int{0, 1}, "some_album", d)
	assertTreeMapHasNameAndType(t, []int{0, 1, 0}, "cd1", p)
	assertTreeMapHasNameAndType(t, []int{0, 1, 0, 0}, "root/some_artist/some_album/cd1/song.mp3", f)
}

func Test__should_support_single_root(t *testing.T) {
	setup()

	addFileToTree("root/A/A1/A-1.mp3")
	addFileToTree("root/A/A1/A-2.mp3")
	addFileToTree("root/A/A2/A-1.mp3")
	addFileToTree("root/B/B1/B-1.mp3")
	addFileToTree("root/C/C1/C-1.mp3")

	fileTreeToItunes(fileTree)

	assertTreeMapHasNameAndType(t, []int{}, "root", d)
	assertTreeMapHasNameAndType(t, []int{0}, "A", d)
	assertTreeMapHasNameAndType(t, []int{1}, "B", d)
	assertTreeMapHasNameAndType(t, []int{2}, "C", d)
	// assertTreeMapHasNameAndType(t, []int{0, 0}, "some_artist", d)
	// assertTreeMapHasNameAndType(t, []int{0, 0, 0}, "some_album", p)
	// assertTreeMapHasNameAndType(t, []int{0, 0, 0, 0}, "root/some_style/some_artist/some_album/song.mp3", f)
}

// checks supplied indexMapping exists and contains an item with specified name and type
func assertTreeMapHasNameAndType(t *testing.T, indexMapping []int, name string, typ itemType) {
	target := resultNode
	for _, i := range indexMapping {
		if len(target.mockNodes) > i {
			target = target.mockNodes[i]
		} else {
			t.Errorf("requested node[%d], but %q has only %d child nodes", i, target.name, len(target.mockNodes))
			t.Fail()
		}
	}
	assert.Equal(t, name, target.name)
	assert.EqualValues(t, typ, target.kind, "expected different itunes item type.")
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
	return fmt.Sprintf("[n:%q id:%d ch:%s]", m.name, m.id, children)
}

const (
	f = iota //file
	d        //dir
	p        //playlist
)

func addNode(name string, t itemType, parent int) int {
	count++
	newNode := &mockNode{name: name, id: count, kind: t}
	parentNode := findParent(resultNode, parent)
	if parentNode != nil && parentNode.name != "" {
		parentNode.mockNodes = append(parentNode.mockNodes, newNode)
	} else {
		resultNode = newNode
	}
	return count
}

func findParent(currentNode *mockNode, parentID int) *mockNode {
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

func (i *mockInterface) NewFolder(name string, id int) (int, error) {
	return addNode(name, d, id), nil
}
func (i *mockInterface) NewPlaylist(name string, id int) (int, error) {
	return addNode(name, p, id), nil
}
func (mockInterface) GetPlaylistIDByName(name string) (int, error) {
	return -1, nil
}
func (mockInterface) GetParentIDForPlaylist(id int) (int, error) {
	return -1, nil
}
func (i *mockInterface) AddFileToPlaylist(filePath string, playlistID int) (int, error) {
	return addNode(filePath, f, playlistID), nil
}
func (mockInterface) DeletePlaylistByID(id int) error {
	return nil
}
