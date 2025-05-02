package writetree

import (
	"github.com/spf13/cobra"
	"mgit/writetree"
)

func NewWriteTreeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "write-tree",
		Short: "hashes and saves whole folders recursively to .mgit folder",
		Run: func(cmd *cobra.Command, args []string) {
			writetree.WriteTree("")
		},
	}
	return cmd
}
