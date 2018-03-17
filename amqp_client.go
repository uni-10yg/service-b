package main

import (
	"fmt"
	"github.com/streadway/amqp"
)


type AmqpClient struct {
	conn *amqp.Connection
}

func (m *AmqpClient) connectToBroker(broker_url string) {
	if broker_url == "" {
		panic("missing broker url")
	}

	var err error
	m.conn, err = amqp.Dial(fmt.Sprintf("%s/", broker_url))
	if err != nil {
		panic("could not connect to message broker at: " + broker_url)
	}
}

func (m *AmqpClient) sendMsg(body []byte, queueName string) error {
	if m.conn == nil {
		panic("could not send message - no connected to message broker")
	}
	ch, err := m.conn.Channel()
	failOnError(err, "could not open a channel")
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "could not declare queue")

	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	fmt.Printf("message sent to queue %v: %v", queueName, body)
	return err
}

func (m *AmqpClient) close() {
	if m.conn != nil {
		m.conn.Close()
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
