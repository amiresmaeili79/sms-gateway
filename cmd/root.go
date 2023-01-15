package cmd

import "github.com/spf13/cobra"

func New() *cobra.Command {
	root := &cobra.Command{
		Use:   "root <subcommand>",
		Short: "root Daemon",
		Run:   nil,
	}
	addServeCmd(root)
	return root
}
