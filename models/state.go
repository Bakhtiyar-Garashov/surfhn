package models

type State struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func NewState(choices []string, cursor int, selected map[int]struct{}) *State {
	return &State{
		choices:  choices,
		cursor:   cursor,
		selected: selected,
	}
}
