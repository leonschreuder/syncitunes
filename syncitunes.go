package main

type itunesInterface interface {
	newFolder(name string) string
	newPlaylist(name, parentID string) string
	getPlaylistIDByName(name string) (string, error)
	getParentIDForPlaylist(id string) (string, error)
	deletePlaylistByID(id string)
	addFileToPlaylist(id string)
}

var iTunes itunesInterface

func createPlaylist(file string) string {
	id := iTunes.newPlaylist("album", "")
	iTunes.addFileToPlaylist("1.mp3")
	return id
}
