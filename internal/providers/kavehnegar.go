package providers

import (
	"fmt"
	"github.com/amir79esmaeili/sms-gateway/internal/cfg"
	"github.com/amir79esmaeili/sms-gateway/internal/model"
	"net/http"
	"time"
)

const KavehNegarStr = "KavehNegar"
const KavehNegarStrFarsi = "کاوه نگار"

type KavehNegar struct {
	apiKey string
	url    string
	name   string
}

func NewKavehNegarClient(config *cfg.Config) *KavehNegar {
	return &KavehNegar{
		apiKey: config.KavehNegarAPIKey,
		url:    fmt.Sprintf(config.KavehNegarURL, config.KavehNegarAPIKey),
		name:   KavehNegarStr,
	}
}

func (k KavehNegar) sendSMS(message *model.Message) error {
	req, err := http.NewRequest("GET", k.url, nil)
	if err != nil {
		return err
	}
	query := req.URL.Query()
	query.Add("receptor", message.Recipient)
	query.Add("message", message.Body)
	query.Add("sender", message.Sender)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	do, err := client.Do(req)
	if err != nil {
		return err
	}

	if do.StatusCode != 200 {
		return fmt.Errorf("sms could not be sent")
	}
	return nil
}

func (k KavehNegar) Name() string {
	return k.name
}
