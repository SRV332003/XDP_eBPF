package handlers

import (
	"fmt"

	"github.com/SRV332003/XDP_eBPF/functions"
)

func HandleInput() (int, int, error) {
	//take user input
	var inputport int
	fmt.Print("Enter the port number to block (press enter to pickup from .env): ")
	fmt.Scanln(&inputport)
	if inputport < 0 || inputport > 65535 {
		fmt.Println("Invalid port number")
		return 0, 0, fmt.Errorf("Invalid port number")
	}
	if inputport == 0 {
		inputport = functions.EnvPort()
	}

	//take user input
	var ifaceName string
	fmt.Print("Enter the interface name (press enter to pickup from .env): ")
	fmt.Scanln(&ifaceName)
	if ifaceName == "" {
		ifaceName = functions.EnvIFace()
	}

	ifaceIndex, err := functions.GetIfaceIndex(ifaceName)
	if err != nil {
		return 0, 0, err
	}

	fmt.Println("-------------------------")
	fmt.Println("Interface index:", ifaceIndex, "\nInterface name:", ifaceName)

	return ifaceIndex, inputport, nil

}
