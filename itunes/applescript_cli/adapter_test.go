package applescript_cli

import (
	"os"
	"testing"

	"github.com/meonlol/syncitunes/filescanner"
	"github.com/meonlol/syncitunes/tree"
	"github.com/stretchr/testify/assert"
)

var itunesIF = &Adapter{}

func Test__should_be_able_to_make_retreive_and_delete_playlists(t *testing.T) {
	name := "go-osascript-itunes-test"

	playlistID, _ := itunesIF.NewPlaylist(name, 0)
	result, _ := itunesIF.GetPlaylistIDByName(name)
	assert.Equal(t, playlistID, result)

	itunesIF.DeletePlaylistByID(playlistID)
	result, err := itunesIF.GetPlaylistIDByName(name)
	assert.Error(t, err)
}

func Test__should_be_able_to_make_retreive_and_delete_folders(t *testing.T) {
	parent := "go-osascript-itunes-test_folder"

	playlistID, _ := itunesIF.NewFolder(parent, 0)
	result, _ := itunesIF.GetPlaylistIDByName(parent)
	assert.Equal(t, playlistID, result)

	itunesIF.DeletePlaylistByID(playlistID)
	result, err := itunesIF.GetPlaylistIDByName(parent)
	assert.Error(t, err)
}

func Test__should_be_able_to_make_sub_folders(t *testing.T) {
	parentID, _ := itunesIF.NewFolder("go-osascript-itunes-test_sub_folder_test", 0)
	defer itunesIF.DeletePlaylistByID(parentID)

	subFolderID, _ := itunesIF.NewFolder("go-osascript-itunes-test_sub_folder", parentID)

	resultID, _ := itunesIF.GetParentIDForPlaylist(subFolderID)
	assert.Equal(t, parentID, resultID)
}

func Test__should_be_able_to_make_rereive_and_delete_playlists_inside_parent_folders(t *testing.T) {
	name := "go-osascript-itunes-test_playlist"
	folderID, _ := itunesIF.NewFolder("go-osascript-itunes-test_folder", 0)
	defer itunesIF.DeletePlaylistByID(folderID)

	playlistID, _ := itunesIF.NewPlaylist(name, folderID)
	parentID, _ := itunesIF.GetParentIDForPlaylist(playlistID)

	assert.Equal(t, parentID, folderID)
}

func Test__should_correctly_handle_retreiving_non_existent_parent(t *testing.T) {

	result, err := itunesIF.GetParentIDForPlaylist(1)

	assert.Error(t, err)
	assert.Empty(t, result)
}

func Test__should_add_file_to_itunes(t *testing.T) {
	playlistID, _ := itunesIF.NewPlaylist("test-playlist", 0)

	wd, _ := os.Getwd()
	fileID, _ := itunesIF.AddFileToPlaylist(wd+"/../t/empty.mp3", playlistID)

	assert.NotEqual(t, "", fileID)
	itunesIF.DeletePlaylistByID(playlistID)
}

// func Test__should_get_current_itunes_library(t *testing.T) {

// 	result, err := itunesIF.GetLibrary()

// 	fmt.Println("RESULT:" + result)
// 	assert.NoError(t, err)
// 	assert.NotEmpty(t, result)
// }
