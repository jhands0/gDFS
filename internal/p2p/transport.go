package p2p

// Peer represents a remote node over any type of connection
type Peer interface{}

// Transport represents any type of connection e.g. TCP, UDP, websockets, gRPC
type Transport interface {
	ListenAndAccept() error
}
