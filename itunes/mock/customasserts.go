package mock

import (
	"testing"

	"github.com/meonlol/syncitunes/itunes"
	"github.com/stretchr/testify/assert"
)

// checks supplied indexMapping exists and contains an item with specified name and type
func AssertTreeMapHasNameAndType(t *testing.T, target *MockNode, indexMapping []int, name string, typ itunes.ItemType) {
	// target := mock.MockTree
	for _, i := range indexMapping {
		if len(target.MockNodes) > i {
			target = target.MockNodes[i]
		} else {
			t.Errorf("requested Node[%d], but %q has only %d child nodes", i, target.Name, len(target.MockNodes))
			t.Fail()
		}
	}
	assert.Equal(t, name, target.Name)
	assert.Equal(t, typ, target.Kind, "expected different itunes item type.")
}
