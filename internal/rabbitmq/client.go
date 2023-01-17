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
	connectionUrl string
	topic         string
	conn          *amqp.Connection
	ch            *amqp.Channel
	q             *amqp.Queue
}

func NewRabbitMQClient(config *cfg.Config) (*RabbitMQ, error) {
	url := fmt.Sprintf(ConfigURL,
		config.RabbitUser, config.RabbitPass, config.RabbitHost, config.RabbitPort)

	client := &RabbitMQ{
		connectionUrl: url,
		topic:         config.RabbitTopic,
		conn:          nil,
		ch:            nil,
		q:             nil,
	}

	err := client.Connect()
	if err != nil {
		return nil, err
	}

	return client, err

}

func (r *RabbitMQ) Connect() error {
	conn, err := amqp.Dial(r.connectionUrl)
	if err != nil {
		return err
	}
	r.conn = conn

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	r.ch = ch

	q, err := ch.QueueDeclare(
		r.topic, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return err
	}

	r.q = &q

	log.Println("[INFO] RabbitMQ connected")
	return nil
}

func (r *RabbitMQ) Send(message *model.NewMessageRequest) error {

	if r.conn.IsClosed() {
		err := r.Connect()
		if err != nil {
			return err
		}
	}
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

func (r *RabbitMQ) Receive() (chan *model.NewMessageRequest, error) {

	if r.conn.IsClosed() {
		err := r.Connect()
		if err != nil {
			return nil, err
		}
	}

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

	incoming := make(chan *model.NewMessageRequest)
	go func() {
		for d := range messages {
			if r.conn.IsClosed() {
				err := r.Connect()
				if err != nil {
					return
				}
			}
			var newMsg *model.NewMessageRequest
			err := json.Unmarshal(d.Body, &newMsg)
			if err != nil {
				log.Println("[Warning] a message failed to unmarshal")
			}
			incoming <- newMsg
		}
	}()

	return incoming, err
}

func (r *RabbitMQ) Close() error {
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
