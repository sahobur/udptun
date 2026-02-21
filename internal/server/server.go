package server

import (
	"fmt"

	"log/slog"
	"net"

	"strings"

	"github.com/sahobur/udptun/internal/model"
	"github.com/sahobur/udptun/internal/tunnel"

	//"github.com/sahobur/udptun/cmd"
	"github.com/songgao/water"
)

type clientID string

type Server struct {
	Port    int
	Listen  net.IP
	TunIf   *water.Interface
	TunAddr string
	Conn    *net.UDPConn
	Clients []model.Client
}

func NewServer(ip string, port int, locTunAddr string, cl []model.Client) *Server {
	slog.Info("Creating server", "ip", ip, "port", port, "tunAddr", locTunAddr)
	return &Server{
		Listen:  net.ParseIP(ip),
		Port:    port,
		TunAddr: locTunAddr,
		Clients: cl,
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

	if err = tunnel.SetupInterface(ifce.Name(), s.TunAddr); err != nil {
		return nil, err
	}

	return ifce, nil
}

func (s *Server) AuthUserConn(packet []byte, n int) model.Client {
	key := strings.Trim(string(packet[:n]), "\n")
	return s.clientKeyExist(key)
}

func (s *Server) clientKeyExist(key string) model.Client {
	for _, v := range s.Clients {
		if v.Key == key {
			return v
		}
	}
	return model.Client{}
}
