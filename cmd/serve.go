package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/amir79esmaeili/sms-gateway/internal/middleware"
	"github.com/amir79esmaeili/sms-gateway/internal/postgres"
	"github.com/amir79esmaeili/sms-gateway/internal/rabbitmq"
	"github.com/amir79esmaeili/sms-gateway/internal/repository"
	"github.com/amir79esmaeili/sms-gateway/internal/service"
	"github.com/spf13/cobra"
)

func addServeCmd(root *cobra.Command) {
	serveCmd := &cobra.Command{
		Use:   "serve",
		Short: "Start web server to receive requests",
		Long:  "This command starts a web server which takes requests and add them to rabbit MQ.",
		Run: func(cmd *cobra.Command, args []string) {
			serve(cmd)
		},
	}

	root.AddCommand(serveCmd)
	serveCmd.Flags().StringP("cfg", "c", ".env", "Config file path")
}

func serve(cmd *cobra.Command) {
	config := loadConfig(cmd)

	db, err := postgres.ConnectToDB(&config)
	defer postgres.KillConnection(db)

	if err != nil {
		log.Fatalf("Could not connect to db, %v", err)
	}
	rabbit, err := rabbitmq.NewRabbitMQClient(&config)
	if err != nil {
		log.Fatalf("Could not connect to db, %v", err)
	}
	defer rabbit.Close()

	msgRepo := repository.NewMessageRepository(db)
	msgServices := service.NewServices(msgRepo, rabbit, nil)

	mux := http.NewServeMux()

	mux.Handle("/messages", middleware.LogCurrentRequest(
		http.HandlerFunc(msgServices.GetMessages)))

	mux.Handle("/new-message", middleware.LogCurrentRequest(
		http.HandlerFunc(msgServices.SendNewMessage)))

	mux.Handle("/providers", middleware.LogCurrentRequest(
		http.HandlerFunc(msgServices.GetProviders)))

	log.Printf("[INFO] listening on :%s\n", config.AppPort)
	err = http.ListenAndServe(fmt.Sprintf(":%v", config.AppPort), mux)
	if err != nil {
		log.Fatalf("Could not start listening on the given port, %v", err)
	}
}
