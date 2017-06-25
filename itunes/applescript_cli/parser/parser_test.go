package parser

import (
	"testing"

	"github.com/meonlol/syncitunes/filescanner"
	"github.com/meonlol/syncitunes/tree"
	"github.com/stretchr/testify/assert"
)

// https://youtu.be/HxaD_trXwRE
// https://golang.org/doc/effective_go.html#channels
// https://golang.org/src/text/template/parse/parse.go
// https://github.com/golang/go/blob/master/src/text/template/parse/lex.go

// {
// 	{"Library", 19192, -1, {}},
// 	 {"albums", 35077, -1, {}},
// 	 {"Air", 38327, 35077, {}},
// 	 {"Air - 10000 Hz Legend", 38361, 38327, {
// 			alias "Macintosh HD:Users:leonmoll:leon:@music:albums:Air:Air - 10000 Hz Legend:01-Electronic Performers.mp3",
// 			alias "Macintosh HD:Users:leonmoll:leon:@music:albums:Air:Air - 10000 Hz Legend:02-How Does It Make You Feel_.mp3",
// 			alias "Macintosh HD:Users:leonmoll:leon:@music:albums:Air:Air - 10000 Hz Legend:03-Radio #1.mp3",
// 			alias "Macintosh HD:Users:leonmoll:leon:@music:albums:Air:Air - 10000 Hz Legend:04-The Vagabond (feat. Beck).mp3",
// 			alias "Macintosh HD:Users:leonmoll:leon:@music:albums:Air:Air - 10000 Hz Legend:05-Radian.mp3",
// 			alias "Macintosh HD:Users:leonmoll:leon:@music:albums:Air:Air - 10000 Hz Legend:06-Lucky and Unhappy.mp3",
// 			alias "Macintosh HD:Users:leonmoll:leon:@music:albums:Air:Air - 10000 Hz Legend:07-Sex Born Poison (feat. Buffalo Daughter).mp3",
// 			alias "Macintosh HD:Users:leonmoll:leon:@music:albums:Air:Air - 10000 Hz Legend:08-People in the City.mp3",
// 			alias "Macintosh HD:Users:leonmoll:leon:@music:albums:Air:Air - 10000 Hz Legend:09-Wonder Milky Bitch.mp3",
// 			alias "Macintosh HD:Users:leonmoll:leon:@music:albums:Air:Air - 10000 Hz Legend:10-Don't Be Light.mp3",
// 			alias "Macintosh HD:Users:leonmoll:leon:@music:albums:Air:Air - 10000 Hz Legend:11-Caramel Prisoner.mp3"}},

func Test__should_parse_root_object(t *testing.T) {
	input := `{{"bla", 14, -1, {}}}`

	result := parse(input)

	assert.NotNil(t, result)
	assert.Equal(t, "bla", result[0].name)
	assert.Equal(t, 14, result[0].ID)
}

func Test__should_parse_non_root_object(t *testing.T) {
	input := `{{"bla", 15, 13, {}}}`

	result := parse(input)

	assert.NotNil(t, result)
	assert.Equal(t, "bla", result[0].name)
	assert.Equal(t, 15, result[0].ID)
	assert.Equal(t, 13, result[0].parentID)
}

func Test__should_parse_multiple_objects(t *testing.T) {
	input := `{{"root", 14}{"lvl1A", 15, 14}{"lvl1B", 16, 14}}`

	result := parse(input)

	assert.NotNil(t, result)
	assert.Equal(t, "root", result[0].name)
	assert.Equal(t, 14, result[0].ID)
	assert.Equal(t, "lvl1A", result[1].name)
	assert.Equal(t, 15, result[1].ID)
	assert.Equal(t, 14, result[1].parentID)
	assert.Equal(t, "lvl1B", result[2].name)
	assert.Equal(t, 16, result[2].ID)
	assert.Equal(t, 14, result[2].parentID)
}

func Test__should_parse_tree_with_a_song(t *testing.T) {
	input := `{{"albums", 10, -1, {}},{"Air", 11, 10, {}},{"Air - 10000 Hz Legend", 12, 11, {alias "Macintosh HD:Users:music:albums:Air:Air - 10000 Hz Legend:01-Electronic Performers.mp3"}}}`

	result := parse(input)

	assert.NotNil(t, result)
	assert.Equal(t, "albums", result[0].name)
	assert.Equal(t, 10, result[0].ID)
	assert.Equal(t, "Air", result[1].name)
	assert.Equal(t, 11, result[1].ID)
	assert.Equal(t, 10, result[1].parentID)

	assert.Equal(t, "Air - 10000 Hz Legend", result[2].name)
	assert.Equal(t, 12, result[2].ID)
	assert.Equal(t, 11, result[2].parentID)
	assert.Equal(t, 1, len(result[2].songs))
	assert.Equal(t, 1, len(result[2].songs))
	assert.Equal(t, 7, len(result[2].songs[0].pathElements))
}

func Test__should_parse_tree_with_multiple_songs(t *testing.T) {
	input := `{{"mixes", 10, -1, {alias "Disk:mixes:mix a.mp3", alias "Disk:mixes:mix b.mp3"}}}`

	result := parse(input)

	assert.NotNil(t, result)
	assert.Equal(t, "mixes", result[0].name)
	assert.Equal(t, 10, result[0].ID)
	assert.Equal(t, 2, len(result[0].songs))
	assert.Equal(t, 3, len(result[0].songs[0].pathElements))
	assert.Equal(t, 3, len(result[0].songs[1].pathElements))
}

func Test__should_update_simplest_tree(t *testing.T) {
	filescanner.FileTree = &tree.Node{}
	filescanner.AddFileToTree("root/some_album/song.mp3")

	input := `{{"root", 1}{"some_album", 2, 1}}`

	parseUpdatingTree(input, filescanner.FileTree)

	nt := filescanner.FileTree
	assert.Equal(t, 1, nt.ID)
	assert.Equal(t, 2, nt.Nodes[0].ID)
}

func Test__should_update_tree_with_multiple_subfolders(t *testing.T) {

	filescanner.FileTree = &tree.Node{}
	filescanner.AddFileToTree("root/subFolder/1.mp3")
	filescanner.AddFileToTree("root/subFolder2/2.mp3")

	input := `{{"root", 1}{"subFolder", 2, 1}{"subFolder2", 3, 1}}`

	parseUpdatingTree(input, filescanner.FileTree)

	nt := filescanner.FileTree
	assert.Equal(t, 1, nt.ID)
	assert.Equal(t, 2, nt.Nodes[0].ID)
	assert.Equal(t, 3, nt.Nodes[1].ID)
}

func Test__should_update_tree_for_complex_folder_sturcture(t *testing.T) {

}
