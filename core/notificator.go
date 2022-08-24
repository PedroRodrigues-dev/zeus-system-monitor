package core

import (
	"context"
	"log"
	"time"
	"zeus/api/controllers"
	"zeus/configs"
	"zeus/models"
	"zeus/scripts"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Notificator(msg string) {
	logging(msg)

	// Check if rabbit is started
	if configs.RabbitStarted {
		rabbit(msg)
	}
}

func logging(msg string) {
	log.Println(msg)
}

func rabbit(msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	configs.RabbitChannel.PublishWithContext(ctx, "", scripts.ENV.RabbitQueue, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
}

func WebsocketNotificator(percent int, option int) {
	// Choose monitor
	var monitorName string
	switch {
	case option == 0:
		monitorName = "cpu"
	case option == 1:
		monitorName = "memory"
	case option == 2:
		monitorName = "disk"
	}

	// Create json response for websocket
	var message models.WebsocketMessage
	message.MonitorName = monitorName
	message.Percent = percent

	// Add json to reference position in message array
	controllers.Messages[option] = message
}
