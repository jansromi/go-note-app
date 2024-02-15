package core

import (
	"fmt"
	"os"
)

type App struct {
	items    *Items
	settings *Settings
}

// Create an instance of App
func NewApp() *App {
	settings, err := NewSettings()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	itemsFilePath := settings.ItemPath
	if _, err := os.Stat(itemsFilePath); os.IsNotExist(err) {
		// Create a new items file
		file, err := os.Create(itemsFilePath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()
	}

	return &App{items: NewItemsWithFilePath(settings.ItemPath), settings: settings}
}

func (a *App) AddItem(title, desc, content string) {
	a.items.AddItem(title, desc, content)
}

func (a *App) PrintItems() {
	a.items.PrintItems()
}

func (a *App) GetItemsForListView() []*Item {
	return a.items.GetItemsForListView()
}
