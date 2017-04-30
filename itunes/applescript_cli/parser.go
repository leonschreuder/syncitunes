package applescript_cli

type itunesObject struct {
	name     string
	ID       string
	parentID string
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
			if lastObj.ID == "" {
				objs[len(objs)-1].ID = it.val
			} else {
				objs[len(objs)-1].parentID = it.val
			}
		}
	}
	return objs
}
