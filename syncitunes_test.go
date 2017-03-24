package main

import "testing"

func setup() {
	fileTree = &node{}
	mock := newMockInterface()
	iTunes = &mock
}

func Test__should_support_folder_playlist_and_song(t *testing.T) {
	setup()

	addFileToTree("root/some_album/song.mp3")

	fileTreeToItunes(fileTree, true)

	assertTreeMapHasNameAndType(t, []int{}, "root", d)
	assertTreeMapHasNameAndType(t, []int{0}, "some_album", p)
	assertTreeMapHasNameAndType(t, []int{0, 0}, "root/some_album/song.mp3", f)
}

func Test__should_support_files_in_differnt_folders(t *testing.T) {
	setup()

	addFileToTree("root/some_album/song.mp3")
	addFileToTree("root/some_other_album/song.mp3")

	fileTreeToItunes(fileTree, true)

	assertTreeMapHasNameAndType(t, []int{0}, "some_album", p)
	assertTreeMapHasNameAndType(t, []int{1}, "some_other_album", p)
	assertTreeMapHasNameAndType(t, []int{0, 0}, "root/some_album/song.mp3", f)
	assertTreeMapHasNameAndType(t, []int{1, 0}, "root/some_other_album/song.mp3", f)
}

func Test__should_support_multiple_audio_files_in_same_playlist(t *testing.T) {
	setup()

	addFileToTree("root/some_album/song.mp3")
	addFileToTree("root/some_album/song2.mp3")
	addFileToTree("root/some_album/song3.mp3")

	fileTreeToItunes(fileTree, true)

	assertTreeMapHasNameAndType(t, []int{0}, "some_album", p)
	assertTreeMapHasNameAndType(t, []int{0, 0}, "root/some_album/song.mp3", f)
	assertTreeMapHasNameAndType(t, []int{0, 1}, "root/some_album/song2.mp3", f)
	assertTreeMapHasNameAndType(t, []int{0, 2}, "root/some_album/song3.mp3", f)
}

func Test__should_support_recursive_nesting_of_nodes(t *testing.T) {
	setup()

	addFileToTree("root/some_style/some_artist/some_album/song.mp3")

	fileTreeToItunes(fileTree, true)

	assertTreeMapHasNameAndType(t, []int{}, "root", d)
	assertTreeMapHasNameAndType(t, []int{0}, "some_style", d)
	assertTreeMapHasNameAndType(t, []int{0, 0}, "some_artist", d)
	assertTreeMapHasNameAndType(t, []int{0, 0, 0}, "some_album", p)
	assertTreeMapHasNameAndType(t, []int{0, 0, 0, 0}, "root/some_style/some_artist/some_album/song.mp3", f)
}

func Test__should_support_mixed_folder_and_playlists(t *testing.T) {
	setup()

	addFileToTree("root/some_artist/some_album/song.mp3")
	addFileToTree("root/some_artist/some_album/song1.mp3")
	addFileToTree("root/some_artist/some_album/cd1/song.mp3")

	fileTreeToItunes(fileTree, true)

	assertTreeMapHasNameAndType(t, []int{0, 0}, "some_album", p)
	assertTreeMapHasNameAndType(t, []int{0, 0, 0}, "root/some_artist/some_album/song.mp3", f)
	assertTreeMapHasNameAndType(t, []int{0, 0, 1}, "root/some_artist/some_album/song1.mp3", f)
	assertTreeMapHasNameAndType(t, []int{0, 1}, "some_album", d)
	assertTreeMapHasNameAndType(t, []int{0, 1, 0}, "cd1", p)
	assertTreeMapHasNameAndType(t, []int{0, 1, 0, 0}, "root/some_artist/some_album/cd1/song.mp3", f)
}

func Test__should_support_single_root(t *testing.T) {
	setup()

	addFileToTree("root/A/A1/A-1.mp3")
	addFileToTree("root/A/A1/A-2.mp3")
	addFileToTree("root/A/A2/A-1.mp3")
	addFileToTree("root/B/B1/B-1.mp3")
	addFileToTree("root/C/C1/C-1.mp3")

	fileTreeToItunes(fileTree, true)

	assertTreeMapHasNameAndType(t, []int{}, "root", d)
	assertTreeMapHasNameAndType(t, []int{0}, "A", d)
	assertTreeMapHasNameAndType(t, []int{1}, "B", d)
	assertTreeMapHasNameAndType(t, []int{2}, "C", d)
}

func Test__should_support_not_including_root(t *testing.T) {
	setup()

	addFileToTree("root/A/A1/A-1.mp3")
	addFileToTree("root/B/B1/B-1.mp3")
	addFileToTree("root/C/C1/C-1.mp3")

	fileTreeToItunes(fileTree, false)

	assertTreeMapHasNameAndType(t, []int{}, "", 0)
	assertTreeMapHasNameAndType(t, []int{0}, "A", d)
	assertTreeMapHasNameAndType(t, []int{1}, "B", d)
	assertTreeMapHasNameAndType(t, []int{2}, "C", d)
}

// func Test__should_not_include_folder_if_already_added(t *testing.T) {
// 	setup()

// 	addFileToTree("root/art/alb/sng.mp3")

// 	fileTreeToItunes(fileTree, true)
// 	fileTreeToItunes(fileTree, true)

// 	// printMockTree(resultNode, 0)

// 	assertTreeMapHasNameAndType(t, []int{}, "root", d)
// 	assertTreeMapHasNameAndType(t, []int{0}, "art", d)
// }
