package providers

import "github.com/amir79esmaeili/sms-gateway/internal/model"

type SMSProvider interface {
	SendSMS(message *model.Message) error
	Name() string
	SelectSender() string
}

var AvailableProviders = []model.Providers{
	{
		Name:        KavehNegarStr,
		PersianName: KavehNegarStrFarsi,
	},
}
