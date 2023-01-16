package providers

import "github.com/amir79esmaeili/sms-gateway/internal/model"

type SMSProvider interface {
	sendSMS(message *model.Message) error
	Name() string
}

var AvailableProviders = []model.Providers{
	{
		Name:        KavehNegarStr,
		PersianName: KavehNegarStrFarsi,
	},
}
