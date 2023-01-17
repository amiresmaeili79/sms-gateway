package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/amir79esmaeili/sms-gateway/internal/cfg"
	"github.com/amir79esmaeili/sms-gateway/internal/model"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

const (
	ConfigURL = "amqp://%s:%s@%s:%s/"
)

type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

func NewRabbitMQClient(config *cfg.Config) (*RabbitMQ, error) {
	conn, err := amqp.Dial(fmt.Sprintf(ConfigURL,
		config.RabbitUser, config.RabbitPass, config.RabbitHost, config.RabbitPort))
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		conn: conn,
		ch:   ch,
		q:    q,
	}, nil
}

func (r RabbitMQ) Send(message *model.NewMessageRequest) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}
	err = r.ch.PublishWithContext(ctx,
		"",       // exchange
		r.q.Name, // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})

	if err != nil {
		return err
	}
	return nil
}

func (r RabbitMQ) Receive() (chan *model.NewMessageRequest, error) {
	messages, err := r.ch.Consume(
		r.q.Name, // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)

	if err != nil {
		return nil, err
	}

	var incomingMessages chan *model.NewMessageRequest

	go func() {
		for d := range messages {
			var newMsg *model.NewMessageRequest
			err := json.Unmarshal(d.Body, &newMsg)
			if err != nil {
				log.Println("[Warning] a message failed to unmarshal")
			}
			incomingMessages <- newMsg
		}
	}()

	return incomingMessages, nil
}

func (r RabbitMQ) Close() error {
	err := r.ch.Close()
	if err != nil {
		return err
	}
	err = r.conn.Close()
	if err != nil {
		return err
	}
	return nil
}
