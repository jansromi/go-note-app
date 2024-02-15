package core

import "strconv"

// A class for a single item
type Item struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func (i *Item) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"id":      i.ID,
		"title":   i.Title,
		"desc":    i.Desc,
		"content": i.Content,
	}

}

func (i Item) String() string {
	return strconv.Itoa(i.ID) + " " + i.Title
}

func NewItem(id int, title, desc, content string) *Item {
	return &Item{id, title, desc, content}
}
