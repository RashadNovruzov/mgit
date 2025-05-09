package cmd

import (
	"github.com/spf13/cobra"
	"mgit/commit"
)

var message string

func NewCommitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "commit",
		Short: "commit command",
		Run: func(cmd *cobra.Command, args []string) {
			println(commit.Commit(message))
		},
	}
	cmd.Flags().StringVarP(&message, "message", "m", "", "commit message")
	err := cmd.MarkFlagRequired("message")
	if err != nil {
		cmd.HelpFunc()
	}
	return cmd
}
