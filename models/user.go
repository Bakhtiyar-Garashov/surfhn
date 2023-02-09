package models

// https://github.com/HackerNews/API
// UNIX timestamp, refers to the number of seconds that have elapsed since the epoch. 1 Jan, 1970.
type unixTime int64

type User struct {
	id int32

	createdAt unixTime

	karma int32

	about string

	submitted []int32
}

func NewUser(id int32, createdAt unixTime, karma int32, about string, submitted []int32) *User {
	return &User{
		id:        id,
		createdAt: createdAt,
		karma:     karma,
		about:     about,
		submitted: submitted,
	}
}
