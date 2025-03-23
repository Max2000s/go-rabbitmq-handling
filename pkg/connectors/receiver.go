package connectors

import (
	"log"

	"github.com/Max2000s/go-rabbitmq-handling/pkg/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Receive() {
	log.Println("Will run the reciever now!")

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

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	utils.FailOnError(err, "Consuming the channel failed!")

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Recieved message on channel %s with body %s", queue.Name, d.Body)
		}
	}()
	<-forever

}
