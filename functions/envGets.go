package functions

import (
	"os"
	"strconv"
)

func EnvPort() int {
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

func EnvIFace() string {
	// Get the interface name from the environment variable.
	iface := os.Getenv("IFACE")
	if iface == "" {
		return "wlp3s0"
	}
	return iface
}

func EnvProcess() string {
	// Get the process name from the environment variable.
	process := os.Getenv("PROCESS")
	if process == "" {
		return "brave-browser"
	}
	return process
}
