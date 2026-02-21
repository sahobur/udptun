//go:build darwin

package tunnel

import (
	"net"
	"os/exec"
)

import (
	"log"
)

func SetupInterface(name, localIP string) error {
	// macOS требует: ifconfig <name> <local> <remote> up
	ip, _, err := net.ParseCIDR(localIP)
	if err != nil {
		return err
	}
	ips := ip.String()
	cmd := exec.Command("ifconfig", name, localIP, ips, "up")
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Printf("macOS: %s up", name)
	return nil
}
