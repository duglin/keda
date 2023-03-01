package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	queueName := "hello"
	url := os.Args[1]
	messageCount, err := strconv.Atoi(os.Args[2])
	failOnError(err, "Failed to parse second arg as messageCount : int")
	if len(os.Args) > 3 {
		queueName = os.Args[3]
	}
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	log.SetFlags(0)
	time.Sleep(1 * time.Second)
	log.Printf("Sending %d messages...\n", messageCount)

	for i := 0; i < messageCount; i++ {
		body := fmt.Sprintf("Hello World: %d", i)
		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		// log.Printf(" [x] Sent %q", body)
		failOnError(err, "Failed to publish a message")
	}

	time.Sleep(500 * time.Millisecond)
	log.Printf("Done, now watching...\n")

	oldSize := -1
	for oldSize != 0 {
		q, err := ch.QueueInspect(queueName)
		failOnError(err, "Failed to inspect queue")
		if oldSize != q.Messages {
			log.Printf("Queue size: %d\n", q.Messages)
			oldSize = q.Messages
		}
		time.Sleep(1 * time.Second)
	}

	log.Printf("All done!\n")
}
