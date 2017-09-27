package main

import (
	"net"
	"fmt"
	"encoding/json"
)

type IConnection struct {
	connType string
	message string
}

func main()  {
	fmt.Println("Launching server...")


	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":14460")

	// close the listener when the app closes
	defer ln.Close()

	for {
		// accept connection on port
		conn, err := ln.Accept()

		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	d := json.NewDecoder(conn)

	var mes IConnection

	err := d.Decode(&mes)

	fmt.Println(mes, err)

	/*message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Received message", string(message))

	newmessage := strings.ToUpper(message)*/

	// send new string to receiver
	//conn.Write([]byte(newmessage))
}
