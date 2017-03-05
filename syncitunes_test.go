package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test__sould_be_able_to_create_playlist_with_file(t *testing.T) {
	mock := mockInterface{}
	iTunes = &mock

	wd, _ := os.Getwd()
	root = wd
	absFile := wd + "/itunes-boundry/empty.mp3"
	id := createPlaylist(absFile)

	assert.Equal(t, "1337", id)
	assert.Equal(t, "itunes-boundry", mock.playlistName)
	assert.Equal(t, absFile, mock.file)
}

func Test__sould_be_able_to_create_playlist_in_folder(t *testing.T) {
	mock := mockInterface{}
	iTunes = &mock

	wd, _ := os.Getwd()
	root = wd
	absFile := wd + "/someFolder/itunes-boundry/empty.mp3"
	createPlaylist(absFile)

	assert.Equal(t, "itunes-boundry", mock.playlistName)
	assert.Equal(t, "someFolder", mock.folderCreated)
	assert.Equal(t, absFile, mock.file)
}

type mockInterface struct {
	file          string
	playlistName  string
	folderCreated string
}

func (i *mockInterface) newFolder(name string) string {
	i.folderCreated = name
	return ""
}
func (i *mockInterface) newPlaylist(name, parentID string) string {
	i.playlistName = name
	return "1337"
}
func (mockInterface) getPlaylistIDByName(name string) (string, error) {
	return "", nil
}
func (mockInterface) getParentIDForPlaylist(id string) (string, error) {
	return "", nil
}
func (i *mockInterface) addFileToPlaylist(filePath, playlistID string) error {
	i.file = filePath
	return nil
}

func (mockInterface) deletePlaylistByID(id string) error {
	return nil
}
