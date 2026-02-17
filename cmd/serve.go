/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/saho/udptun/internal/server"
)

var localPort string
var localTunAddr string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run tunnel server",
	Long:  `ip tunnel server`,
	Run: func(cmd *cobra.Command, args []string) {
		server.startServer(localPort, localTunAddr)
		
	},
}


func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringVarP(&localPort, "local-port", "p", "", "Local port")
	serveCmd.Flags().StringVarP(&localTunAddr, "local-tun-addr", "l", "", "Local tunnel address")
	serveCmd.MarkFlagRequired("local-port")
	serveCmd.MarkFlagRequired("local-tun-addr")

}
