package mock

import (
	"fmt"

	"github.com/meonlol/syncitunes/itunes"
	"github.com/meonlol/syncitunes/tree"
)

var MockTree *MockNode

type MockNode struct {
	Name      string
	Id        int
	Parent    int
	Kind      itunes.ItemType
	MockNodes []*MockNode
}

var currentIdCount = 0

type MockInterface struct {
	currentID   int
	pathCreated string
}

func NewMockInterface() MockInterface {
	currentIdCount = 0
	MockTree = &MockNode{}
	return MockInterface{}
}

func (i *MockInterface) NewFolder(name string, id int) (int, error) {
	return addNode(name, itunes.Dir, id), nil
}
func (i *MockInterface) NewPlaylist(name string, id int) (int, error) {
	return addNode(name, itunes.Playlist, id), nil
}
func (MockInterface) GetPlaylistIDByName(name string) (int, error) {
	return -1, nil
}
func (MockInterface) GetParentIDForPlaylist(id int) (int, error) {
	return -1, nil
}
func (i *MockInterface) AddFileToPlaylist(filePath string, playlistID int) (int, error) {
	return addNode(filePath, itunes.File, playlistID), nil
}
func (MockInterface) DeletePlaylistByID(id int) error {
	return nil
}

func (MockInterface) UpdateTreeWithExisting(tree *tree.Node) {
}

func printMockTree(n *MockNode, depth int) {
	var indent []byte
	for i := 0; i < depth; i++ {
		indent = append(indent, []byte(".")...)
	}
	fmt.Println(string(indent) + n.Name)
	for _, subN := range n.MockNodes {
		printMockTree(subN, depth+1)
	}
}

func (m *MockNode) String() string {
	var children []string
	if len(m.MockNodes) > 0 {
		for _, child := range m.MockNodes {
			children = append(children, child.String())
		}
	}
	return fmt.Sprintf("[n:%q id:%d ch:%s]", m.Name, m.Id, children)
}

func addNode(name string, t itunes.ItemType, parent int) int {
	currentIdCount++
	newNode := &MockNode{Name: name, Id: currentIdCount, Kind: t}
	parentNode := findParent(MockTree, parent)
	if parentNode == nil || parentNode.Name == "" {
		if rootNotSet() {
			MockTree = newNode
		} else if rootSetWithNormalNode() {
			addSecondRoot(newNode)
		} else {
			addChildToNode(newNode, MockTree)
		}
	} else {
		addChildToNode(newNode, parentNode)
	}
	return currentIdCount
}

func findParent(currentNode *MockNode, parentID int) *MockNode {
	if currentNode.Id == parentID {
		return currentNode
	}
	for _, n := range currentNode.MockNodes {
		result := findParent(n, parentID)
		if result != nil {
			return result
		}
	}
	return nil
}

func rootNotSet() bool {
	return MockTree.Name == "" && len(MockTree.MockNodes) < 1
}

func rootSetWithNormalNode() bool {
	return MockTree.Name != ""
}

func addChildToNode(child *MockNode, parent *MockNode) {
	parent.MockNodes = append(parent.MockNodes, child)
}

func addSecondRoot(n *MockNode) {
	MockTree = &MockNode{MockNodes: []*MockNode{MockTree, n}}
}
