package cmd

import (
	"github.com/spf13/cobra"
	"mgit/readtree"
)

func NewReadTreeCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "read-tree TREEOID",
		Args:  cobra.ExactArgs(1),
		Short: "read-tree of mgit repo and apply changes",
		Run: func(cmd *cobra.Command, args []string) {
			oid := getOid(args[0])
			readtree.ReadTree(oid)
		},
	}

	return cmd
}
