package tunnel

import (
	"log/slog"
	"net"

	"github.com/songgao/water"
)

type UDPTunnel struct {
	localPort     int
	conn          *net.UDPConn
	localTunnelIP string
	ifce          *water.Interface
	RemoteAddress *net.UDPAddr
}

func NewUDPTunnel(localPort int, conn *net.UDPConn, remoteAddress *net.UDPAddr,
	localTunnelIP string, ifce *water.Interface) UDPTunnel {
	//todo validate
	return UDPTunnel{
		localPort:     localPort,
		conn:          conn,
		localTunnelIP: localTunnelIP,
		ifce:          ifce,
		RemoteAddress: remoteAddress,
	}
}

func (t UDPTunnel) Send() {
	packet := make([]byte, 2000)
	for {
		_, err := t.ifce.Read(packet)
		if err != nil {
			slog.Error("read from tunnel",
				slog.String("interface", t.ifce.Name()),
				slog.String("error", err.Error()))
		}
		_, err = t.conn.WriteTo(packet, t.RemoteAddress)
		if err != nil {
			slog.Error("write to udp connection",
				slog.String("connection", t.conn.RemoteAddr().String()),
				slog.String("error", err.Error()))
		}
	}
}

func (t UDPTunnel) Recv() {
	packet := make([]byte, 2000)
	for {
		n, _, err := t.conn.ReadFromUDP(packet)
		if err != nil {
			slog.Error("read from",
				slog.String("udp connection", t.conn.RemoteAddr().String()),
				slog.String("error", err.Error()))
			return
		}

		n, err = t.ifce.Write(packet[:n])
		if err != nil {
			slog.Error("write to interface",
				slog.String("interface", t.ifce.Name()),
				slog.String("error", err.Error()))
			continue
		}
	}
}
