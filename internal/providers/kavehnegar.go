package providers

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/amir79esmaeili/sms-gateway/internal/cfg"
	"github.com/amir79esmaeili/sms-gateway/internal/model"
)

const KavehNegarStr = "KavehNegar"
const KavehNegarStrFarsi = "کاوه نگار"

type KavehNegar struct {
	apiKey  string
	url     string
	name    string
	numbers []string
}

func NewKavehNegarClient(config *cfg.Config) *KavehNegar {
	return &KavehNegar{
		apiKey:  config.KavehNegarAPIKey,
		url:     fmt.Sprintf(config.KavehNegarURL, config.KavehNegarAPIKey),
		name:    KavehNegarStr,
		numbers: strings.Split(config.KavehNegarNumbers, ","),
	}
}

func (k KavehNegar) SendSMS(message *model.Message) error {
	req, err := http.NewRequest("GET", k.url, nil)
	if err != nil {
		return err
	}
	query := req.URL.Query()
	query.Add("receptor", message.Recipient)
	query.Add("message", message.Body)
	query.Add("sender", message.Sender)
	req.URL.RawQuery = query.Encode()

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

func (k KavehNegar) SelectSender() string {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s) // initialize local pseudorandom generator
	idx := r.Intn(len(k.numbers))
	return k.numbers[idx]
}
