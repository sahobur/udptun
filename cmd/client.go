/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
)

var remoteAddr string

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "start client",
	Long:  `client`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("client start")
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.Flags().StringVarP(&remoteAddr, "remote-addr", "r", "", "Remote address")
	clientCmd.Flags().IntVarP(&localPort, "local-port", "p", 8888, "Local port")
	clientCmd.Flags().StringVarP(&tunnelAddr, "tunnel-addr", "t", "10.30.1.2/24", "Tunnel address")
	clientCmd.MarkFlagRequired("remote-addr")
	clientCmd.MarkFlagRequired("local-port")
	clientCmd.MarkFlagRequired("tunnel-addr")
}
