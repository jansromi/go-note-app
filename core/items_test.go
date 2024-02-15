package core

import (
	"io/ioutil"
	"os"
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
func TestPrintItems(t *testing.T) {
	items := &Items{
		items: []*Item{
			{
				id:      1,
				title:   "Test Item 1",
				desc:    "This is a test item 1",
				content: "Lorem ipsum dolor sit amet",
			},
			{
				id:      2,
				title:   "Test Item 2",
				desc:    "This is a test item 2",
				content: "Lorem ipsum dolor sit amet",
			},
			{
				id:      3,
				title:   "Test Item 3",
				desc:    "This is a test item 3",
				content: "Lorem ipsum dolor sit amet",
			},
		},
	}

	// Redirect stdout to capture the output
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	items.PrintItems()

	// Restore stdout
	w.Close()
	os.Stdout = old

	out, _ := ioutil.ReadAll(r)

	expectedOutput := "1 Test Item 1\n2 Test Item 2\n3 Test Item 3\n"

	if string(out) != expectedOutput {
		t.Errorf("PrintItems() failed, expected output:\n%s\ngot:\n%s", expectedOutput, string(out))
	}
}
