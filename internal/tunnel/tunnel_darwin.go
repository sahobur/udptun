//go:build darwin

package tunnel

import "os/exec"

import (
	"log"
)

func SetupInterface(name, localIP string) error {
	// macOS требует: ifconfig <name> <local> <remote> up
	cmd := exec.Command("ifconfig", name, localIP, "10.30.1.1", "up")
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Printf("macOS: %s up", name)
	return nil
}
