package node

import (
	"net"
	"pad/PADLabs/lab2/common"
	"encoding/json"
	"log"
)

func Start(connectedNodesAddrs []int, id string) {
	Addr, err := net.ResolveUDPAddr("udp", common.Addr + ":12450")
	common.CheckError(err)
	go startUDP(Addr, id, connectedNodesAddrs)
	go startTCP(connectedNodesAddrs, id)
}

func startUDP(serverAddress *net.UDPAddr, id string, connectedNodes []int) {
	conn, err := net.ListenUDP("udp", serverAddress)
	common.CheckError(err)
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buf)
		var request common.Request
		err = json.Unmarshal(buf[0:n], &request)
		common.CheckError(err)
		switch request.Msg {
		case common.GET_NODES_INFO:
			sendMetaData(request, id, connectedNodes)
		default:
			log.Print("Incorrect message type ", request.Msg)
		}

	}
}

func startTCP(connectedNodes []int, id string) {
	listen, err := net.ListenTCP("tcp", common.TCPAddr + ":" + id)
	common.CheckError(err)
	// defer listen.Close()
}

func sendMetaData(request common.Request, id string, connectedNodes []int) {
	conn, err := net.Dial("udp", request.ResponseAddr)
	common.CheckError(err)
	nodeInfo := common.NodeInfo{
		Addr: ":" + id,
		NodesCount: len(connectedNodes),
	}
	marshal, err := json.Marshal(nodeInfo)
	common.CheckError(err)
	conn.Write(marshal)
}
