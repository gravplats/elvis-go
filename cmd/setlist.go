package cmd

import (
	"github.com/mrydengren/elvis/pkg/cmd/setlist"
	"github.com/spf13/cobra"
)

func NewCmdSetlist() *cobra.Command {
	cmdSetlist := cobra.Command{
		Use:   "setlist <value>",
		Short: "Generate a playlist from a setlist.fm setlist",
		Example: "" +
			"  elvis setlist https://www.setlist.fm/setlist/opeth/2017/finlandia-talo-helsinki-finland-53e3dba9.html\n" +
			"  elvis setlist 53e3dba9",
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			setlist.Setlist(args[0])
		},
	}

	return &cmdSetlist
}

func init() {
	rootCmd.AddCommand(NewCmdSetlist())
}
