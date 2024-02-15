package core

import (
	"encoding/json"
	"fmt"
	"os"
)

// this class is responsible for managing items and file operations

type Items struct {
	items []*Item
	id    int
}

// Create an instance of Items
func NewItems() *Items {
	return &Items{}
}

func NewItemsWithFilePath(filePath string) *Items {
	itemSlice := ParseItemsFromFile(filePath)
	if itemSlice == nil {
		return nil
	}
	maxId := getHighestId(itemSlice)
	return &Items{items: itemSlice, id: maxId}
}

func NewItemsWithItems(items []*Item) *Items {
	return &Items{items: items}
}

func getHighestId(items []*Item) int {
	highestId := 0
	for _, v := range items {
		if v.ID > highestId {
			highestId = v.ID
		}
	}
	return highestId
}

// Add a new item
func (i *Items) AddItem(title, desc, content string) {
	i.id++
	i.items = append(i.items, NewItem(i.id, title, desc, content))
}

// Delete an item
func (i *Items) DeleteItem(id int) {
	for index, v := range i.items {
		if v.ID == id {
			// create two slices, one from the start to the index and one from the index to the end.
			// then append the two slices together to remove the item at the index
			i.items = append(i.items[:index], i.items[index+1:]...)
			break
		}
	}
}

func (i *Items) GetItemsForListView() []*Item {
	return i.items
}

// Print all items
func (i *Items) PrintItems() {
	for _, v := range i.items {
		fmt.Println(v)
	}
}

func ParseItemsFromFile(filepath string) []*Item {
	openFile, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer openFile.Close()

	items := []*Item{}
	decoder := json.NewDecoder(openFile)
	err = decoder.Decode(&items)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return items
}
