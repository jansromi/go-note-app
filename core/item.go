package core

import "strconv"

// A class for a single item
type Item struct {
	id                   int
	title, desc, content string
}

func (i Item) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"id":      i.id,
		"title":   i.title,
		"desc":    i.desc,
		"content": i.content,
	}

}

func (i Item) String() string {
	return strconv.Itoa(i.id) + " " + i.title
}

func NewItem(id int, title, desc, content string) *Item {
	return &Item{id, title, desc, content}
}
