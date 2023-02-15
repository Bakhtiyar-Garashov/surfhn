package models

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewState(t *testing.T) {
	expected := &State{
		Choices:  []string{"Foo", "Bar", "Baz"},
		Cursor:   42,
		Selected: map[int]struct{}{},
	}

	got := NewState([]string{"Foo", "Bar", "Baz"}, 42, map[int]struct{}{})
	// Do we really need DeepEqual here ? TODO
	if !cmp.Equal(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}

}
