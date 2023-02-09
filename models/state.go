package models

type State struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func NewState() *State {
	return &State{
		choices:  []string{"Foo", "Bar", "Baz"},
		cursor:   42,
		selected: make(map[int]struct{}),
	}
}
