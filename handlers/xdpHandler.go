package handlers

import (
	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/asm"
)

func GetXDPProgram(port int) *ebpf.ProgramSpec {
	return &ebpf.ProgramSpec{
		Type:    ebpf.XDP,
		License: "GPL",
		Instructions: asm.Instructions{
			// Load pointers to the start and end of the packet
			asm.LoadMem(asm.R6, asm.R1, 0, asm.Word), // Load data pointer into R6
			asm.LoadMem(asm.R7, asm.R1, 4, asm.Word), // Load data_end pointer into R7

			asm.Mov.Reg(asm.R2, asm.R6),
			asm.Add.Imm(asm.R2, 23), // Offset for IP protocol field (14+9=23)

			// Check if the IP protocol field is within packet bounds
			asm.JGE.Reg(asm.R2, asm.R7, "exit"), // if R2 > R7 (end of packet), jump to exit

			// find ip header length
			asm.LoadMem(asm.R8, asm.R6, 14, asm.Byte),
			asm.And.Imm(asm.R8, 0x0F),   // Mask out the lower 4 bits
			asm.Mov.Imm(asm.R9, 2),      // Multiply by 4 (IP header length is in 4-byte words)
			asm.LSh.Reg(asm.R8, asm.R9), // Shift R8 left by R9 bits

			// Load IP protocol field
			asm.LoadMem(asm.R2, asm.R6, 23, asm.Byte),
			asm.JNE.Imm(asm.R2, 0x06, "exit"), // Jump to exit if not TCP (0x06)

			// Calculate the offset of IP protocol field
			asm.Add.Reg(asm.R8, asm.R6), // Offset for IP protocol field
			asm.Add.Imm(asm.R8, 14),
			asm.Mov.Reg(asm.R9, asm.R8), // Offset for IP protocol field (14+9=23)
			asm.Add.Imm(asm.R9, 4),      // Offset for TCP dest port field (23+13=36)

			asm.JGE.Reg(asm.R9, asm.R7, "exit"), // if R8 > R7 (end of packet), jump to exit

			// Load TCP destination port field
			asm.LoadMem(asm.R2, asm.R8, 3, asm.Half), // Load TCP dest port field into R2 (offset 40)
			asm.JEq.Imm(asm.R2, int32(port), "exit"), // Jump to exit if not port 4040

			asm.Mov.Imm(asm.R0, 1), // Set return code to XDP_DROP (1)
			asm.Return(),           // Return from program

			// Exit label
			asm.Mov.Imm(asm.R0, 2).WithSymbol("exit"), // Set return code to XDP_PASS (2)
			asm.Return(), // Return from program
		},
	}
}
