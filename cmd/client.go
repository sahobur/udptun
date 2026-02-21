/*
Copyright © 2026 Alexandr Lykshitov 278999@mail.ru
*/
package cmd

import (
	"log/slog"
	"net"
	"os"

	"github.com/sahobur/udptun/internal/server"
	"github.com/sahobur/udptun/internal/tunnel"
	"github.com/spf13/cobra"
)

var remoteAddr string
var remotePort int
var clientKey string

// clientCmd represents the client command
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "start client",
	Long:  `client`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("client start")
		server := server.NewServer("0.0.0.0", localPort, tunnelAddr, cfg.Clients)
		if err := server.StartServer(); err != nil {
			slog.Error("server sart", slog.String("error", err.Error()))
			os.Exit(1)
		}

		n, err := server.Conn.WriteTo([]byte(clientKey), &net.UDPAddr{IP: net.ParseIP(remoteAddr), Port: remotePort})
		if err != nil {
			slog.Error("write to", slog.String("error", err.Error()))
			os.Exit(1)
		}
		slog.Info("write to", slog.Int("length", n))

		t := tunnel.NewUDPTunnel(localPort, server.Conn,
			&net.UDPAddr{IP: net.ParseIP(remoteAddr), Port: remotePort}, tunnelAddr, server.TunIf)

		go func() {
			t.Send()
		}()
		t.Recv()
	},
}

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.Flags().StringVarP(&remoteAddr, "remote-addr", "r", "109.207.171.14", "Remote address")
	clientCmd.Flags().IntVarP(&localPort, "local-port", "s", 8888, "Local port")
	clientCmd.Flags().IntVarP(&remotePort, "remote-port", "p", 9999, "remote port")
	clientCmd.Flags().StringVarP(&tunnelAddr, "tunnel-addr", "t", "10.30.1.2/24", "Tunnel address")
	clientCmd.Flags().StringVarP(&clientKey, "client-key", "k", "key", "preauth key")
	clientCmd.MarkFlagRequired("remote-addr")
	clientCmd.MarkFlagRequired("tunnel-addr")
}
