package cmd

import (
	"log"

	"github.com/amir79esmaeili/sms-gateway/internal/postgres"
	"github.com/amir79esmaeili/sms-gateway/internal/providers"
	"github.com/amir79esmaeili/sms-gateway/internal/rabbitmq"
	"github.com/amir79esmaeili/sms-gateway/internal/repository"
	"github.com/amir79esmaeili/sms-gateway/internal/service"
	"github.com/spf13/cobra"
)

func addConsumeCmd(root *cobra.Command) {
	consumeCmd := &cobra.Command{
		Use:   "consume",
		Short: "Starts a consumer to receive messages and send them",
		Long:  "This command starts a consumer that receives messages from RabbitMQ and sends them to the requested provided",
		Run: func(cmd *cobra.Command, args []string) {
			consume(cmd)
		},
	}

	root.AddCommand(consumeCmd)
	consumeCmd.Flags().StringP("cfg", "c", ".env", "Config file path")
}

func consume(cmd *cobra.Command) {
	config := loadConfig(cmd)

	db, err := postgres.ConnectToDB(&config)
	if err != nil {
		log.Fatalf("Could not connect to db, %v", err)
	}
	rabbit, err := rabbitmq.NewRabbitMQClient(&config)
	if err != nil {
		log.Fatalf("Could not connect to db, %v", err)
	}

	msgRepo := repository.NewMessageRepository(db)

	providerRegistry := providers.NewProviderRegistry(
		providers.NewKavehNegarClient(&config),
		providers.NewGhasedakProvider(&config),
	)

	msgServices := service.NewServices(msgRepo, rabbit, providerRegistry)
	msgServices.HandleSendingNewMessages()
}
