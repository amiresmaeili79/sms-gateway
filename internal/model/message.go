package model

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Message struct {
	Id        uuid.UUID `json:"id,omitempty"`
	Recipient string    `json:"recipient,omitempty"`
	Sender    string    `json:"sender,omitempty"`
	Body      string    `json:"body,omitempty"`
	Date      time.Time `json:"date,omitempty"`
	Provider  string    `json:"provider,omitempty"`
}

type MessageRepository interface {
	List(ctx context.Context) ([]*Message, error)
	Get(id uuid.UUID, ctx context.Context) (*Message, error)
	Create(entity *Message, ctx context.Context) error
}

type NewMessageRequest struct {
	Recipient string `json:"recipient,omitempty"`
	Body      string `json:"body,omitempty"`
	Provider  string `json:"provider,omitempty"`
}
