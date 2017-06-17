package parser

import (
	"strconv"

	"github.com/meonlol/syncitunes/tree"
)

type itunesObject struct {
	name     string
	ID       int
	parentID int
}

func parseUpdatingTree(input string, nodeTree *tree.Node) {
	parseResult := parse(input)
	for _, v := range parseResult {
		findItemInTree(v, nodeTree)
	}
}

func findItemInTree(item itunesObject, node *tree.Node) {
	if item.parentID == 0 {
		node.ID = item.ID
		node.ParentID = item.parentID
	}
	if node.ID == item.parentID {
		//this item is parent of one we are looking for
		for _, cn := range node.Nodes {
			findItemUnderCorrectParent(item, cn)
		}
	}
}

func findItemUnderCorrectParent(item itunesObject, node *tree.Node) {
	if item.name == node.Name {
		node.ID = item.ID
		node.ParentID = item.parentID
		return
	}
}

func parse(input string) []itunesObject {
	l := lex(input)
	objs := []itunesObject{}
	for {
		it, ok := <-l.items
		if !ok {
			break
		}
		if it.typ == itemName {
			objs = append(objs, itunesObject{name: it.val})
		}
		if it.typ == itemId {
			lastObj := objs[len(objs)-1]
			n, _ := strconv.Atoi(it.val)
			if lastObj.ID == 0 {
				objs[len(objs)-1].ID = n
			} else {
				objs[len(objs)-1].parentID = n
			}
		}
	}
	return objs
}
