package main

import (
	"fmt"
	"log"

	//import env package

	"github.com/SRV332003/XDP_eBPF/handlers"
	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
	"github.com/joho/godotenv"
)

func main() {

	ifaceIndex, port, err := handlers.HandleInput()
	if err != nil {
		log.Fatalf("Failed to get interface index: %v", err)
	}

	// Allow the current process to lock memory for eBPF maps.
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatalf("Failed to remove memlock limit: %v", err)
	}

	preCompiled := handlers.GetXDPProgram(port)

	// Load the eBPF program into the kernel.
	prog, err := ebpf.NewProgram(preCompiled)
	if err != nil {
		log.Fatalf("Failed to load eBPF program: %v", err)
	}
	defer prog.Close()

	// Attach the eBPF program to a network interface.
	l, err := link.AttachXDP(link.XDPOptions{
		Program:   prog,
		Interface: ifaceIndex,
		Flags:     link.XDPGenericMode, // Use XDPGenericMode if your NIC doesn't support native XDP
	})
	if err != nil {
		log.Fatalf("Failed to attach XDP program: %v", err)
	}
	defer l.Close()

	fmt.Printf("Started dropping TCP packets on port %d\n", port)

	select {}
}

func init() {
	//load .env file
	godotenv.Load(".env")
}
