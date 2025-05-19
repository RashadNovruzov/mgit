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
			oid := getOid(args[0])
			println(catfile.CatFile(oid, ""))
		},
	}
	return cmd
}
