package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test__should_nest_and_add_different_types(t *testing.T) {
	newMockInterface()

	rootID := addNode("root", d, 0)
	playlistID := addNode("some_album", p, rootID)
	addNode("root/some_album/song.mp3", f, playlistID)
	resultID := addNode("root/some_album/song2.mp3", f, playlistID)

	assert.Equal(t, 4, resultID)
	assertTreeMapHasNameAndType(t, []int{}, "root", d)
	assertTreeMapHasNameAndType(t, []int{0}, "some_album", p)
	assertTreeMapHasNameAndType(t, []int{0, 0}, "root/some_album/song.mp3", f)
	assertTreeMapHasNameAndType(t, []int{0, 1}, "root/some_album/song2.mp3", f)
}

func Test__should_handle_adding_on_different_nesting_depths(t *testing.T) {
	newMockInterface()

	rootID := addNode("root", d, 0)
	playlistID := addNode("some_album", p, rootID)
	addNode("root/some_album/song.mp3", f, playlistID)
	addNode("some_album2", p, rootID)

	assertTreeMapHasNameAndType(t, []int{}, "root", d)
	assertTreeMapHasNameAndType(t, []int{0}, "some_album", p)
	assertTreeMapHasNameAndType(t, []int{1}, "some_album2", p)
	assertTreeMapHasNameAndType(t, []int{0, 0}, "root/some_album/song.mp3", f)
}

// Helpers
//--------------------------------------------------------------------------------

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
