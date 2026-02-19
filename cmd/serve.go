/*
Copyright © 2026 Alexandr Lykshitov 278999@mail.ru
*/
package cmd

import (
	"log/slog"
	"os"
	"strings"

	"github.com/sahobur/udptun/internal/server"
	"github.com/sahobur/udptun/internal/tunnel"
	"github.com/spf13/cobra"
)

var localPort int
var listenAddr string
var tunnelAddr string

const clientKey = "pingpingping"

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run tunnel server",
	Long:  `ip tunnel server`,
	Run: func(cmd *cobra.Command, args []string) {
		server := server.NewServer(listenAddr, localPort, tunnelAddr)
		if err := server.StartServer(); err != nil {
			slog.Error("server sart", slog.String("error", err.Error()))
			os.Exit(1)
		}
		// todo receive udp packets for handshake and create tunnel
		packet := make([]byte, 1500)
		for {
			p, remoteAddr, err := server.Conn.ReadFromUDP(packet)
			if err != nil {
				slog.Error("read from", slog.String("error", err.Error()))
				continue
			}
			slog.Info("received packet", slog.Int("length", p), slog.String("remoteAddr", remoteAddr.String()))
			if strings.Trim(string(packet[:p]), "\n") == clientKey {
				slog.Info("client connected", slog.String("remoteAddr", remoteAddr.String()))
				t := tunnel.NewUDPTunnel(localPort, server.Conn, remoteAddr, tunnelAddr, server.TunIf)
				go func() {
					t.Send()
				}()
				t.Recv()
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&localPort, "local-port", "p", 9999, "Local port")
	serveCmd.Flags().StringVarP(&listenAddr, "listen-addr", "l", "0.0.0.0", "Local tunnel address")
	serveCmd.Flags().StringVarP(&tunnelAddr, "tunnel-addr", "t", "10.30.1.1/24", "Tunnel address")
	serveCmd.MarkFlagRequired("local-port")
	serveCmd.MarkFlagRequired("listen-addr")
	serveCmd.MarkFlagRequired("tunnel-addr")
}
