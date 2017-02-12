package main

import "testing"

func Test__should_be_able_to_make_retreive_and_delete_playlists(t *testing.T) {
	name := "go-osascript-itunes-test"

	playlistID := newPlaylist(name)

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
