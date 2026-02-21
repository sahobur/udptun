//go:build linux

package tunnel

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/vishvananda/netlink"
)

func SetupInterface(name string, ipNet string) error {
	link, err := netlink.LinkByName(name)
	if err != nil {
		return fmt.Errorf("setupInterface: %v", err)
	}

	addr, err := netlink.ParseAddr(ipNet)
	if err != nil {
		return fmt.Errorf("setupInterface: %v", err)
	}

	if err := netlink.AddrAdd(link, addr); err != nil {
		return fmt.Errorf("setupInterface: %v", err)
	}

	if err := netlink.LinkSetUp(link); err != nil {
		return fmt.Errorf("setupInterface: %v", err)
	}

	return nil
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
