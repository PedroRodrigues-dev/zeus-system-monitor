package configs

import (
	"fmt"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var RabbitChannel *amqp.Channel

func StartRabbit(rabbitUser string, rabbitPassword string, rabbitHost string, rabbitPort int) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", rabbitUser, rabbitPassword, rabbitHost, rabbitPort))

	// Checks if the rabbit is up
	if err != nil {
		log.Println("Unable to connect to rabbit")
		fmt.Println("[Unable to connect to rabbit]")
		os.Exit(0)
	}

	ch, _ := conn.Channel()
	RabbitChannel = ch
}
