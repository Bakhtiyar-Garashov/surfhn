package models

type State struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func NewState() *State {
	return &State{
		choices:  []string{"Get all", "Ask", "Top Posts"},
		cursor:   42,
		selected: make(map[int]struct{}),
	}
}
