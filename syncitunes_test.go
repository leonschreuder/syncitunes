package main

import (
	"testing"

	"github.com/meonlol/syncitunes/filescanner"
	"github.com/meonlol/syncitunes/itunes"
	"github.com/meonlol/syncitunes/itunes/mock"
	"github.com/meonlol/syncitunes/tree"
)

func setup() {
	filescanner.FileTree = &tree.Node{}
	mock := mock.NewMockInterface()
	iTunes = &mock
}

func Test__should_support_folder_playlist_and_song(t *testing.T) {
	setup()

	filescanner.AddFileToTree("root/some_album/song.mp3")
	fileTree = filescanner.FileTree

	fileTreeToItunes(fileTree, true)

	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{}, "root", itunes.Dir)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0}, "some_album", itunes.Playlist)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0, 0}, "root/some_album/song.mp3", itunes.File)
}

func Test__should_support_files_in_differnt_folders(t *testing.T) {
	setup()

	filescanner.AddFileToTree("root/some_album/song.mp3")
	filescanner.AddFileToTree("root/some_other_album/song.mp3")
	fileTree = filescanner.FileTree

	fileTreeToItunes(fileTree, true)

	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0}, "some_album", itunes.Playlist)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{1}, "some_other_album", itunes.Playlist)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0, 0}, "root/some_album/song.mp3", itunes.File)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{1, 0}, "root/some_other_album/song.mp3", itunes.File)
}

func Test__should_support_multiple_audio_files_in_same_playlist(t *testing.T) {
	setup()

	filescanner.AddFileToTree("root/some_album/song.mp3")
	filescanner.AddFileToTree("root/some_album/song2.mp3")
	filescanner.AddFileToTree("root/some_album/song3.mp3")
	fileTree = filescanner.FileTree

	fileTreeToItunes(fileTree, true)

	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0}, "some_album", itunes.Playlist)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0, 0}, "root/some_album/song.mp3", itunes.File)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0, 1}, "root/some_album/song2.mp3", itunes.File)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0, 2}, "root/some_album/song3.mp3", itunes.File)
}

func Test__should_support_recursive_nesting_of_nodes(t *testing.T) {
	setup()

	filescanner.AddFileToTree("root/some_style/some_artist/some_album/song.mp3")
	fileTree = filescanner.FileTree

	fileTreeToItunes(fileTree, true)

	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{}, "root", itunes.Dir)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0}, "some_style", itunes.Dir)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0, 0}, "some_artist", itunes.Dir)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0, 0, 0}, "some_album", itunes.Playlist)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0, 0, 0, 0}, "root/some_style/some_artist/some_album/song.mp3", itunes.File)
}

func Test__should_support_mixed_folder_and_playlists(t *testing.T) {
	setup()

	filescanner.AddFileToTree("root/some_artist/some_album/song.mp3")
	filescanner.AddFileToTree("root/some_artist/some_album/song1.mp3")
	filescanner.AddFileToTree("root/some_artist/some_album/cd1/song.mp3")
	fileTree = filescanner.FileTree

	fileTreeToItunes(fileTree, true)

	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0, 0}, "some_album", itunes.Playlist)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0, 0, 0}, "root/some_artist/some_album/song.mp3", itunes.File)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0, 0, 1}, "root/some_artist/some_album/song1.mp3", itunes.File)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0, 1}, "some_album", itunes.Dir)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0, 1, 0}, "cd1", itunes.Playlist)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0, 1, 0, 0}, "root/some_artist/some_album/cd1/song.mp3", itunes.File)
}

func Test__should_support_single_root(t *testing.T) {
	setup()

	filescanner.AddFileToTree("root/A/A1/A-1.mp3")
	filescanner.AddFileToTree("root/A/A1/A-2.mp3")
	filescanner.AddFileToTree("root/A/A2/A-1.mp3")
	filescanner.AddFileToTree("root/B/B1/B-1.mp3")
	filescanner.AddFileToTree("root/C/C1/C-1.mp3")
	fileTree = filescanner.FileTree

	fileTreeToItunes(fileTree, true)

	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{}, "root", itunes.Dir)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0}, "A", itunes.Dir)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{1}, "B", itunes.Dir)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{2}, "C", itunes.Dir)
}

func Test__should_support_not_including_root(t *testing.T) {
	setup()

	filescanner.AddFileToTree("root/A/A1/A-1.mp3")
	filescanner.AddFileToTree("root/B/B1/B-1.mp3")
	filescanner.AddFileToTree("root/C/C1/C-1.mp3")
	fileTree = filescanner.FileTree

	fileTreeToItunes(fileTree, false)

	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{}, "", 0)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0}, "A", itunes.Dir)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{1}, "B", itunes.Dir)
	mock.AssertTreeMapHasNameAndType(t, mock.MockTree, []int{2}, "C", itunes.Dir)
}

// func Test__should_not_include_folder_if_already_added(t *testing.T) {
// 	setup()

// 	filescanner.AddFileToTree("root/art/alb/sng.mp3")

// 	fileTreeToItunes(fileTree, true)
// 	fileTreeToItunes(fileTree, true)

// 	// printMockTree(resultNode, 0)

// 	AssertTreeMapHasNameAndType(t, mock.MockTree, []int{}, "root", itunes.Dir)
// 	AssertTreeMapHasNameAndType(t, mock.MockTree, []int{0}, "art", itunes.Dir)
// }
