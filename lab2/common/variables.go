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

type NodeInfo struct {
	Addr 		string
	NodesCount 	int
}
