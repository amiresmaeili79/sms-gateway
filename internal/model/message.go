package model

import "time"

type Message struct {
	Recipient string    `json:"recipient,omitempty"`
	Sender    string    `json:"sender,omitempty"`
	Body      string    `json:"body,omitempty"`
	Date      time.Time `json:"date,omitempty"`
}
