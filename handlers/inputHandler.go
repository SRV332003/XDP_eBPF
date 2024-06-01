package handlers

import (
	"fmt"
	"os"
	"strconv"
)

func HandleInput() (string, int) {
	//take user input
	var inputport int
	fmt.Print("Enter the port number to block (press enter to pickup from .env): ")
	fmt.Scanln(&inputport)
	if inputport < 0 || inputport > 65535 {
		fmt.Println("Invalid port number")
		return "", 0
	}
	if inputport == 0 {
		inputport = getPortFromEnv()
	}

	//take user input
	var ifaceName string
	fmt.Print("Enter the interface name (press enter to pickup from .env): ")
	fmt.Scanln(&ifaceName)
	if ifaceName == "" {
		ifaceName = getIfaceFromEnv()
	}

	return ifaceName, inputport

}

func getPortFromEnv() int {
	// Get the port number from the environment variable.
	port := os.Getenv("PORT")
	if port == "" {
		return 5173
	}
	p, err := strconv.Atoi(port)
	if err != nil {
		return 5173
	}
	return p
}

func getIfaceFromEnv() string {
	// Get the interface name from the environment variable.
	iface := os.Getenv("IFACE")
	if iface == "" {
		return "wlp3s0"
	}
	return iface
}
