package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/amir79esmaeili/sms-gateway/internal/model"
	"github.com/amir79esmaeili/sms-gateway/internal/providers"
	"github.com/google/uuid"
)

type MessageBroker interface {
	Send(message *model.NewMessageRequest) error
	Receive() (chan *model.NewMessageRequest, error)
	Close() error
}

type Services struct {
	repository       model.MessageRepository
	messageBroker    MessageBroker
	providerRegistry *providers.ProviderRegistry
}

func NewServices(repo model.MessageRepository,
	broker MessageBroker, providerRegistry *providers.ProviderRegistry) *Services {
	return &Services{
		repository:       repo,
		messageBroker:    broker,
		providerRegistry: providerRegistry,
	}
}

func (s Services) GetMessages(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
	messages, err := s.repository.List(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(messages)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (s Services) SendNewMessage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}

	decoder := json.NewDecoder(r.Body)
	var newMsgRequest model.NewMessageRequest
	err := decoder.Decode(&newMsgRequest)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	err = s.messageBroker.Send(&newMsgRequest)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (s Services) GetProviders(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(providers.AvailableProviders)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	return
}

func (s Services) HandleSendingNewMessages() {

	incoming, err := s.messageBroker.Receive()

	if err != nil {
		log.Fatalf("could not commuincate with rabbit mq to receive messages, %v\n", err)
	}

	for m := range incoming {

		provider := s.providerRegistry.GetProvider(m.Provider)

		id, _ := uuid.NewUUID()
		newMessage := model.Message{
			Id:        id,
			Recipient: m.Recipient,
			Sender:    provider.SelectSender(),
			Provider:  provider.Name(),
			Body:      m.Body,
			Date:      time.Now(),
		}

		go func() {
			err := provider.SendSMS(&newMessage)
			if err != nil {
				log.Printf("could not send message, %v\n", err)
			}
		}()
		go func() {
			err = s.repository.Create(&newMessage, context.Background())
			if err != nil {
				log.Printf("could not save message to db, %v\n", err)
			}
		}()
	}
}
