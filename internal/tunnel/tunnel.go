package main

import (
	"net"

	"github.com/songgao/water"
)

type UDPTunnel struct {
	remoteIP      string
	localPort     int
	conn          *net.UDPConn
	localTunnelIP string
	ifce          *water.Interface
}

func NewUDPTunnel(remoteIP string, localPort int, conn *net.UDPConn, localTunnelIP string, ifce *water.Interface) UDPTunnel {
	//todo validate
	return UDPTunnel{
		remoteIP:      remoteIP,
		localPort:     localPort,
		conn:          conn,
		localTunnelIP: localTunnelIP,
		ifce:          ifce,
	}
}

func (t *UDPTunnel) Send(data []byte) error {
	_, err := t.conn.Write(data)
	return err
}

func (t *UDPTunnel) Recv() ([]byte, error) {
	data := make([]byte, 2000)
	_, err := t.conn.Read(data)
	return data, err
}
