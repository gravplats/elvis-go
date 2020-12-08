package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/mrydengren/elvis/pkg/config"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

func NewCmdInit() *cobra.Command {
	cmdInit := cobra.Command{
		Use:     "init",
		Short:   "Generate a config file in the home directory (will not overwrite)",
		Example: "  elvis init",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) > 0 {
				fmt.Printf("Ignoring extra arguments: %+v\n", args)
			}

			home, err := homedir.Dir()
			if err != nil {
				return err
			}

			filename := filepath.Join(home, ".elvis.json")

			if _, err := os.Stat(filename); err == nil {
				return fmt.Errorf("Config file already exists and won't be overwritten.")
			}

			config := config.Config{}
			buf, err := json.MarshalIndent(config, "", "  ")

			err = ioutil.WriteFile(filename, buf, 0644)
			if err != nil {
				return err
			}

			fmt.Printf("Config file has been written to %v.\n", filename)

			return nil
		},
	}

	return &cmdInit
}

func init() {
	rootCmd.AddCommand(NewCmdInit())
}
