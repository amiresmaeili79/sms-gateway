package model

type Message struct {
	Recipient string `json:"recipient,omitempty"`
	Sender    string `json:"sender,omitempty"`
	Body      string `json:"body,omitempty"`
}

type MessageRepository interface {
}
