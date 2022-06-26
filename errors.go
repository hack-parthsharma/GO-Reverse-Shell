package gorevshell

import (
	"log"
)

func termErr(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err.Error())
	}
}
