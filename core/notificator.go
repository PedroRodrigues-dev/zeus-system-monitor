package core

import (
	"context"
	"log"
	"time"
	"zeus/configs"
	"zeus/scripts"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Notificator(msg string) {
	logging(msg)
	rabbit(msg)
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
