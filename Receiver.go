package main

import (
	"net"
	"bufio"
	"fmt"
)

func main() {
	conn, _ := net.Dial("tcp", "0.0.0.0:14460")

	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if message != "" {
			fmt.Println("Message from brocker:", message)
		}

	}
}
