package core

import "fmt"

// this class is responsible for managing items and file operations

type Items struct {
	items []*Item
	id    int
}

func NewItems() *Items {
	return &Items{}
}

func (i *Items) AddItem(title, desc, content string) {
	i.id++
	i.items = append(i.items, NewItem(i.id, title, desc, content))
}

func (i *Items) PrintItems() {
	for _, v := range i.items {
		fmt.Println(v)
	}
}
