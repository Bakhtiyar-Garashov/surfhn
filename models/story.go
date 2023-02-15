package models

// A story: https://hacker-news.firebaseio.com/v0/item/8863.json?print=pretty

//	{
//	  "by" : "dhouston",
//	  "descendants" : 71,
//	  "id" : 8863,
//	  "kids" : [ 8952, 9224, 8917, 8884, 8887, 8943, 8869, 8958, 9005, 9671, 8940, 9067, 8908, 9055, 8865, 8881, 8872, 8873, 8955, 10403, 8903, 8928, 9125, 8998, 8901, 8902, 8907, 8894, 8878, 8870, 8980, 8934, 8876 ],
//	  "score" : 111,
//	  "time" : 1175714200,
//	  "title" : "My YC app: Dropbox - Throw away your USB drive",
//	  "type" : "story",
//	  "url" : "http://www.getdropbox.com/u/2/screencast.html"
//	}

type Story struct {
	Username string

	Descendants int32

	Id int32

	Kids []int32

	Score int32

	Time UnixTime

	Title string

	EntryType string

	Url string
}

func NewStory(username string, descendants int32, id int32, kids []int32, score int32, time UnixTime, title string, url string) *Story {
	return &Story{
		Username:    username,
		Descendants: descendants,
		Id:          id,
		Kids:        kids,
		Score:       score,
		Time:        time,
		Title:       title,
		EntryType:   "story",
		Url:         url,
	}
}
