package models

type State struct {
	Choices  []string
	Cursor   int
	Selected map[int]struct{}
}

func NewState(choices []string, cursor int, selected map[int]struct{}) *State {
	return &State{
		Choices:  choices,
		Cursor:   cursor,
		Selected: selected,
	}
}
