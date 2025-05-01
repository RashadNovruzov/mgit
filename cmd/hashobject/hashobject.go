package hashobject

import (
	"github.com/spf13/cobra"
	"mgit/hashobject"
)

func NewHashObjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hashobject FILENAME",
		Args:  cobra.ExactArgs(1),
		Short: "hashes and saves object to .mgit folder",
		Long:  `hashes and saves object to .mgit folder`,
		Run: func(cmd *cobra.Command, args []string) {
			hashobject.HashObject(args[0])
		},
	}
	return cmd
}
