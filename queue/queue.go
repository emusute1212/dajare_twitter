package queue

import (
	"container/list"
	"errors"
)

var q *list.List

type TweetData struct {
	ID    string
	Tweet string
}

func Init() {
	q = list.New()
}

func Enqueue(tweetData TweetData) {
	q.PushBack(tweetData)
}

func Dequeue() (TweetData, error) {
	front := q.Front()
	if front == nil {
		return TweetData{"", ""}, errors.New("キューに値は入ってないです")
	}
	r := q.Remove(front).(TweetData)
	return r, nil
}
