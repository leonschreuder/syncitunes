package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test__should_lex_root_item(t *testing.T) {
	l := lex("{\"bla\", 19}")

	assert.Equal(t, item{plainText, ""}, <-l.items)
	assert.Equal(t, item{opener, "{"}, <-l.items)
	assert.Equal(t, item{openQuote, "\""}, <-l.items)
	assert.Equal(t, item{itemName, "bla"}, <-l.items)
	assert.Equal(t, item{closeQuote, "\""}, <-l.items)
	assert.Equal(t, item{separator, ", "}, <-l.items)
	assert.Equal(t, item{itemId, "19"}, <-l.items)
	assert.Equal(t, item{closer, "}"}, <-l.items)
}

func Test__should_lex_complete_item(t *testing.T) {
	l := lex("{\"bla\", 19, 20}")

	assert.Equal(t, item{plainText, ""}, <-l.items)
	assert.Equal(t, item{opener, "{"}, <-l.items)
	assert.Equal(t, item{openQuote, "\""}, <-l.items)
	assert.Equal(t, item{itemName, "bla"}, <-l.items)
	assert.Equal(t, item{closeQuote, "\""}, <-l.items)
	assert.Equal(t, item{separator, ", "}, <-l.items)
	assert.Equal(t, item{itemId, "19"}, <-l.items)
	assert.Equal(t, item{separator, ", "}, <-l.items)
	assert.Equal(t, item{itemId, "20"}, <-l.items)
	assert.Equal(t, item{closer, "}"}, <-l.items)
}

func Test__should_lex_multiple_items(t *testing.T) {
	l := lex("{{\"bla\"}, {\"B\", 19}}")

	assert.Equal(t, item{plainText, ""}, <-l.items)
	assert.Equal(t, item{opener, "{"}, <-l.items)
	assert.Equal(t, item{opener, "{"}, <-l.items)
	assert.Equal(t, item{openQuote, "\""}, <-l.items)
	assert.Equal(t, item{itemName, "bla"}, <-l.items)
	assert.Equal(t, item{closeQuote, "\""}, <-l.items)
	assert.Equal(t, item{closer, "}"}, <-l.items)
	assert.Equal(t, item{separator, ", "}, <-l.items)
	assert.Equal(t, item{opener, "{"}, <-l.items)
	assert.Equal(t, item{openQuote, "\""}, <-l.items)
	assert.Equal(t, item{itemName, "B"}, <-l.items)
	assert.Equal(t, item{closeQuote, "\""}, <-l.items)
	assert.Equal(t, item{separator, ", "}, <-l.items)
	assert.Equal(t, item{itemId, "19"}, <-l.items)
	assert.Equal(t, item{closer, "}"}, <-l.items)
	assert.Equal(t, item{closer, "}"}, <-l.items)
}
