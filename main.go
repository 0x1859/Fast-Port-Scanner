package main

import (
	"fmt"
	"net"
	"os"

	"github.com/MasterDimmy/go-cls"
	"github.com/admin100/util/console"
)

var (
	targetIP  string
	startPort int
	endPort   int
)

func main() {
	cls.CLS()
	console.SetConsoleTitle("Port Scanner")

	fmt.Print("[>] Enter the target IP: ")
	fmt.Scanln(&targetIP)

	fmt.Print("[>] Enter the start port: ")
	fmt.Scanln(&startPort)

	fmt.Print("[>] Enter the end port: ")
	fmt.Scanln(&endPort)

	file, err := os.Create("open.txt")
	if err != nil {
		fmt.Println("[!] Error creating file:", err)
		return
	}
	defer file.Close()

	cls.CLS()
	fmt.Printf("Scanning %s from port %d to %d", targetIP, startPort, endPort)
	fmt.Println("")

	for port := startPort; port <= endPort; port++ {
		address := fmt.Sprintf("%s:%d", targetIP, port)
		conn, err := net.Dial("tcp", address)
		if err == nil {
			file.WriteString(fmt.Sprintf("%s:%d\n", targetIP, port))
			fmt.Printf("Port %d is Open\n", port)
			conn.Close()
		}
	}
	cls.CLS()
	fmt.Println("[+] Open ports have been written to open.txt.")
}
