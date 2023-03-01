package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
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
	messageCount, err := strconv.Atoi(os.Args[2])
	failOnError(err, "Failed to parse second arg as messageCount : int")

	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	queueName := "hello"

	log.SetFlags(0)
	time.Sleep(1 * time.Second)
	log.Printf("Sending %d messages...\n", messageCount)

	for i := 0; i < messageCount; i++ {
		msg := fmt.Sprintf("Hello World: %d", i)

		intCmd := rdb.RPush(context.Background(), queueName, msg)
		if intCmd != nil && intCmd.Err() != nil {
			failOnError(intCmd.Err(), "Can't send")
			os.Exit(1)
		}
	}

	time.Sleep(500 * time.Millisecond)
	log.Printf("Done, now watching...\n")

	oldSize := int64(-1)
	for oldSize != 0 {
		len := rdb.LLen(context.Background(), queueName)
		if len != nil && len.Err() != nil {
			failOnError(len.Err(), "Failed to get len")
		}
		if oldSize != len.Val() {
			log.Printf("Queue size: %d\n", len.Val())
			oldSize = len.Val()
		}
		time.Sleep(1 * time.Second)
	}

	log.Printf("All done!\n")
}
