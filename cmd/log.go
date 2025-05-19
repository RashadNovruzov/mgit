package cmd

import (
	"github.com/spf13/cobra"
	"mgit/log"
)

func NewLogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "log COMMIT",
		Short: "log command",
		Args:  cobra.RangeArgs(0, 1),
		Run: func(cmd *cobra.Command, args []string) {
			commitOid := ""
			if len(args) != 0 {
				commitOid = getOid(args[0])
			}

			log.PrintLog(commitOid)
		},
	}
	return cmd
}
