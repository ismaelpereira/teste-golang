package rabbit

import (
	"fmt"

	"github.com/streadway/amqp"
)

func SendMessage(message string) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5673/")

	if err != nil {
		return fmt.Errorf("Failed to connect to RabbitMQ", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return fmt.Errorf("Failed to open a channel", err)
	}

	defer ch.Close()

	queue, err := ch.QueueDeclare(
		"api_messages",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return fmt.Errorf("Failed to declare queue", err)
	}

	body := message

	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err != nil {
		return fmt.Errorf("Failed to publish a message", err)
	}

	return nil
}
