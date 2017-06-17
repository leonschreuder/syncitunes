package filescanner

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/meonlol/syncitunes/tree"
	"github.com/stretchr/testify/assert"
)

func Test__should_walk(t *testing.T) {
	FileTree = &tree.Node{}
	// fileTree = &tree.Node{}
	wd, _ := os.Getwd()
	tRoot := wd + "/t_mp/"
	files := []string{
		"artist/album/1.mp3",
		"artist/album/2.mp3",
		"other_artist/other_album/1.mp3",
		"artist3/album3/1.mp3",
	}
	createFiles(tRoot, files)
	defer os.RemoveAll(tRoot)

	ScanFolder(tRoot)

	assertTreeMapHasNode(t, []int{}, "t_mp")
	assertTreeMapHasNode(t, []int{0}, "artist")
	assertTreeMapHasNode(t, []int{0, 0}, "album")
	assertTreeMapHasNode(t, []int{0, 0, 0}, "1.mp3")
	assertTreeMapHasNode(t, []int{0, 0, 1}, "2.mp3")
	assertTreeMapHasNode(t, []int{2}, "other_artist")
	assertTreeMapHasNode(t, []int{2, 0}, "other_album")
	assertTreeMapHasNode(t, []int{1}, "artist3")
	assertTreeMapHasNode(t, []int{1, 0}, "album3")
	assert.Equal(t, tRoot+files[0], FileTree.Nodes[0].Nodes[0].Nodes[0].Path)
	assert.Equal(t, tRoot+files[1], FileTree.Nodes[0].Nodes[0].Nodes[1].Path)
}

func createFiles(root string, files []string) {
	wd, _ := os.Getwd()
	for _, f := range files {
		os.MkdirAll(filepath.Dir(root+f), 0777)
		copy(wd+"/../t/empty.mp3", root+f)
	}
}

func Test__visit_should_ignore_dirs(t *testing.T) {
	FileTree = &tree.Node{}

	visit("/root/someDir", mockFileInfo{name: "someDir"}, nil)

	assert.Empty(t, FileTree.Name)
}

func Test__visit_should_add_valid_types_to_list(t *testing.T) {
	FileTree = &tree.Node{}

	visit("root/someDir/someFile.mp3", mockFileInfo{"someFile.mp3", true}, nil)
	visit("root/someOtherDir/someFile.aac", mockFileInfo{"someFile.aac", true}, nil)
	visit("root/someOtherDir/someFile.aac.bak", mockFileInfo{"someFile.aac.bak", true}, nil)

	assertTreeMapHasNode(t, []int{0, 0}, "someFile.mp3")
	assertTreeMapHasNode(t, []int{1, 0}, "someFile.aac")
	assert.Equal(t, 1, len(FileTree.Nodes[1].Nodes))
}

func Test__should_add_sub_node(t *testing.T) {
	FileTree = &tree.Node{}

	AddFileToTree("root/1.mp3")

	assertTreeMapHasNode(t, []int{}, "root")
	assertTreeMapHasNode(t, []int{0}, "1.mp3")
}

func Test__should_add_second_sub_node(t *testing.T) {
	FileTree = &tree.Node{}

	AddFileToTree("root/1.mp3")
	AddFileToTree("root/2.mp3")

	assertTreeMapHasNode(t, []int{0}, "1.mp3")
	assertTreeMapHasNode(t, []int{1}, "2.mp3")
}

func Test__should_add_second_level_node(t *testing.T) {
	FileTree = &tree.Node{}

	AddFileToTree("root/subFolder/1.mp3")

	assertTreeMapHasNode(t, []int{0}, "subFolder")
	assertTreeMapHasNode(t, []int{0, 0}, "1.mp3")
}

func Test__should_add_multiple_second_level_nodes(t *testing.T) {
	FileTree = &tree.Node{}

	AddFileToTree("root/subFolder/1.mp3")
	AddFileToTree("root/subFolder/2.mp3")

	assertTreeMapHasNode(t, []int{0}, "subFolder")
	assertTreeMapHasNode(t, []int{0, 0}, "1.mp3")
	assertTreeMapHasNode(t, []int{0, 1}, "2.mp3")
}

