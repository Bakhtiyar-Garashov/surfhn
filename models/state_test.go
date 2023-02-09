package models

import (
	"reflect"
	"testing"
)

func TestNewState(t *testing.T) {
	expected := &State{
		choices:  []string{"Foo", "Bar", "Baz"},
		cursor:   42,
		selected: map[int]struct{}{},
	}

	got := NewState()
	// Do we really need DeepEqual here ? TODO
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected %v, got %v", expected, got)
	}

}
