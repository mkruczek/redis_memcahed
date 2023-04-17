package main

import (
	"fmt"
	"log"
	"redismemcache/redis"
)

func main() {

	c, err := redis.New()
	if err != nil {
		log.Fatal(err)
	}

	pubsub, err := c.SubscribeMessage("cache.restart.now")
	if err != nil {
		panic(err)
	}

	go func() {
		for {
			msg, err := pubsub.ReceiveMessage()
			if err != nil {
				panic(err)
			}
			fmt.Println("Received message:", msg.Payload)
		}
	}()

	select {}
}
