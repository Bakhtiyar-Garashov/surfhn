package models

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewStort(t *testing.T) {
	expected := &Story{
		Username:    "Road Runner",
		Descendants: 3,
		Id:          4,
		Kids:        []int32{4, 5, 2, 24},
		Score:       111,
		Time:        1175714200,
		Title:       "How fast Road Runner is ?",
		EntryType:   "story",
		Url:         "http://www.speedtest.org/road-runner",
	}

	got := NewStory("Road Runner", 3, 4, []int32{4, 5, 2, 24}, 111, 1175714200, "How fast Road Runner is ?", "http://www.speedtest.org/road-runner")

	if !cmp.Equal(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
