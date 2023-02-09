package models

import "net/url"

type Story struct {
	username User

	descendants int32

	id int32

	kids []int32

	score int32

	time unixTime

	title string

	storyType string

	url *url.URL
}

func NewStory(username User, descendants int32, id int32, kids []int32, score int32, time unixTime, title string, storyType string, url *url.URL) *Story {
	return &Story{
		username:    username,
		descendants: descendants,
		id:          id,
		kids:        kids,
		score:       score,
		time:        time,
		title:       title,
		storyType:   storyType,
		url:         url,
	}
}
