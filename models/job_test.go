package models

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewJob(t *testing.T) {
	expected := &Job{
		Username:  "Road Runner",
		Id:        4,
		Score:     878,
		Text:      "Lorem ipsum dolor sit amet",
		Time:      1210981217,
		Title:     "Acme Inc. is looking for a new marketing manager",
		EntryType: "job",
		Url:       "https://www.acme.inc/careers",
	}

	got := NewJob("Road Runner", 4, 878, "Lorem ipsum dolor sit amet", 1210981217, "Acme Inc. is looking for a new marketing manager", "https://www.acme.inc/careers")

	if !cmp.Equal(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
