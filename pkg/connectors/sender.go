package connectors

import (
	"context"
	"log"
	"time"

	"github.com/Max2000s/go-rabbitmq-handling/pkg/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Send() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnError(err, "Failed to connect to rabbitmq")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"test_queue", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	utils.FailOnError(err, "Failed to declare queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body := "that is a test message"
	err = ch.PublishWithContext(ctx, "", queue.Name, false, false, amqp.Publishing{ContentType: "text.plain", Body: []byte(body)})
	utils.FailOnError(err, "Publish for message failed")
	log.Printf(" Sent %s\n", body)

}
