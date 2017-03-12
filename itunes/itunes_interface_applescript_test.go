package itunes

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var bitunes = &ApplescriptInterface{}

func Test__should_be_able_to_make_retreive_and_delete_playlists(t *testing.T) {
	name := "go-osascript-itunes-test"

	playlistID, _ := bitunes.NewPlaylist(name, 0)

	result, _ := bitunes.GetPlaylistIDByName(name)
	if result != playlistID {
		t.Errorf("Expected %q, got %q", playlistID, result)
	}

	bitunes.DeletePlaylistByID(playlistID)

	result, err := bitunes.GetPlaylistIDByName(name)
	if err == nil {
		t.Errorf("Expected an error, got nil with result %q", result)
	}
}

func Test__should_be_able_to_make_retreive_and_delete_folders(t *testing.T) {
	parent := "go-osascript-itunes-test_folder"

	playlistID, _ := bitunes.NewFolder(parent, 0)

	result, _ := bitunes.GetPlaylistIDByName(parent)
	if result != playlistID {
		t.Errorf("Expected %q, got %q", playlistID, result)
	}

	bitunes.DeletePlaylistByID(playlistID)

	result, err := bitunes.GetPlaylistIDByName(parent)
	if err == nil {
		t.Errorf("Expected an error, got nil with result %q", result)
	}
}

func Test__should_be_able_to_make_rereive_and_delete_playlists_inside_parent_folders(t *testing.T) {
	parent := "go-osascript-itunes-test_folder"
	name := "go-osascript-itunes-test_playlist"

	folderID, _ := bitunes.NewFolder(parent, 0)
	playlistID, _ := bitunes.NewPlaylist(name, folderID)
	parentID, _ := bitunes.GetParentIDForPlaylist(playlistID)

	if folderID != parentID {
		t.Errorf("Expected %q, got %q", folderID, parentID)
	}
	bitunes.DeletePlaylistByID(folderID)
}

func Test__should_correctly_handle_retreiving_non_existent_parent(t *testing.T) {

	result, err := bitunes.GetParentIDForPlaylist(1)

	assert.Error(t, err)
	assert.Empty(t, result)
}

func Test__should_add_file_to_itunes(t *testing.T) {
	playlistID, _ := bitunes.NewPlaylist("test-playlist", 0)

	wd, _ := os.Getwd()
	fileID, _ := bitunes.AddFileToPlaylist(wd+"/../t/empty.mp3", playlistID)

	assert.NotEqual(t, "", fileID)
	bitunes.DeletePlaylistByID(playlistID)
}
