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

func serverUDP(localPort int) (*net.UDPConn, error) {

	addr := &net.UDPAddr{IP: net.IPv4(0, 0, 0, 0), Port: localPort}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("не удалось запустить UDP-сервер: %w", err)
	}
	return conn, nil
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

func createTunInterface(ip string) (*water.Interface, error) {
	config := water.Config{
		DeviceType: water.TUN,
	}

	// Создание интерфейса
	ifce, err := water.New(config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Интерфейс создан: %s\n", ifce.Name())

	setupInterface(ifce.Name(), ip+"/30")

	return ifce, nil
}

func setupInterface(name string, ip string) {
	exec.Command("ip", "addr", "add", ip, "dev", name).Run()
	exec.Command("ip", "link", "set", "dev", name, "up").Run()
}
func setupRoute(remoteIP string) {
	if err := exec.Command("ip", "route", "add", remoteIP, "via", "tun0").Run(); err != nil {
		log.Fatal(err)
	}

	if getDefaultRoute() != "" {
		if err := exec.Command("ip", "route", "add", "default", "via", getDefaultRoute()).Run(); err != nil {
			log.Fatal(err)
		}
	}
}

func getDefaultRoute() string {
	routes, err := netlink.RouteList(nil, netlink.FAMILY_V4)
	if err != nil {
		log.Fatal(err)
	}
	for _, route := range routes {
		if route.Gw != nil {
			return route.Gw.String()
		}
	}
	return ""
}
