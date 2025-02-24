package scanner

import (
	"net"
	"strconv"
	"time"
)

type PortInfo struct {
	Port   int
	Status string
}

func CheckPort(protocol, host string, port int) PortInfo {
	portData := PortInfo{Port: port}
	address := net.JoinHostPort(host, strconv.Itoa(port))

	conn, err := net.DialTimeout(protocol, address, 500*time.Millisecond)
	if err != nil {
		portData.Status = "Closed"
		return portData
	}
	defer conn.Close()

	portData.Status = "Open"
	return portData
}

func ScanPorts(host string, startPort, endPort int) []PortInfo {
	var results []PortInfo

	for port := startPort; port <= endPort; port++ {
		results = append(results, CheckPort("tcp", host, port))
	}

	return results
}
