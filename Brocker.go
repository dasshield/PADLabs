package main

import (
	"net"
	"fmt"
	"bufio"
	"strings"
)

func main()  {
	fmt.Println("Launching server...")


	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":14460")

	// close the listener when the app closes
	defer ln.Close()

	for {
		// accept connection on port
		conn, _ := ln.Accept()

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Received message", string(message))

		newmessage := strings.ToUpper(message)

		// send new string to receiver
		conn.Write([]byte(newmessage))
	}
}
