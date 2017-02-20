package itunes

import (
	"os"
	"testing"
)

var bitunes = &applescriptInterface{}

func Test__should_be_able_to_make_retreive_and_delete_playlists(t *testing.T) {
	name := "go-osascript-itunes-test"

	playlistID := bitunes.newPlaylist(name, "")

	result, _ := bitunes.getPlaylistIDByName(name)
	if result != playlistID {
		t.Errorf("Expected %q, got %q", playlistID, result)
	}

	bitunes.deletePlaylistByID(playlistID)

	result, err := bitunes.getPlaylistIDByName(name)
	if err == nil {
		t.Errorf("Expected an error, got nil with result %q", result)
	}
}

func Test__should_be_able_to_make_retreive_and_delete_folders(t *testing.T) {
	parent := "go-osascript-itunes-test_folder"

	playlistID := bitunes.newFolder(parent)

	result, _ := bitunes.getPlaylistIDByName(parent)
	if result != playlistID {
		t.Errorf("Expected %q, got %q", playlistID, result)
	}

	bitunes.deletePlaylistByID(playlistID)

	result, err := bitunes.getPlaylistIDByName(parent)
	if err == nil {
		t.Errorf("Expected an error, got nil with result %q", result)
	}
}

func Test__should_be_able_to_make_rereive_and_delete_playlists_inside_parent_folders(t *testing.T) {
	parent := "go-osascript-itunes-test_folder"
	name := "go-osascript-itunes-test_playlist"

	folderID := bitunes.newFolder(parent)
	playlistID := bitunes.newPlaylist(name, folderID)
	parentID, _ := bitunes.getParentIDForPlaylist(playlistID)

	if folderID != parentID {
		t.Errorf("Expected %q, got %q", folderID, parentID)
	}
	bitunes.deletePlaylistByID(folderID)
}

func Test__should_correctly_handle_retreiving_non_existent_parent(t *testing.T) {

	result, err := bitunes.getParentIDForPlaylist("1")

	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	if result != "" {
		t.Errorf("Expected an empty stdout, got %q", result)
	}
}

func Test__should_add_file_to_itunes(t *testing.T) {
	id := bitunes.newPlaylist("test-playlist", "")

	wd, _ := os.Getwd()
	bitunes.addFileToPlaylist(wd+"/empty.mp3", id)

	bitunes.deletePlaylistByID(id)
}
