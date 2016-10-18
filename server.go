package main

import (
	"fmt"
	"net"
	"os"
)

var conn *net.UDPConn

// Server UDP server
type Server struct {
	port int
}

//Start nodoc
func (server Server) Start() {
	serverAddress, err := net.ResolveUDPAddr("udp", ":10001")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	conn, err = net.ListenUDP("udp", serverAddress)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}

//Close nodoc
func (server Server) Close() {
	conn.Close()
}
