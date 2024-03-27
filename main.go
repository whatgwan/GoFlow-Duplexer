package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
)

func main() {
	var (
		deviceAddresses string
		listenPort      int
	)
	flag.StringVar(&deviceAddresses, "devices", "", "Comma-separated list of device addresses in the format 'ip1:port1,ip2:port2'")
	flag.IntVar(&listenPort, "port", 2055, "Port to listen for NetFlow packets")
	flag.Parse()

	if deviceAddresses == "" {
		fmt.Println("Please provide device addresses using -devices flag")
		return
	}

	devices := parseDeviceAddresses(deviceAddresses)

	deviceConnections := make([]*net.UDPConn, len(devices))
	for i, address := range devices {
		deviceAddr, err := net.ResolveUDPAddr("udp", address)
		if err != nil {
			fmt.Println("Error resolving address:", err)
			return
		}

		conn, err := net.DialUDP("udp", nil, deviceAddr)
		if err != nil {
			fmt.Println("Error connecting to device:", err)
			return
		}
		defer conn.Close()

		deviceConnections[i] = conn
	}

	serverAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", listenPort))
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	serverConn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer serverConn.Close()

	fmt.Printf("Listening for NetFlow packets on port %d...\n", listenPort)

	// Buffer to store incoming packet data
	buffer := make([]byte, 65535)

	for {
		n, addr, err := serverConn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			continue
		}

		fmt.Printf("Received NetFlow packet from: %s\n", addr.IP)

		for _, conn := range deviceConnections {
			_, err := conn.Write(buffer[:n])
			if err != nil {
				fmt.Println("Error forwarding packet to device:", err)
			}
		}
	}
}

func parseDeviceAddresses(addresses string) []string {
	return splitTrim(addresses, ",")
}

func splitTrim(s, sep string) []string {
	var result []string
	parts := strings.Split(s, sep)
	for _, part := range parts {
		result = append(result, strings.TrimSpace(part))
	}
	return result
}

