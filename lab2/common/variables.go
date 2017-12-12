package common

const (
	Addr = "224.0.0.224"
	TCPAddr = "0.0.0.0"
	ProxyPort = "8001"
	ProxyRespPort = "8088"
	UDPProxyPort = "8089"
	GET_NODES_INFO = "GET_NODES_INFO"
)

type Request struct {
	ResponseAddr string
	Msg          string
}

type Character struct {
	Nickname string `json:"nickname"`
	Location string `json:"location"`
	CoordX int 	`json:"coordX"`
	CoordY int 	`json:"coordY"`
}

type NodeInfo struct {
	Addr 		string
	NodesCount 	int
}

type NodeDataRequest struct {
	Type string
}
