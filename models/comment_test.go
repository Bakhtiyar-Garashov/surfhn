package models

import (
	"reflect"
	"testing"
)

func TestNewComment(t *testing.T) {
	expected := &Comment{
		username:  "Road Runner",
		id:        2,
		kids:      []int32{1, 2, 3},
		parent:    3,
		text:      "HN: <h1>Hello world</h1>",
		time:      58378738,
		entryType: "comment",
	}

	got := NewComment("Road Runner", 2, []int32{1, 2, 3}, 3, "HN: <h1>Hello world</h1>", 58378738)

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}

}
