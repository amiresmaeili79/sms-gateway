package service

import (
	"encoding/json"
	"github.com/amir79esmaeili/sms-gateway/internal/model"
	"github.com/amir79esmaeili/sms-gateway/internal/providers"
	"net/http"
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
