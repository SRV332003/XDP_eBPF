package main

import (
	"errors"
	"net"
)

// getIfaceIdex returns the index of the network interface with the given name.
// If the interface is not found, an error is returned.
func getIfaceIdex(ifaceName string) (int, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return 0, err
	}

	// Find the index of the interface we want to attach the eBPF program to.
	var ifaceIndex int
	for _, iface := range ifaces {
		if iface.Name == ifaceName {
			ifaceIndex = iface.Index
			break
		}
	}

	if ifaceIndex == 0 {
		return 0, errors.New("Interface not found")
	}

	return ifaceIndex, nil
}
