package queue

import "container/list"

var q *list.List

func Init() {
	q = list.New()
}

func Enqueue(text string) {
	q.PushBack(text)
}

func Dequeue() string {
	front := q.Front()
	if front == nil {
		return ""
	}
	r := q.Remove(front).(string)
	return r
}
