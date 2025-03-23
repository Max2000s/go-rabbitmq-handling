package main

import (
	"github.com/Max2000s/go-rabbitmq-handling/pkg/connectors"
	"log"
)

func main() {
	log.Println(" === Starting the Program now ===")
	log.Println("Starting the channel for receiving!")
	go connectors.Receive()
	log.Println("Will now send several messages!")
	connectors.Send()
	connectors.Send()
	connectors.Send()
	connectors.Send()
	log.Println("Program wil end now!")
}
