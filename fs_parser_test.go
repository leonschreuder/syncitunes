package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test__visit_should_ignore_nodes(t *testing.T) {
	filesList = []string{}

	visit("/root/someDir", mockFileInfo{name: "someDir"}, nil)

	assert.Empty(t, filesList)
}

func Test__visit_should_add_valid_types_to_list(t *testing.T) {
	filesList = []string{}

	filePath := "/root/someDir/someFile.mp3"
	filePath2 := "/root/someOtherDir/someFile.aac"
	filePath3 := "/root/someOtherDir/someFile.aac.bak"
	visit(filePath, mockFileInfo{"someFile.mp3", true}, nil)
	visit(filePath2, mockFileInfo{"someFile.aac", true}, nil)
	visit(filePath3, mockFileInfo{"someFile.aac.bak", true}, nil)

	assert.Equal(t, 2, len(filesList))
	assert.Equal(t, filePath, filesList[0])
	assert.Equal(t, filePath2, filesList[1])
}

func Test__visit_should_not_pick_other_files(t *testing.T) {
	filesList = []string{}

	visit("/root/someDir/someFile.txt", mockFileInfo{"someFile.txt", true}, nil)

	assert.Empty(t, filesList)
}

// func Test__should_build_tree_from_single_file_in_several_folders(t *testing.T) {
// 	wd, _ := os.Getwd()
// 	os.MkdirAll(wd+"/t_mp/root/playlist/", 0777)
// 	copy(wd+"/t/empty.mp3", wd+"/t_mp/root/playlist/1.mp3")
// 	defer os.RemoveAll(wd + "/t_mp/")

// 	findMusic(wd + "/t_mp/root/")

// 	fmt.Println("tree:", treeFound)
// 	assert.Equal(t, "root", treeFound.name)
// 	assert.Equal(t, 1, len(treeFound.nodes))
// 	assert.Equal(t, "playlist", treeFound.nodes[0].name)
// 	assert.Equal(t, 1, len(treeFound.nodes[0].nodes))
// 	assert.Equal(t, "1.mp3", treeFound.nodes[0].nodes[0].name)
// }

// func Test__visit_should_add_multiple_nodes_on_same_level(t *testing.T) {
// 	treeFound = neoNode{name: "root"}
// 	currentNode = &treeFound

// 	visit("root/subDir", mockFileInfo{"subDir"}, nil)
// 	visit("root/otherDir", mockFileInfo{"otherDir"}, nil)

// 	assert.Equal(t, 2, len(treeFound.nodes))
// 	assert.Equal(t, treeFound.nodes[1], currentNode)
// }

// func Test__visit_should_add_nodes_recursively(t *testing.T) {
// 	treeFound = neoNode{name: "root"}
// 	currentNode = &treeFound

// 	visit("root/subDir", mockFileInfo{"subDir"}, nil)
// 	visit("root/subDir/subSubDir", mockFileInfo{"subSubDir"}, nil)

// 	assert.Equal(t, 1, len(treeFound.nodes))
// }

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

func Test__should_walk(t *testing.T) {
	wd, _ := os.Getwd()
	os.MkdirAll(wd+"/t_mp/root/playlist/", 0777)
	firstFile := wd + "/t_mp/root/playlist/1.mp3"
	secondFile := wd + "/t_mp/root/playlist/2.mp3"
	copy(wd+"/t/empty.mp3", firstFile)
	copy(wd+"/t/empty.mp3", secondFile)
	defer os.RemoveAll(wd + "/t_mp/")

	findMusic(wd + "/t_mp/root/")

	fmt.Println("tree:", filesList)
	assert.Equal(t, firstFile, filesList[0])
	assert.Equal(t, secondFile, filesList[1])
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
