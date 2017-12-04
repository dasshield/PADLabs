package common

const (
	Addr = "224.0.0.1"
	TCPAddr = "0.0.0.0"
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
