package tunnel

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/songgao/water"
)

type UDPTunnel struct {
	remoteIP      string
	localPort     int
	conn          *net.UDPConn
	localTunnelIP string
	ifce          *water.Interface
	RemoteAddress *net.UDPAddr
}

func NewUDPTunnel(remoteIP string, localPort int, conn *net.UDPConn, remoteAddress *net.UDPAddr,
	localTunnelIP string, ifce *water.Interface) UDPTunnel {
	//todo validate
	return UDPTunnel{
		remoteIP:      remoteIP,
		localPort:     localPort,
		conn:          conn,
		localTunnelIP: localTunnelIP,
		ifce:          ifce,
		RemoteAddress: remoteAddress,
	}
}

func (t *UDPTunnel) Send() {
	packet := make([]byte, 2000)
	for {
		_, err := t.ifce.Read(packet)
		if err != nil {
			slog.Error("read from", slog.String("error", err.Error()))
		}
		fmt.Println("read from", packet)
		//_, err = t.conn.WriteTo(data, t.RemoteAddress)
	}
}

func (t *UDPTunnel) Recv() {
	data := make([]byte, 2000)
	for {
		n, remoteAddr, err := t.conn.ReadFromUDP(data)
		if err != nil {
			slog.Error("read from", slog.String("error", err.Error()))
			return
		}

		fmt.Printf("recv packet from %s:\n%v\n", remoteAddr.String(), data[:n]) // which client?
	}
}
