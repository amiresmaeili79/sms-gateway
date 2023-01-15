package service

import (
	"encoding/json"
	"github.com/amir79esmaeili/sms-gateway/internal/model"
	"net/http"
)

type Services struct {
	repository model.MessageRepository
}

func NewServices(repo model.MessageRepository) *Services {
	return &Services{
		repository: repo,
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