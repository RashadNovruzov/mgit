package cmd

import (
	"github.com/spf13/cobra"
	"mgit/catfile"
)

func NewCatFileCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cat-file FILENAME",
		Args:  cobra.ExactArgs(1),
		Short: "cat-file opens oboject(file) with given file hash sum",
		Run: func(cmd *cobra.Command, args []string) {
			catfile.CatFile(args[0], "")
		},
	}
	return cmd
}
