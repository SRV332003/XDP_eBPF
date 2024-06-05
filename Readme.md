# TCP Port Blocker (Problem 1)

This is a Go project that uses eBPF to drop TCP packets on a specific port for a specific network interface. This is the assessment for recruitment process at Accuknox. Current Readme.md describes the solution of 1st problem statement. For other problems, visit [`AccuknowTest Repository`](https://github.com/SRV332003/AccuknoxTest "All Problems")

## Usage
This program needs to be executed as super user or previledged-user as it needs to lock memory for `ebpf maps`.
To run this project, install dependencies and execute the [`main.go`](https://github.com/SRV332003/XDP_eBPF/blob/main/main.go "main.go") file:

```bash
go mod tidy
sudo go run main.go
```

You will be prompted to enter the port number to block and the interface name. If you press enter without providing any input, the values will be picked up from the environment variables.

## Environment Variables
This project uses the following environment variables:

- `PORT` : The port number to block. If not provided, the default value is `5173`.
- `IFACE` : The name of the network interface. If not provided, the default value is `wlp3s1`.

```env
# .env
PORT=8080     # drop packets on port 8080
IFACE=wlp3s0  # attach program to wlp3s0 network interface
 ```
## Code Structure
- **[`main.go`](https://github.com/SRV332003/XDP_eBPF/blob/main/main.go "main.go")** : This is the entry point of the application. It handles user input, loads the eBPF program into the kernel, and attaches the eBPF program to a network interface.
- **[`handlers`](https://github.com/SRV332003/XDP_eBPF/blob/main/handlers/inputHandler.go "handlers module") :** This module container higher level abstraction allowing operations such as handling `input` and providing `ebpf` program from the information. This `go module` contains the following files:
    - **[`handlers/inputHandler.go`](https://github.com/SRV332003/XDP_eBPF/blob/main/handlers/inputHandler.go "handlers/inputHandler.go") :** This file contains the HandleInput function which handles `user input` or `.env` for the `port` number and the `interface` name. This uses the `functions module` to process info.
    - **[`handlers/xdpHandler.go`](https://github.com/SRV332003/XDP_eBPF/blob/main/handlers/xdpHandler.go "handlers/inputHandler.go") :** This file contains the GetXDPProgram function which returns an object `ebpf.ProgramSpec` that contains instructions as the part of `ebpf program`. This file is made to separate the assembly code written using `github.com/cilium/ebpf/asm` module from the rest of linking code.
- **[`functions`](https://github.com/SRV332003/XDP_eBPF/blob/main/functions "functions module") :**
    - **[`functions/envGets.go`](https://github.com/SRV332003/XDP_eBPF/blob/main/functions/envGets.go "functions/envGets.go") :** This file contains functions to get environment variables.
    - **[`functions/osGets.go`](https://github.com/SRV332003/XDP_eBPF/blob/main/functions/osGets.go "functions/envGets.go") :** This file contains functions to get OS-related information.
## Dependencies
This project uses the following dependencies:

- `github.com/cilium/ebpf`: A package to work with eBPF programs in Go.
- `github.com/joho/godotenv`: A package to load environment variables from a .env file.