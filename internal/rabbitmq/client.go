package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/amir79esmaeili/sms-gateway/internal/cfg"
	"github.com/amir79esmaeili/sms-gateway/internal/model"
	amqp "github.com/rabbitmq/amqp091-go"
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

func (r RabbitMQ) Send(message *model.Message) error {
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
