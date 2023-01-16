package providers

import (
	"github.com/amir79esmaeili/sms-gateway/internal/cfg"
	"github.com/amir79esmaeili/sms-gateway/internal/model"
)

const KavehNegarStr = "KavehNegar"
const KavehNegarStrFarsi = "کاوه نگار"

type KavehNegar struct {
	apiKey string
	name   string
}

func NewKavehNegarClient(config *cfg.Config) *KavehNegar {
	return &KavehNegar{
		apiKey: config.KavehNegarAPIKey,
		name:   KavehNegarStr,
	}
}

func (k KavehNegar) sendSMS(message *model.Message) error {
	return nil
}

func (k KavehNegar) Name() string {
	return k.name
}
