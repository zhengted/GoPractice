package main

import (
	"GoPractice/Routines/PubAndSubs/pubsub"
	"fmt"
	"strings"
	"time"
)

func main() {
	p := pubsub.NewPublisher(5*time.Second, 64)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello,world")
	p.Publish("hello,golang")

	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	time.Sleep(3 * time.Second)
}
