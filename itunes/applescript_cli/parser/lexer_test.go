package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test__should_lex_root_item(t *testing.T) {
	l := lex(`{{"bla", 19}}`)

	assert.Equal(t, item{opener, "{"}, <-l.items)
	assert.Equal(t, item{opener, "{"}, <-l.items)
	assert.Equal(t, item{openQuote, "\""}, <-l.items)
	assert.Equal(t, item{itemName, "bla"}, <-l.items)
	assert.Equal(t, item{closeQuote, "\""}, <-l.items)
	assert.Equal(t, item{separator, ","}, <-l.items)
	assert.Equal(t, item{itemId, "19"}, <-l.items)
	assert.Equal(t, item{closer, "}"}, <-l.items)
	assert.Equal(t, item{closer, "}"}, <-l.items)
}

func Test__should_lex_complete_item(t *testing.T) {
	l := lex(`{{"bla", 19, 20, {}}}`)

	assert.Equal(t, item{opener, "{"}, <-l.items)
	assert.Equal(t, item{opener, "{"}, <-l.items)
	assert.Equal(t, item{openQuote, "\""}, <-l.items)
	assert.Equal(t, item{itemName, "bla"}, <-l.items)
	assert.Equal(t, item{closeQuote, "\""}, <-l.items)
	assert.Equal(t, item{separator, ","}, <-l.items)
	assert.Equal(t, item{itemId, "19"}, <-l.items)
	assert.Equal(t, item{separator, ","}, <-l.items)
	assert.Equal(t, item{itemId, "20"}, <-l.items)
	assert.Equal(t, item{separator, ","}, <-l.items)
	assert.Equal(t, item{opener, "{"}, <-l.items)
	assert.Equal(t, item{closer, "}"}, <-l.items)
	assert.Equal(t, item{closer, "}"}, <-l.items)
	assert.Equal(t, item{closer, "}"}, <-l.items)
}

func Test__should_lex_multiple_items(t *testing.T) {
	l := lex(`{{"bla"}, {"B", 19}}`)

	assert.Equal(t, item{opener, "{"}, <-l.items)
	assert.Equal(t, item{opener, "{"}, <-l.items)
	assert.Equal(t, item{openQuote, "\""}, <-l.items)
	assert.Equal(t, item{itemName, "bla"}, <-l.items)
	assert.Equal(t, item{closeQuote, "\""}, <-l.items)
	assert.Equal(t, item{closer, "}"}, <-l.items)
	assert.Equal(t, item{separator, ","}, <-l.items)
	assert.Equal(t, item{opener, "{"}, <-l.items)
	assert.Equal(t, item{openQuote, "\""}, <-l.items)
	assert.Equal(t, item{itemName, "B"}, <-l.items)
	assert.Equal(t, item{closeQuote, "\""}, <-l.items)
	assert.Equal(t, item{separator, ","}, <-l.items)
	assert.Equal(t, item{itemId, "19"}, <-l.items)
	assert.Equal(t, item{closer, "}"}, <-l.items)
	assert.Equal(t, item{closer, "}"}, <-l.items)
}

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

func Test__should_lex_songs(t *testing.T) {
	l := lex(`{{"playlist", 1, -1, {alias "Macintosh HD:Users:someuser:music:playlist:Some Artist - Some Song.mp3"}}}`)

	assert.Equal(t, item{opener, "{"}, <-l.items)
	assert.Equal(t, item{opener, "{"}, <-l.items)
	assert.Equal(t, item{openQuote, "\""}, <-l.items)
	assert.Equal(t, item{itemName, "playlist"}, <-l.items)
	assert.Equal(t, item{closeQuote, "\""}, <-l.items)
	assert.Equal(t, item{separator, ","}, <-l.items)
	assert.Equal(t, item{itemId, "1"}, <-l.items)
	assert.Equal(t, item{separator, ","}, <-l.items)
	assert.Equal(t, item{itemId, "-1"}, <-l.items)
	assert.Equal(t, item{separator, ","}, <-l.items)
	assert.Equal(t, item{opener, "{"}, <-l.items)
	assert.Equal(t, item{aliasIndicator, "alias"}, <-l.items)
	assert.Equal(t, item{filePathOpener, "\""}, <-l.items)
	assert.Equal(t, item{filePath, "Macintosh HD:Users:someuser:music:playlist:Some Artist - Some Song.mp3"}, <-l.items)
	assert.Equal(t, item{closeQuote, "\""}, <-l.items)
	assert.Equal(t, item{closer, "}"}, <-l.items)
	assert.Equal(t, item{closer, "}"}, <-l.items)
	assert.Equal(t, item{closer, "}"}, <-l.items)
}
