package models

// comment: https://hacker-news.firebaseio.com/v0/item/2921983.json?print=pretty

// {
//   "by" : "norvig",
//   "id" : 2921983,
//   "kids" : [ 2922097, 2922429, 2924562, 2922709, 2922573, 2922140, 2922141 ],
//   "parent" : 2921506,
//   "text" : "Aw shucks, guys ... you make me blush with your compliments.<p>Tell you what, Ill make a deal: I'll keep writing if you keep reading. K?",
//   "time" : 1314211127,
//   "type" : "comment"
// }

type Comment struct {
	username string

	id int32

	kids []int32

	parent int32

	text string

	time UnixTime

	entryType string
}

func NewComment(username string, id int32, kids []int32, parent int32, text string, time UnixTime) *Comment {
	return &Comment{
		username:  username,
		id:        id,
		kids:      kids,
		parent:    parent,
		text:      text,
		time:      time,
		entryType: "comment",
	}
}