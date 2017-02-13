package main

import "testing"

func Test__should_be_able_to_make_retreive_and_delete_playlists(t *testing.T) {
	name := "go-osascript-itunes-test"

	playlistID := newPlaylist(name, "")

	result, _ := getPlaylistIDByName(name)
	if result != playlistID {
		t.Errorf("Expected %q, got %q", playlistID, result)
	}

	deletePlaylistByID(playlistID)

	result, err := getPlaylistIDByName(name)
	if err == nil {
		t.Errorf("Expected an error, got nil with result %q", result)
	}
}

func Test__should_be_able_to_make_retreive_and_delete_folders(t *testing.T) {
	parent := "go-osascript-itunes-test_folder"

	playlistID := newFolder(parent)

	result, _ := getPlaylistIDByName(parent)
	if result != playlistID {
		t.Errorf("Expected %q, got %q", playlistID, result)
	}

	deletePlaylistByID(playlistID)

	result, err := getPlaylistIDByName(parent)
	if err == nil {
		t.Errorf("Expected an error, got nil with result %q", result)
	}
}

func Test__should_be_able_to_make_rereive_and_delete_playlists_inside_parent_folders(t *testing.T) {
	parent := "go-osascript-itunes-test_folder"
	name := "go-osascript-itunes-test_playlist"

	folderID := newFolder(parent)
	playlistID := newPlaylist(name, folderID)
	parentID, _ := getParentIDForPlaylist(playlistID)

	if folderID != parentID {
		t.Errorf("Expected %q, got %q", folderID, parentID)
	}
	deletePlaylistByID(folderID)
}

func Test__should_correctly_handle_retreiving_non_existent_parent(t *testing.T) {

	result, err := getParentIDForPlaylist("1")

	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
	if result != "" {
		t.Errorf("Expected an empty stdout, got %q", result)
	}
}
