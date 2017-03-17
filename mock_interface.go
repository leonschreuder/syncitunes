package main

import "fmt"

var resultNode *mockNode

type itemType int

var count = 0

type mockNode struct {
	name      string
	id        int
	parent    int
	kind      itemType
	mockNodes []*mockNode
}

func printMockTree(n *mockNode, depth int) {
	var indent []byte
	for i := 0; i < depth; i++ {
		indent = append(indent, []byte(".")...)
	}
	fmt.Println(string(indent) + n.name)
	for _, subN := range n.mockNodes {
		printMockTree(subN, depth+1)
	}
}

func (m *mockNode) String() string {
	var children []string
	if len(m.mockNodes) > 0 {
		for _, child := range m.mockNodes {
			children = append(children, child.String())
		}
	}
	return fmt.Sprintf("[n:%q id:%d ch:%s]", m.name, m.id, children)
}

const (
	f = iota //file
	d        //dir
	p        //playlist
)

func addNode(name string, t itemType, parent int) int {
	count++
	newNode := &mockNode{name: name, id: count, kind: t}
	parentNode := findParent(resultNode, parent)
	if parentNode == nil || parentNode.name == "" {
		if resultNode.name == "" && len(resultNode.mockNodes) < 1 {
			resultNode = newNode
		} else if resultNode.name != "" {
			resultNode = &mockNode{mockNodes: []*mockNode{resultNode, newNode}}
		} else {
			resultNode.mockNodes = append(resultNode.mockNodes, newNode)
		}
	} else {
		parentNode.mockNodes = append(parentNode.mockNodes, newNode)
	}
	return count
}

func findParent(currentNode *mockNode, parentID int) *mockNode {
	if currentNode.id == parentID {
		return currentNode
	}
	for _, n := range currentNode.mockNodes {
		result := findParent(n, parentID)
		if result != nil {
			return result
		}
	}
	return nil
}

type mockInterface struct {
	currentID   int
	pathCreated string
}

func (i *mockInterface) NewFolder(name string, id int) (int, error) {
	return addNode(name, d, id), nil
}
func (i *mockInterface) NewPlaylist(name string, id int) (int, error) {
	return addNode(name, p, id), nil
}
func (mockInterface) GetPlaylistIDByName(name string) (int, error) {
	return -1, nil
}
func (mockInterface) GetParentIDForPlaylist(id int) (int, error) {
	return -1, nil
}
func (i *mockInterface) AddFileToPlaylist(filePath string, playlistID int) (int, error) {
	return addNode(filePath, f, playlistID), nil
}
func (mockInterface) DeletePlaylistByID(id int) error {
	return nil
}
