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
		title:       "How fast Road Runner is ?",
		entryType:   "story",
		url:         "http://www.speedtest.org/road-runner",
	}

	got := NewStory("Road Runner", 3, 4, []int32{4, 5, 2, 24}, 111, 1175714200, "How fast Road Runner is ?", "http://www.speedtest.org/road-runner")

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
