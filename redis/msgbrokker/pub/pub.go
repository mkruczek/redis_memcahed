package main

import (
	"fmt"
	"log"
	"redismemcache/redis"
	"time"
)

func main() {

	c, err := redis.New()
	if err != nil {
		log.Fatal(err)
	}

	i := 0

	for {
		time.Sleep(5 * time.Second)
		i++
		c.PublishMessage("cache.update.now", fmt.Sprintf("update := %d", i))
		c.PublishMessage("cache.restart.now", fmt.Sprintf("restart := %d", i))
	}
}
