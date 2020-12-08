package cmd

import (
	"github.com/mrydengren/elvis/pkg/cmd/setlist"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"log"
)

func NewCmdSetlist() *cobra.Command {
	cmdSetlist := cobra.Command{
		Use:   "setlist <value>",
		Short: "Generate a playlist from a setlist.fm setlist",
		Example: "" +
			"  elvis setlist https://www.setlist.fm/setlist/opeth/2017/finlandia-talo-helsinki-finland-53e3dba9.html\n" +
			"  elvis setlist 53e3dba9",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return pflag.ErrHelp
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 1 {
				log.Printf("Ignoring extra arguments: %+v\n", args[1:])
			}
			return setlist.Setlist(args[0])
		},
	}

	return &cmdSetlist
}

func init() {
	rootCmd.AddCommand(NewCmdSetlist())
}
