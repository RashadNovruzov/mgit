package cmd

import (
	"github.com/spf13/cobra"
	initcmd "mgit/init"
)

func NewInitCommand() *cobra.Command {
	initCommand := &cobra.Command{
		Use:   "init",
		Short: "init command",
		Long:  `initialize git repository`,
		Run: func(cmd *cobra.Command, args []string) {
			initcmd.InitializeRepo()
		},
	}
	return initCommand
}
