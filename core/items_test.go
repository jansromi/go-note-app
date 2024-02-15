package core

import (
	"reflect"
	"testing"
)

func TestAddItem(t *testing.T) {
	items := &Items{}
	items.AddItem("Test Item", "This is a test item", "Lorem ipsum dolor sit amet")

	expected := &Items{
		id: 1,
		items: []*Item{
			{
				id:      1,
				title:   "Test Item",
				desc:    "This is a test item",
				content: "Lorem ipsum dolor sit amet",
			},
		},
	}

	if len(items.items) != len(expected.items) {
		t.Errorf("AddItem() failed, expected %d item(s), got %d", len(expected.items), len(items.items))
	}

	if items.id != expected.id {
		t.Errorf("AddItem() failed, expected id %d, got %d", expected.id, items.id)
	}

	if len(items.items) > 0 && !reflect.DeepEqual(items.items[0], expected.items[0]) {
		t.Errorf("AddItem() failed, expected item %v, got %v", expected.items[0], items.items[0])
	}
}

func TestAddManyItemsAndTestForCorrectId(t *testing.T) {
	items := &Items{}
	items.AddItem("Test Item 1", "This is a test item 1", "Lorem ipsum dolor sit amet")
	items.AddItem("Test Item 2", "This is a test item 2", "Lorem ipsum dolor sit amet")
	items.AddItem("Test Item 3", "This is a test item 3", "Lorem ipsum dolor sit amet")

	if items.items[0].id != 1 {
		t.Errorf("AddItem() failed, expected id %d, got %d", 1, items.items[0].id)
	}

	if items.items[1].id != 2 {
		t.Errorf("AddItem() failed, expected id %d, got %d", 2, items.items[1].id)
	}

	if items.items[2].id != 3 {
		t.Errorf("AddItem() failed, expected id %d, got %d", 3, items.items[2].id)
	}
}
