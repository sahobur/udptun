package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/songgao/water"
	"github.com/vishvananda/netlink"
)

const (
// remoteIP   = "109.207.171.14:8000"
)




func main1() {
	if len(os.Args) < 3 {
		panic("Usage: ./stun <local_port> <remote_ip>")
	}

	localPort, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	remoteAddr := os.Args[2]
	localTunAddr := os.Args[3]
	remoteIP := strings.Split(remoteAddr, ":")[0]

	conn, err := serverUDP(localPort)
	if err != nil {
		panic(err)
	}
	c := 0
	ifce, err := createTunInterface(localTunAddr)
	if err != nil {
		panic(err)
	}

	tunnel := NewUDPTunnel(remoteIP, localPort, conn, localTunAddr, ifce)
	setupRoute(remoteIP)

	go func() {
		if err := receiveUDPPacket(ifce, conn, remoteAddr, "ddddddddddddddddddddd"); err != nil {
			panic(err)
		}
	}()

	for {
		err = sendUDPPacket(ifce, conn, remoteAddr, "ddddddddddddddddddddd")
		if err != nil {
			slog.Error("error send msg ", err)
		}
		c++
	}

}



func sendUDPPacket(ifce *water.Interface, conn *net.UDPConn, remoteAddr string, msg string) error {
	remoteUDPAddr, err := net.ResolveUDPAddr("udp", remoteAddr)
	if err != nil {
		return fmt.Errorf("неверный адрес: %w", err)
	}

	packet := make([]byte, 2000)
	for {
		_, err = ifce.Read(packet)
		if err != nil {
			return fmt.Errorf("ошибка чтения пакета c интерфейса: %w", err)
		}
		_, err = conn.WriteTo(packet, remoteUDPAddr)
		if err != nil {
			return fmt.Errorf("ошибка отправки пакета: %w", err)
		}
	}
	return nil
}

func receiveUDPPacket(ifce *water.Interface, conn *net.UDPConn, remoteAddr string, msg string) error {
	packet := make([]byte, 2000)

	for {
		n, _, err := conn.ReadFromUDP(packet)
		if err != nil {
			fmt.Printf("Ошибка при приёме: %v", err)
			continue
		}
		n, err = ifce.Write(packet[:n])
		if err != nil {
			fmt.Printf("Ошибка при отправке: %v", err)
			continue
		}
	}

	return nil
}


