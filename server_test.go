package main

import (
	"net"
	"testing"
)

func TestServerInit(t *testing.T) {
	server := Server{port: 10001}
	server.Start()
	defer server.Close()
}

func TestUDPServer(t *testing.T) {
	server := Server{port: 10001}
	server.Start()
	defer server.Close()

	serverAddress, err := net.ResolveUDPAddr("udp", "127.0.0.1:10001")
	if err != nil {
		t.Fatal(err)
	}
	localAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	conn, err1 := net.DialUDP("udp", localAddr, serverAddress)
	if err1 != nil {
		t.Fatal(err1)
	}
	defer conn.Close()
}
