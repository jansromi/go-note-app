package core

type App struct {
	items *Items
}

func NewApp() *App {
	return &App{items: NewItems()}
}

func (a *App) AddItem(title, desc, content string) {
	a.items.AddItem(title, desc, content)
}
