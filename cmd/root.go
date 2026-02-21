/*
Copyright © 2026 Alexandr Lykshitov 278999@mail.ru
*/
package cmd

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/spf13/cobra"
)

var cfg *Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "udptun",
	Short: "udp tunneling app",
}

func Execute() {
	w := os.Stderr
	slog.SetDefault(slog.New(
		tint.NewHandler(w, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.TimeOnly,
			AddSource:  true,
		}),
	))

	err := cfg.Load()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
