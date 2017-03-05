package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test__should_walk(t *testing.T) {
	fileTree = &node{}
	wd, _ := os.Getwd()
	os.MkdirAll(wd+"/t_mp/root/playlist/", 0777)
	firstFile := wd + "/t_mp/root/playlist/1.mp3"
	secondFile := wd + "/t_mp/root/playlist/2.mp3"
	copy(wd+"/t/empty.mp3", firstFile)
	copy(wd+"/t/empty.mp3", secondFile)
	defer os.RemoveAll(wd + "/t_mp/")

	scanFolder(wd + "/t_mp/root/")

	assert.Equal(t, "playlist", fileTree.name)
	assert.Equal(t, 2, len(fileTree.nodes))
	assert.Equal(t, "1.mp3", fileTree.nodes[0].name)
	assert.Equal(t, firstFile, fileTree.nodes[0].path)
	assert.Equal(t, "2.mp3", fileTree.nodes[1].name)
	assert.Equal(t, secondFile, fileTree.nodes[1].path)
}

func Test__visit_should_ignore_dirs(t *testing.T) {
	fileTree = &node{}

	visit("/root/someDir", mockFileInfo{name: "someDir"}, nil)

	assert.Empty(t, fileTree.name)
}

func Test__visit_should_add_valid_types_to_list(t *testing.T) {
	fileTree = &node{}

	filePath := "root/someDir/someFile.mp3"
	filePath2 := "root/someOtherDir/someFile.aac"
	filePath3 := "root/someOtherDir/someFile.aac.bak"
	visit(filePath, mockFileInfo{"someFile.mp3", true}, nil)
	visit(filePath2, mockFileInfo{"someFile.aac", true}, nil)
	visit(filePath3, mockFileInfo{"someFile.aac.bak", true}, nil)

	assert.Equal(t, "root", fileTree.name)
	assert.Equal(t, "someFile.mp3", fileTree.nodes[0].nodes[0].name)
	assert.Equal(t, 1, len(fileTree.nodes[1].nodes))
	assert.Equal(t, "someFile.aac", fileTree.nodes[1].nodes[0].name)
}

func Test__should_add_single_node(t *testing.T) {
	fileTree = &node{}

	addFileToTree("1.mp3")

	assert.Equal(t, "1.mp3", fileTree.name)
}

func Test__should_add_sub_node(t *testing.T) {
	fileTree = &node{}

	addFileToTree("root/1.mp3")

	assert.Equal(t, "root", fileTree.name)
	assert.Equal(t, 1, len(fileTree.nodes))
	assert.Equal(t, "1.mp3", fileTree.nodes[0].name)
}

func Test__should_add_second_sub_node(t *testing.T) {
	fileTree = &node{}

	addFileToTree("root/1.mp3")
	addFileToTree("root/2.mp3")

	assert.Equal(t, "root", fileTree.name)
	assert.Equal(t, 2, len(fileTree.nodes))
	assert.Equal(t, "1.mp3", fileTree.nodes[0].name)
	assert.Equal(t, "2.mp3", fileTree.nodes[1].name)
}

// func Test__scanning_example(t *testing.T) {
// 	fileTree = &node{}

// 	scanFolder("/Users/leonmoll/leon/@music/")

// 	printTree(fileTree, 0)
// }

func Test__should_add_second_level_node(t *testing.T) {
	fileTree = &node{}

	addFileToTree("root/subFolder/1.mp3")

	assert.Equal(t, "root", fileTree.name)
	assert.Len(t, fileTree.nodes, 1)
	assert.Equal(t, "subFolder", fileTree.nodes[0].name)
	assert.Len(t, fileTree.nodes[0].nodes, 1)
}

func Test__should_add_multiple_second_level_nodes(t *testing.T) {
	fileTree = &node{}

	addFileToTree("root/subFolder/1.mp3")
	addFileToTree("root/subFolder/2.mp3")

	assert.Equal(t, "root", fileTree.name)
	assert.Equal(t, 1, len(fileTree.nodes))
	assert.Equal(t, "subFolder", fileTree.nodes[0].name)
	assert.Equal(t, 2, len(fileTree.nodes[0].nodes))
}

func Test__should_add_multiple_multilevel_nodes(t *testing.T) {
	fileTree = &node{}

	addFileToTree("root/subFolder/1.mp3")
	addFileToTree("root/subFolder/2.mp3")
	addFileToTree("root/subFolder2/3.mp3")

	assert.Equal(t, "root", fileTree.name)
	assert.Equal(t, 2, len(fileTree.nodes))
	assert.Equal(t, "subFolder", fileTree.nodes[0].name)
	assert.Equal(t, "subFolder2", fileTree.nodes[1].name)
	assert.Equal(t, "1.mp3", fileTree.nodes[0].nodes[0].name)
	assert.Equal(t, "2.mp3", fileTree.nodes[0].nodes[1].name)
	assert.Equal(t, "3.mp3", fileTree.nodes[1].nodes[0].name)
}

func Test__should_add_realistic_folder_structure(t *testing.T) {
	fileTree = &node{}

	addFileToTree("root/Aphrodite/Aphrodite - Urban Jungle/1.mp3")
	addFileToTree("root/Aphrodite/Aphrodite - Urban Jungle/2.mp3")
	addFileToTree("root/Arctic Monkeys/Arctic Monkeys - Artic Monkeys (Album)/1.mp3")
	addFileToTree("root/Arctic Monkeys/Arctic Monkeys - Artic Monkeys (Album)/2.mp3")
	addFileToTree("root/Arctic Monkeys/Arctic Monkeys - Artic Monkeys (Album)/3.mp3")
	addFileToTree("root/Arctic Monkeys/Arctic Monkeys - Favourite Worst Nightmare/1.mp3")
	addFileToTree("root/Arctic Monkeys/Arctic Monkeys - Favourite Worst Nightmare/2.mp3")

	assert.Equal(t, "root", fileTree.name)
	assert.Equal(t, 2, len(fileTree.nodes))
	assert.Equal(t, "Aphrodite", fileTree.nodes[0].name)
	assert.Equal(t, "Arctic Monkeys", fileTree.nodes[1].name)
	assert.Equal(t, 1, len(fileTree.nodes[0].nodes))
	assert.Equal(t, "Aphrodite - Urban Jungle", fileTree.nodes[0].nodes[0].name)
	assert.Equal(t, 2, len(fileTree.nodes[0].nodes[0].nodes))
	assert.Equal(t, 2, len(fileTree.nodes[1].nodes))
	assert.Equal(t, "Arctic Monkeys - Artic Monkeys (Album)", fileTree.nodes[1].nodes[0].name)
	assert.Equal(t, 3, len(fileTree.nodes[1].nodes[0].nodes))
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
