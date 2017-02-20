package main

import "testing"

// func setupTest(t *testing.T) func() {
// 	// Test setup
// 	t.Log("setupTest()")

// 	// Test teardown - return a closure for use by 'defer'
// 	return func() {
// 		t.Log("teardownTest()")
// 	}
// }

func Test__sould_be_able_to_create_folder_structure(t *testing.T) {
	// defer setupTest(t)()
	mock := mockInterface{"b"}
	iTunes = &mock

	id := createPlaylist("album/1.mp3")

	if id != "1337" {
		t.Errorf("Expected mock id '1337', got %q", id)
	}
	if mock.file != "1.mp3" {
		t.Errorf("Expected '1.mp3', got %q", iTunes)
	}
}

type mockInterface struct {
	file string
}

func (mockInterface) newFolder(name string) string {
	return ""
}
func (i mockInterface) newPlaylist(name, parentID string) string {
	i.file = "bla"
	return "1337"
}
func (mockInterface) getPlaylistIDByName(name string) (string, error) {
	return "", nil
}
func (mockInterface) getParentIDForPlaylist(id string) (string, error) {
	return "", nil
}
func (mockInterface) deletePlaylistByID(id string) {
}
func (i *mockInterface) addFileToPlaylist(f string) {
	i.file = f
}
