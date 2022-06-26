package gorevshell

import (
	"net"
)

type GoRevShell struct {
	ip      string
	port    string
	address *net.TCPAddr
	con     *net.TCPConn
}
