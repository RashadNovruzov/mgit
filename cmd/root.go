package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "mgit",
		Short: "Mini git short",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Usage(); err != nil {
				log.Fatalln(err)
			}
		},
	}
	return rootCmd
}
