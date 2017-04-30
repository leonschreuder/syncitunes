package mock

import (
	"testing"

	"github.com/meonlol/syncitunes/itunes"
	"github.com/stretchr/testify/assert"
)

func Test__should_nest_and_add_different_types(t *testing.T) {
	NewMockInterface()

	rootID := addNode("root", itunes.Dir, 0)
	playlistID := addNode("some_album", itunes.Playlist, rootID)
	addNode("root/some_album/song.mp3", itunes.File, playlistID)
	resultID := addNode("root/some_album/song2.mp3", itunes.File, playlistID)

	assert.Equal(t, 4, resultID)
	AssertTreeMapHasNameAndType(t, MockTree, []int{}, "root", itunes.Dir)
	AssertTreeMapHasNameAndType(t, MockTree, []int{0}, "some_album", itunes.Playlist)
	AssertTreeMapHasNameAndType(t, MockTree, []int{0, 0}, "root/some_album/song.mp3", itunes.File)
	AssertTreeMapHasNameAndType(t, MockTree, []int{0, 1}, "root/some_album/song2.mp3", itunes.File)
}

func Test__should_handle_adding_on_different_nesting_depths(t *testing.T) {
	NewMockInterface()

	rootID := addNode("root", itunes.Dir, 0)
	playlistID := addNode("some_album", itunes.Playlist, rootID)
	addNode("root/some_album/song.mp3", itunes.File, playlistID)
	addNode("some_album2", itunes.Playlist, rootID)

	AssertTreeMapHasNameAndType(t, MockTree, []int{}, "root", itunes.Dir)
	AssertTreeMapHasNameAndType(t, MockTree, []int{0}, "some_album", itunes.Playlist)
	AssertTreeMapHasNameAndType(t, MockTree, []int{1}, "some_album2", itunes.Playlist)
	AssertTreeMapHasNameAndType(t, MockTree, []int{0, 0}, "root/some_album/song.mp3", itunes.File)
}

// Helpers
//--------------------------------------------------------------------------------
