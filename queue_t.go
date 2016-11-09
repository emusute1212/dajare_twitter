package main

import (
	"fmt"
	"time"

	"github.com/emusute1212/dajare_twitter/queue"
)

func main() {
	queue.Init()

	queue.Enqueue("test1")
	queue.Enqueue("test2")
	queue.Enqueue("test3")
	queue.Enqueue("test4")
	queue.Enqueue("test5")
	queue.Enqueue("あいうえお")
	for {
		s := queue.Dequeue()
		if s == "" {
			break
		}
		fmt.Println(s)
		time.Sleep(1000 * time.Millisecond)
	}
}
