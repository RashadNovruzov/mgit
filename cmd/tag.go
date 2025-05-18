package cmd

import (
	"github.com/spf13/cobra"
	"mgit/commit"
	"mgit/tag"
)

func NewTagCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tag NAME OID",
		Args:  cobra.RangeArgs(1, 2),
		Short: "tag command",
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			oid := ""
			if name == "" {
				panic("tag name is required")
			}
			if len(args) == 2 {
				oid = args[1]
			}
			if oid == "" {
				oid = commit.GetHead()
			}
			tag.CreateTag(name, oid)
		},
	}
	return cmd
}
