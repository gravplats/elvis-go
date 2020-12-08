package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/mrydengren/elvis/pkg/api/lastfm"
	"github.com/mrydengren/elvis/pkg/api/setlistfm"
	"github.com/mrydengren/elvis/pkg/config"
	"github.com/mrydengren/elvis/pkg/debug"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strconv"
	"time"
)

type RootOptions struct {
	Debug bool
}

var rootOpts = RootOptions{
	Debug: false,
}

var rootCmd = cobra.Command{
	Use:   "elvis",
	Short: "A Spotify playlist generator",
	Long:  "Elvis is a CLI for generating Spotify playlists from various web APIs.",
	// Only print error once if there is an error in a sub-command.
	SilenceErrors: true,
	// Don't show usage if there is an error in a sub-command. Each sub-command will show usage if args are incorrect.
	SilenceUsage: true,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		initDebug()

		// Don't try to read config file if we are trying to create a config file.
		if cmd.Name() == "init" {
			return nil
		}

		err := initConfig()
		if err != nil {
			return err
		}

		return nil
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		debug.DumpInput()
		if dir, ok := debug.GetDebugDir(); ok {
			fmt.Printf("\nWrote debug information to %s.\n", dir)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initConfig() error {
	home, err := homedir.Dir()
	if err != nil {
		return fmt.Errorf("Unable to get home dir: %v\n", err)
	}

	viper.SetConfigName(".elvis")
	viper.SetConfigType("json")

	viper.AddConfigPath(home)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("Unable to read config file: %v\n", err)
	}

	var cfg config.Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return fmt.Errorf("Unable to unmarshal config: %v", err)
	}

	// The lastfm API wrapper reads these, if present.
	os.Setenv(lastfm.ApiKey, cfg.Lastfm.Credentials.Key)
	os.Setenv(lastfm.ApiSecret, cfg.Lastfm.Credentials.Secret)
	// The setlistfm API wrapper reads these, if present.
	os.Setenv(setlistfm.ApiKey, cfg.Setlistfm.Credentials.Key)
	// The Spotify API wrapper reads these, if present.
	os.Setenv("SPOTIFY_ID", cfg.Spotify.Credentials.Id)
	os.Setenv("SPOTIFY_SECRET", cfg.Spotify.Credentials.Secret)
	// YouTube API keys
	os.Setenv("YOUTUBE_KEY", cfg.YouTube.Credentials.Key)

	return nil
}

func initDebug() {
	os.Setenv(debug.DEBUG, strconv.FormatBool(rootOpts.Debug))
	os.Setenv(debug.DEBUG_SESSION_ID, strconv.FormatInt(time.Now().Unix(), 10))
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&rootOpts.Debug, "debug", false, "Output debug information")
	rootCmd.PersistentFlags().MarkHidden("debug")
}
