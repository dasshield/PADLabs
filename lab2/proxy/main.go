package proxy

import (
	"net"
	"pad/PADLabs/lab2/common"
	"bufio"
	"encoding/json"
	"time"
	"log"
	"fmt"
)

var nodesList []common.NodeInfo = make([]common.NodeInfo, 0)

func StartMediator()  {
	go startTCP()

	go listenUDP()
}

func startTCP()  {
	listen, err := net.Listen("tcp", common.TCPAddr + ":" + common.ProxyPort)
	common.CheckError(err)

	for {
		conn, err := listen.Accept()
		common.CheckError(err)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))

	for {
		request, err := receiveMessage(rw)
		common.CheckError(err)
		handleMessage(request, conn)
	}
}

func receiveMessage(rw *bufio.ReadWriter) (request common.NodeDataRequest, err error) {
	d := json.NewDecoder(rw)
	err = d.Decode(&request)
	common.CheckError(err)
	rw.Flush()
	return request, nil
}

func getMaxNodesTcpPort() string {
	go listenUDP()

	maven := getMaven()
	fmt.Println(maven)
	return maven.Addr
}

func handleMessage(request common.NodeDataRequest, conn net.Conn)  {
	switch request.Type {
	case "get_data":
		maven := getMaxNodesTcpPort()
		fmt.Println(maven)
		nodeData := tcpDial(maven, request)
		serializeResponse(request, nodeData, conn)
	default:
		log.Println("Invalid request type")
	}
}

func serializeResponse(request common.NodeDataRequest, data []common.Character, conn net.Conn) {
	marshal, err := json.Marshal(data)
	common.CheckError(err)
	log.Println("data + " + string(marshal))
	conn.Write([]byte(string(marshal) + "\n"))
}

func doBroadcast() {
	conn, err := net.Dial("udp", common.Addr + ":" + common.UDPProxyPort)
	common.CheckError(err)

	nodesInfoRequest := common.Request{
		ResponseAddr: common.Addr + ":" + common.ProxyRespPort,
		Msg: common.GET_NODES_INFO,
	}
	marshal, err := json.Marshal(nodesInfoRequest)
	common.CheckError(err)
	conn.Write(marshal)
}

func getMaven() common.NodeInfo {
	doBroadcast()
	time.Sleep(time.Second * 10)
	maven := common.NodeInfo{}
	for _, node := range nodesList {
		if node.NodesCount > maven.NodesCount {
			maven = node
		}
	}
	fmt.Println(maven)
	return maven
}

/*func newBroadcaster() (*net.UDPConn, error) {
	addr, err := net.ResolveUDPAddr("udp", common.Addr + ":" + common.UDPProxyPort)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}*/

func listenUDP() {
	addr, addrErr := net.ResolveUDPAddr("udp", common.Addr + ":" + common.ProxyRespPort)
	common.CheckError(addrErr)
	conn, connErr := net.ListenUDP("udp", addr)
	common.CheckError(connErr)
	defer conn.Close()
	buf := make([]byte, 1024)
	for {
		n, _, err := conn.ReadFromUDP(buf)
		common.CheckError(err)
		var nodeInfo common.NodeInfo
		err = json.Unmarshal(buf[0:n], &nodeInfo)
		fmt.Println(nodeInfo)
		common.CheckError(err)
		nodesList = append(nodesList, nodeInfo)
	}
}

func tcpDial(node string, clientRequest common.NodeDataRequest) []common.Character {
	log.Println("dial to maven : " + common.TCPAddr + ":" + node)
	conn, err := net.Dial("tcp", common.TCPAddr + ":" + node)
	common.CheckError(err)

	request := common.NodeDataRequest{
		Type: clientRequest.Type,
	}
	marshal, _ := json.Marshal(request)
	conn.Write(marshal)
	for {
		decoder := json.NewDecoder(conn)

		var response []common.Character
		err := decoder.Decode(&response)
		common.CheckError(err)
		return response
	}
}