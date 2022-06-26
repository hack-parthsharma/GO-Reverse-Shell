package gorevshell

import (
	"fmt"
	"log"
	"net"
)

func (r *GoRevShell) Init(ip, port string) {
	var err error

	if ip == "" || port == "" {
		log.Fatal("[-] Provide valid argument")
	}
	socket := fmt.Sprintf("%s:%s", ip, port)

	//resolve ip
	r.address, err = net.ResolveTCPAddr("tcp", socket)
	termErr("[-] resolving: ", err)
}
