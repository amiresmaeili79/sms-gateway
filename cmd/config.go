package cmd

import (
	"github.com/amir79esmaeili/sms-gateway/internal/cfg"
	"github.com/spf13/cobra"
)

func loadConfig(cmd *cobra.Command) cfg.Config {
	configPath, _ := cmd.Flags().GetString("cfg")
	return cfg.ParseConfig(configPath)
}
