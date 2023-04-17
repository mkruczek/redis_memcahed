package main

import (
	"fmt"
	"log"
	"redismemcache/memcache"
	"redismemcache/model"
	"redismemcache/redis"
	"sync"
	"time"
)

func main() {

	mc, err := memcache.New()
	if err != nil {
		log.Fatal(err)
	}
	rd, err := redis.New()
	if err != nil {
		log.Fatal(err)
	}

	data := model.GenerateDevice(100)

	wg := sync.WaitGroup{}
	wg.Add(2)
	//memcache
	go func(data []model.Device) {
		start := time.Now()

		for _, d := range data {
			err := mc.SetDevice(d)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Printf("memcache set: %5v\n", time.Since(start))

		for _, d := range data {
			_, err := mc.GetDevice(d.ID.String())
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Printf("memcache END: %5v\n", time.Since(start))
		wg.Done()
	}(data)

	//redis
	go func(data []model.Device) {
		start := time.Now()

		for _, d := range data {
			err := rd.SetDevice(d)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Printf("redis set:    %5v\n", time.Since(start))
		for _, d := range data {
			_, err := rd.GetDevice(d.ID.String())
			if err != nil {
				fmt.Println(err)
			}
		}

		fmt.Printf("redis END:    %5v\n", time.Since(start))
		wg.Done()
	}(data)

	wg.Wait()
}
