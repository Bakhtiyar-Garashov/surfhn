package models

import (
	"reflect"
	"testing"
)

func TestNewAskHN(t *testing.T) {
	expected := &AskHN{
		username:    "Road Runner",
		descendants: 1,
		id:          2,
		kids:        []int32{1, 2, 3},
		score:       3,
		text:        "HN: <h1>Hello world</h1>",
		time:        58378738,
		title:       "Ask HN: lorem ipsum",
		entryType:   "story",
	}

	got := NewAskHN("Road Runner", 1, 2, []int32{1, 2, 3}, 3, "HN: <h1>Hello world</h1>", 58378738, "Ask HN: lorem ipsum")

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}

}