func Test__should_add_multiple_multilevel_nodes(t *testing.T) {
	FileTree = &tree.Node{}

	AddFileToTree("root/subFolder/1.mp3")
	AddFileToTree("root/subFolder/2.mp3")
	AddFileToTree("root/subFolder2/3.mp3")

	assertTreeMapHasNode(t, []int{0}, "subFolder")
	assertTreeMapHasNode(t, []int{0, 0}, "1.mp3")
	assertTreeMapHasNode(t, []int{0, 1}, "2.mp3")
	assertTreeMapHasNode(t, []int{1}, "subFolder2")
	assertTreeMapHasNode(t, []int{1, 0}, "3.mp3")
}

func Test__should_add_realistic_folder_structure(t *testing.T) {
	FileTree = &tree.Node{}

	AddFileToTree("root/Aphrodite/Aphrodite - Urban Jungle/1.mp3")
	AddFileToTree("root/Aphrodite/Aphrodite - Urban Jungle/2.mp3")
	AddFileToTree("root/Arctic Monkeys/Arctic Monkeys - Artic Monkeys (Album)/1.mp3")
	AddFileToTree("root/Arctic Monkeys/Arctic Monkeys - Artic Monkeys (Album)/2.mp3")
	AddFileToTree("root/Arctic Monkeys/Arctic Monkeys - Artic Monkeys (Album)/3.mp3")
	AddFileToTree("root/Arctic Monkeys/Arctic Monkeys - Favourite Worst Nightmare/1.mp3")
	AddFileToTree("root/Arctic Monkeys/Arctic Monkeys - Favourite Worst Nightmare/2.mp3")

	assertTreeMapHasNode(t, []int{}, "root")
	assertTreeMapHasNode(t, []int{0}, "Aphrodite")
	assertTreeMapHasNode(t, []int{1}, "Arctic Monkeys")
	assertTreeMapHasNode(t, []int{0, 0}, "Aphrodite - Urban Jungle")
	assertTreeMapHasNode(t, []int{0, 0, 0}, "1.mp3")
	assertTreeMapHasNode(t, []int{0, 0, 1}, "2.mp3")
	assertTreeMapHasNode(t, []int{1, 0}, "Arctic Monkeys - Artic Monkeys (Album)")
	assertTreeMapHasNode(t, []int{1, 0, 2}, "3.mp3")
	assertTreeMapHasNode(t, []int{1, 1}, "Arctic Monkeys - Favourite Worst Nightmare")
	assertTreeMapHasNode(t, []int{1, 1, 1}, "2.mp3")
}

// checks supplied indexMapping exists and contains an item with specified name and type
func assertTreeMapHasNode(t *testing.T, indexMapping []int, nodeName string) {
	target := FileTree
	for _, i := range indexMapping {
		if len(target.Nodes) > i {
			target = target.Nodes[i]
		} else {
			t.Errorf("requested Node[%d], but %q has only %d child nodes", i, target.Name, len(target.Nodes))
			t.Fail()
		}
	}
	assert.Equal(t, nodeName, target.Name)
}

type mockFileInfo struct {
	name   string
	isFile bool
}

func (m mockFileInfo) Name() string {
	return m.name
}
func (m mockFileInfo) Size() int64 {
	return 0
}
func (m mockFileInfo) Mode() os.FileMode {
	return 0
}
func (m mockFileInfo) ModTime() time.Time {
	return time.Now()
}
func (m mockFileInfo) IsDir() bool {
	return !m.isFile
}
func (m mockFileInfo) Sys() interface{} {
	return nil
}

func copy(src string, dst string) {
	data, err := ioutil.ReadFile(src)
	checkErr(err)
	err = ioutil.WriteFile(dst, data, 0644)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
