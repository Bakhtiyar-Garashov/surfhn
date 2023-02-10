package models

import (
	"reflect"
	"testing"
)

func TestNewJob(t *testing.T) {
	expected := &Job{
		username:  "Road Runner",
		id:        4,
		score:     878,
		text:      "Lorem ipsum dolor sit amet",
		time:      1210981217,
		title:     "Acme Inc. is looking for a new marketing manager",
		entryType: "job",
		url:       "https://www.acme.inc/careers",
	}

	got := NewJob("Road Runner", 4, 878, "Lorem ipsum dolor sit amet", 1210981217, "Acme Inc. is looking for a new marketing manager", "https://www.acme.inc/careers")

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}
}
