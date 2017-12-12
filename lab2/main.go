package main

import (
	"pad/PADLabs/lab2/node"
	"pad/PADLabs/lab2/common"
	"pad/PADLabs/lab2/proxy"
	"net"
	"encoding/json"
	"bufio"
	"log"
	"time"
)

func main()  {
	go proxy.StartMediator()
	runNodes()
	time.Sleep(time.Second * 10)
	dial := tcpDial()
	marshal, _ := json.Marshal(dial)
	log.Println("Dial " + string(marshal) + "\n")
}

func runNodes() {
	// 14400
	nodes := []int{14401, 14403}
	go node.Start([]common.Character {
		{
			"Undertaker",
			"Lorencia",
			96,
			30,
		},
	}, nodes, "14400")
	// 14401
	nodes = []int{14400, 14402}
	go node.Start([]common.Character {
		{
			"Warmashine",
			"Davias",
			51,
			211,
		},
	}, nodes, "14401")
	// 14402
	nodes = []int{14401, 14403}
	go node.Start([]common.Character {
		{
			"Mortarion",
			"Dungeon",
			17,
			98,
		},
	}, nodes, "14402")
	// 14403
	nodes = []int{14400, 14402}
	go node.Start([]common.Character {
		{
			"Boreus",
			"Davias",
			20,
			29,
		},
	}, nodes, "14403")
	// 14404
	nodes = []int{14405}
	go node.Start([]common.Character {
		{
			"Shooter",
			"Tarkan",
			178,
			17,
		},
	}, nodes, "14404")
	// 14405
	nodes = []int{14404}
	go node.Start([]common.Character {
		{
			"21Savage",
			"Atlanta",
			50,
			60,
		},
	}, nodes, "14405")
}

func tcpDial() []common.Character {
	conn, err := net.Dial("tcp", common.TCPAddr + ":" + common.ProxyPort)
	common.CheckError(err)

	request := common.NodeDataRequest{
		Type: "get_data",
	}

	marshal, _ := json.Marshal(request)
	conn.Write(marshal)

	for {
		msg, _ := bufio.NewReader(conn).ReadString('\n')
		log.Println("Client msg" + msg)
	}
}
