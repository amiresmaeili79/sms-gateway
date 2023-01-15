package cmd

import (
	"fmt"
	"github.com/amir79esmaeili/sms-gateway/internal/middleware"
	"github.com/spf13/cobra"
	"log"
	"net/http"
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

	mux := http.NewServeMux()

	mux.Handle("/ping", middleware.LogCurrentRequest(
		http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("Hello World!"))
		})))

	err := http.ListenAndServe(fmt.Sprintf(":%v", config.AppPort), mux)
	if err != nil {
		log.Fatalf("Could not start listening on the given port, %v", err)
	}
}
