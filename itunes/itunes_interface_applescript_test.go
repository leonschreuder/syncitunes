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
	assert.Equal(t, playlistID, result)

	bitunes.DeletePlaylistByID(playlistID)
	result, err := bitunes.GetPlaylistIDByName(name)
	assert.Error(t, err)
}

func Test__should_be_able_to_make_retreive_and_delete_folders(t *testing.T) {
	parent := "go-osascript-itunes-test_folder"

	playlistID, _ := bitunes.NewFolder(parent, 0)
	result, _ := bitunes.GetPlaylistIDByName(parent)
	assert.Equal(t, playlistID, result)

	bitunes.DeletePlaylistByID(playlistID)
	result, err := bitunes.GetPlaylistIDByName(parent)
	assert.Error(t, err)
}

func Test__should_be_able_to_make_sub_folders(t *testing.T) {
	parentID, _ := bitunes.NewFolder("go-osascript-itunes-test_sub_folder_test", 0)
	defer bitunes.DeletePlaylistByID(parentID)

	subFolderID, _ := bitunes.NewFolder("go-osascript-itunes-test_sub_folder", parentID)

	resultID, _ := bitunes.GetParentIDForPlaylist(subFolderID)
	assert.Equal(t, parentID, resultID)
}

func Test__should_be_able_to_make_rereive_and_delete_playlists_inside_parent_folders(t *testing.T) {
	name := "go-osascript-itunes-test_playlist"
	folderID, _ := bitunes.NewFolder("go-osascript-itunes-test_folder", 0)
	defer bitunes.DeletePlaylistByID(folderID)

	playlistID, _ := bitunes.NewPlaylist(name, folderID)
	parentID, _ := bitunes.GetParentIDForPlaylist(playlistID)

	assert.Equal(t, parentID, folderID)
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
