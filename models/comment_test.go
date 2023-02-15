package models

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewComment(t *testing.T) {
	expected := &Comment{
		Username:  "Road Runner",
		Id:        2,
		Kids:      []int32{1, 2, 3},
		Parent:    3,
		Text:      "HN: <h1>Hello world</h1>",
		Time:      58378738,
		EntryType: "comment",
	}

	got := NewComment("Road Runner", 2, []int32{1, 2, 3}, 3, "HN: <h1>Hello world</h1>", 58378738)

	if !cmp.Equal(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}

}
