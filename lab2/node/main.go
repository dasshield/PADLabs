package node

import (
	"net"
	"pad/PADLabs/lab2/common"
	"encoding/json"
	"log"
	"bufio"
)

func Start(values []common.Character, connectedNodesAddrs []int, id string) {
	Addr, err := net.ResolveUDPAddr("udp", common.Addr + ":12450")
	common.CheckError(err)
	go startUDP(Addr, id, connectedNodesAddrs)
	go startTCP(connectedNodesAddrs, id, values)
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

func startTCP(connectedNodes []int, id string, values []common.Character) {
	listen, err := net.ListenTCP("tcp", common.TCPAddr + ":" + id)
	log.Println("Start TCP " + common.TCPAddr + ":" + id)
	common.CheckError(err)
	// defer listen.Close()

	for {
		conn, err := listen.Accept()
		common.CheckError(err)
		go handleTCPConn(conn, connectedNodes, values)
	}
}

func handleTCPConn(conn net.Conn, nodes []int, values []common.Character) {
	defer conn.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	for {
		request, err := receiveMsg(rw)
		common.CheckError(err)
		handleMessage(request, conn, nodes, values)
	}
}

func receiveMsg(rw *bufio.ReadWriter) (request common.NodeDataRequest, err error) {
	d := json.NewDecoder(rw)
	err = d.Decode(&request)
	common.CheckError(err)
	rw.Flush()
	return request, nil
}

func handleMessage(request common.NodeDataRequest, conn net.Conn, nodes []int, valuers []common.Character)  {
	switch request.Type {
	case "get_value":
		
	}
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
