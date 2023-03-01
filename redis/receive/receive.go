package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	host := os.Args[1]

	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	for {
		str := rdb.RPop(context.Background(), "hello")
		log.Printf("str: %#v\n", str)
		if str != nil {
			/*
				if str.Err() != nil {
					failOnError(str.Err(), "Error reading message")
				}
			*/
			if str.Val() != "" {
				log.Printf("Received a message: %s", str.Val())
			}
		}
		time.Sleep(1 * time.Second)
	}
}
