package main

import (
	"fmt"
	"log"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
)

const ifaceName = "wlp3s0"
const port = 5173

func main() {

	//

	ifaceIndex, err := getIfaceIdex(ifaceName)
	if err != nil {
		log.Fatalf("Failed to get interface index: %v", err)
	}

	fmt.Println("Interface index:", ifaceIndex, "\nInterface name:", ifaceName)

	// Allow the current process to lock memory for eBPF maps.
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatalf("Failed to remove memlock limit: %v", err)
	}

	// Load the eBPF program into the kernel.
	prog, err := ebpf.NewProgram(getPacketFilterProgram())
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

	fmt.Printf("eBPF program attached to interface %s, dropping TCP packets on port %d\n", ifaceName, port)

	req()

	// Keep the program running
	select {}
}
