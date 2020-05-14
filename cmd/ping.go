package cmd

import (
	"github.com/mritd/mmh/core"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:     "ping SERVER_NAME",
	Aliases: []string{"mping"},
	Short:   "ping server",
	Long:    "ping server.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			core.Ping(args[0])
		} else {
			_ = cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(pingCmd)
}
