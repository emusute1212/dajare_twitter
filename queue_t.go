package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/emusute1212/dajare_twitter/queue"
)

func main() {
	//キューの初期化
	queue.Init()

	//キューに値のセット
	queue.Enqueue(queue.TweetData{"test1", "testT1"})
	queue.Enqueue(queue.TweetData{"test2", "testT2"})
	t3 := queue.TweetData{"test3", "testT3"}
	queue.Enqueue(t3)
	t4 := queue.TweetData{"test4", "testT4"}
	queue.Enqueue(t4)

	//キューから値を取り出す
	for s, e := queue.Dequeue(); e == nil; s, e = queue.Dequeue() {
		fmt.Println(reflect.TypeOf(s.Tweet), reflect.TypeOf(s.ID))
		time.Sleep(1000 * time.Millisecond)
	}
}
