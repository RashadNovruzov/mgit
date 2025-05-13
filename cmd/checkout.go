package cmd

import (
	"github.com/spf13/cobra"
	"mgit/checkout"
)

func NewCheckoutCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "checkout OID",
		Args:  cobra.ExactArgs(1),
		Short: "Checkout to commit by oid",
		Run: func(cmd *cobra.Command, args []string) {
			oid := args[0]
			checkout.Checkout(oid)
		},
	}
	return cmd
}
