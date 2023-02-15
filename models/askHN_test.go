package models

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewAskHN(t *testing.T) {
	expected := &AskHN{
		Username:    "Road Runner",
		Descendants: 1,
		Id:          2,
		Kids:        []int32{1, 2, 3},
		Score:       3,
		Text:        "HN: <h1>Hello world</h1>",
		Time:        58378738,
		Title:       "Ask HN: lorem ipsum",
		EntryType:   "story",
	}

	got := NewAskHN("Road Runner", 1, 2, []int32{1, 2, 3}, 3, "HN: <h1>Hello world</h1>", 58378738, "Ask HN: lorem ipsum")

	if !cmp.Equal(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}

}
