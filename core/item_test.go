package core

import (
	"reflect"
	"testing"
)

func TestToJson(t *testing.T) {
	item := &Item{
		id:      1,
		title:   "Test Item",
		desc:    "This is a test item",
		content: "Lorem ipsum dolor sit amet",
	}

	expected := map[string]interface{}{
		"id":      1,
		"title":   "Test Item",
		"desc":    "This is a test item",
		"content": "Lorem ipsum dolor sit amet",
	}

	result := item.ToJson()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ToJson() = %v, want %v", result, expected)
	}
}
