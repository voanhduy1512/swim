package main

import "net"

//Status nodoc
type Status int

const (
	//ALIVE nodoc
	ALIVE STATUS = 1 + iota
	//FAILED nodoc
	FAILED
)

//Member nodoc
type Member struct {
	addr   *net.UDPAddr
	status int
}

//Node nodoc
type Node struct {
	period  int
	members []Member
}
