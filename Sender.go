package main

import (
	"net"
	"fmt"
	"encoding/json"
)

func main() {
	conn, _ := net.Dial("tcp", "0.0.0.0:14460")

	fmt.Fprintf(conn, "valet\n")

	c := &IConnection {"client", "valet"}

	encoder := json.NewEncoder(conn)

	encoder.Encode(c)

/*	message, _ := bufio.NewReader(conn).ReadString('\n')
	if message != "" {
		fmt.Println("Message from brocker:", message)
	}*/

}
