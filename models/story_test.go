package models

import (
	"reflect"
	"testing"
)

func TestNewStort(t *testing.T) {
	expected := &Story{
		username:    "Road Runner",
		descendants: 3,
		id:          4,
		kids:        []int32{4, 5, 2, 24},
		score:       111,
		time:        1175714200,
		title:       "My YC app: Dropbox - Throw away your USB drive",
		entryType:   "story",
		url:         "http://www.getdropbox.com/u/2/screencast.html",
	}

	got := NewStory("Road Runner", 3, 4, []int32{4, 5, 2, 24}, 111, 1175714200, "My YC app: Dropbox - Throw away your USB drive", "http://www.getdropbox.com/u/2/screencast.html")

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
