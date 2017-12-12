package node

import (
	"net"
	"pad/PADLabs/lab2/common"
	"encoding/json"
	"log"
	"bufio"
	"strconv"
	"fmt"
)

func Start(values []common.Character, connectedNodesAddrs []int, id string) {
	Addr, err := net.ResolveUDPAddr("udp", common.Addr + ":" + common.UDPProxyPort)
	common.CheckError(err)
	go startUDP(Addr, id, connectedNodesAddrs)
	go startTCP(connectedNodesAddrs, id, values)
}

func startUDP(serverAddress *net.UDPAddr, id string, connectedNodes []int) {
	serverConn, err := net.ListenUDP("udp", serverAddress)
	fmt.Println(serverConn)
	defer serverConn.Close()
	common.CheckError(err)
	buf := make([]byte, 1024)
	for {
		n, _, err := serverConn.ReadFromUDP(buf)
		var request common.Request
		err = json.Unmarshal(buf[0:n], &request)
		fmt.Println(" fsd",request)
		common.CheckError(err)
		if request.Msg == common.GET_NODES_INFO {
			sendMetaData(request, id, connectedNodes)
		} else {
			log.Print("incorrect message " + request.Msg)
		}
	}
}

func startTCP(connectedNodes []int, id string, values []common.Character) {
	listen, err := net.Listen("tcp", common.TCPAddr + ":" + id)
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

func handleMessage(request common.NodeDataRequest, conn net.Conn, nodes []int, values []common.Character)  {
	switch request.Type {
	case "get_data":
		var answer []common.Character

		answer = collectNodeData(nodes, values)
		marshal, err := json.Marshal(answer)
		common.CheckError(err)
		conn.Write(marshal)
	default:
		log.Println("[Node] Invalid request.")
	}
}

func collectNodeData(nodes []int, values []common.Character) []common.Character  {
	var answer []common.Character
	for _, nodeData := range nodes {
		for _, part := range tcpDial(nodeData) {
			answer = appendData(answer, part)
		}
	}
	for _, node := range values {
		answer = append(answer, node)
	}
	return answer
}

func appendData(currentAnswer []common.Character, part common.Character) []common.Character {
	for _, elem := range currentAnswer {
		if elem == part {
			return currentAnswer
		}
	}
	return append(currentAnswer, part)
}

func sendMetaData(request common.Request, id string, connectedNodes []int) {
	conn, err := net.Dial("udp", request.ResponseAddr)
	common.CheckError(err)
	nodeInfo := common.NodeInfo{
		Addr: ":" + id,
		NodesCount: len(connectedNodes),
	}
	marshal, err := json.Marshal(nodeInfo)
	fmt.Println(marshal)
	common.CheckError(err)
	conn.Write(marshal)
}

func tcpDial(port int) []common.Character {
	conn, err := net.Dial("tcp", common.Addr + ":" + strconv.Itoa(port))
	common.CheckError(err)
	request := common.NodeDataRequest{
		Type: "get_data",
	}
	marshal, _ := json.Marshal(request)
	conn.Write(marshal)

	var resp []common.Character
	decoder := json.NewDecoder(conn)
	err = decoder.Decode(&resp)
	common.CheckError(err)
	return resp
}
