package configs

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitChannel *amqp.Channel
var RabbitStarted bool = true

func StartRabbit(rabbitUser string, rabbitPassword string, rabbitHost string, rabbitPort int) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", rabbitUser, rabbitPassword, rabbitHost, rabbitPort))

	// Checks if the rabbit is up
	if err != nil {
		log.Println("Unable to connect to rabbit")
		fmt.Println("[Unable to connect to rabbit]")
		log.Println("Starting only realtime websocket and log")
		fmt.Println("[Starting only realtime websocket and log]")
		RabbitStarted = false
	} else {
		ch, _ := conn.Channel()
		RabbitChannel = ch
	}
}
