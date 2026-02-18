package server

import (
	"fmt"
	"log"
	"log/slog"
	"net"
	"os/exec"

	"github.com/songgao/water"
	"github.com/vishvananda/netlink"
)

type clientID string

type Server struct {
	Port    int
	Listen  net.IP
	TunIf   *water.Interface
	TunAddr string
	Conn    *net.UDPConn
}

func NewServer(ip string, port int, locTunAddr string) *Server {
	slog.Info("Creating server", "ip", ip, "port", port, "tunAddr", locTunAddr)
	return &Server{
		Listen:  net.ParseIP(ip),
		Port:    port,
		TunAddr: locTunAddr,
	}
}

func (s *Server) StartServer() error {
	var err error
	s.Conn, err = s.serverUDP()
	if err != nil {
		return err
	}
	s.TunIf, err = s.createTunInterface()
	if err != nil {
		return err
	}

	// remoteIP := strings.Split(remoteAddr, ":")[0]
	// setupRoute(remoteIP)
	return nil
}

func (s *Server) serverUDP() (*net.UDPConn, error) {
	slog.Info("Starting UDP server", "port", fmt.Sprintf("%s:%d", s.Listen, s.Port))
	addr := &net.UDPAddr{IP: s.Listen, Port: s.Port}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("не удалось запустить UDP-сервер: %w", err)
	}
	return conn, nil
}

func (s *Server) createTunInterface() (*water.Interface, error) {
	config := water.Config{
		DeviceType: water.TUN,
	}

	// Создание интерфейса
	ifce, err := water.New(config)
	if err != nil {
		panic(err)
	}
	s.TunIf = ifce
	slog.Info("TUN interface created", "name", ifce.Name())
	setupInterface(ifce.Name(), s.TunAddr)

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

func (s *Server) authUserConn() {

}
