package main

import (
	"pad/PADLabs/lab1/common"
	"bufio"
	"os"
	"fmt"
	"time"
)

func main()  {
	pub := common.NewPub()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Topic: ")
	topic, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Cannot read inputed topic")
	}

	fmt.Print("Enter message: ")
	msg, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Cannot read inputed message")
	}


	for {



		pub.Publish(topic, msg)

		time.Sleep(time.Second * 10)
	}
}
