package main

type itunesInterface interface {
	newFolder(name string) string
	newPlaylist(name, parentID string) string
	getPlaylistIDByName(name string) (string, error)
	getParentIDForPlaylist(id string) (string, error)
	addFileToPlaylist(filePath, playlistID string) error
	deletePlaylistByID(id string) error
}

var iTunes itunesInterface
var root = ""

func createPlaylist(file string) string {
	parentID := iTunes.newFolder("someFolder")
	id := iTunes.newPlaylist("itunes-boundry", parentID)
	iTunes.addFileToPlaylist(file, id)
	return id
}
