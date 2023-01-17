package providers

import (
	"github.com/amir79esmaeili/sms-gateway/internal/cfg"
	"github.com/amir79esmaeili/sms-gateway/internal/model"
)

const GhasedakStr = "Ghasedak"
const GhasedakStrFarsi = "قاصدک"

type Ghasedak struct {
	name string
}

func NewGhasedakProvider(config *cfg.Config) *Ghasedak {
	return &Ghasedak{
		name: GhasedakStr,
	}
}

func (g Ghasedak) SendSMS(message *model.Message) error {
	return nil
}

func (g Ghasedak) Name() string {
	return g.name
}

func (g Ghasedak) SelectSender() string {
	return "11111111"
}
