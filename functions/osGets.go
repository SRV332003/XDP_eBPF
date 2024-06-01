package functions

import (
	"bytes"
	"errors"
	"net"
	"os/exec"
	"strconv"
)

func GetIfaceIndex(ifaceName string) (int, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return 0, err
	}

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

func GetPIDByName(name string) (int, error) {
	cmd := exec.Command("pgrep", name)
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	pidStr := string(bytes.TrimSpace(output))
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		return 0, err
	}

	return pid, nil
}
