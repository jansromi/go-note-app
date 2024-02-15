package core

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func TestAddItem(t *testing.T) {
	items := NewItems()
	items.AddItem("Test Item", "This is a test item", "Lorem ipsum dolor sit amet")

	expected := []*Item{
		{
			ID:      1,
			Title:   "Test Item",
			Desc:    "This is a test item",
			Content: "Lorem ipsum dolor sit amet",
		},
	}

	if len(items.items) != len(expected) {
		t.Errorf("AddItem() failed, expected %d item(s), got %d", len(expected), len(items.items))
	}

	if len(items.items) > 0 && !reflect.DeepEqual(items.items[0], expected[0]) {
		t.Errorf("AddItem() failed, expected item %v, got %v", expected[0], items.items[0])
	}
}

func TestAddManyItemsAndTestForCorrectId(t *testing.T) {
	items := NewItems()
	items.AddItem("Test Item 1", "This is a test item 1", "Lorem ipsum dolor sit amet")
	items.AddItem("Test Item 2", "This is a test item 2", "Lorem ipsum dolor sit amet")
	items.AddItem("Test Item 3", "This is a test item 3", "Lorem ipsum dolor sit amet")

	if items.items[0].ID != 1 {
		t.Errorf("AddItem() failed, expected id %d, got %d", 1, items.items[0].ID)
	}

	if items.items[1].ID != 2 {
		t.Errorf("AddItem() failed, expected id %d, got %d", 2, items.items[1].ID)
	}

	if items.items[2].ID != 3 {
		t.Errorf("AddItem() failed, expected id %d, got %d", 3, items.items[2].ID)
	}
}

func TestPrintItems(t *testing.T) {
	items := NewItemsWithItems([]*Item{
		{
			ID:      1,
			Title:   "Test Item 1",
			Desc:    "This is a test item 1",
			Content: "Lorem ipsum dolor sit amet",
		},
		{
			ID:      2,
			Title:   "Test Item 2",
			Desc:    "This is a test item 2",
			Content: "Lorem ipsum dolor sit amet",
		},
		{
			ID:      3,
			Title:   "Test Item 3",
			Desc:    "This is a test item 3",
			Content: "Lorem ipsum dolor sit amet",
		},
	})

	// Redirect stdout to capture the output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	items.PrintItems()

	// Restore stdout
	w.Close()
	os.Stdout = old

	out, _ := io.ReadAll(r)

	expectedOutput := "1 Test Item 1\n2 Test Item 2\n3 Test Item 3\n"

	if string(out) != expectedOutput {
		t.Errorf("PrintItems() failed, expected output:\n%s\ngot:\n%s", expectedOutput, string(out))
	}
}
