package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn net.Conn

	// outbound is True if we dial the other peer
	// outbound is False if we accept the other peer
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOptions struct {
	ListenAddress string
	ShakeHands    HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOptions
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(options TCPTransportOptions) *TCPTransport {
	return &TCPTransport{
		TCPTransportOptions: options,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		fmt.Printf("new incoming connection %+v\n", conn)

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	var err error

	peer := NewTCPPeer(conn, true)

	if err = t.ShakeHands(peer); err != nil {
		fmt.Printf("TCP error: %s\n", err)
		conn.Close()
		return
	}

	// Message read loop
	decoderErrorCount := 0
	msg := &Message{}
	buf := make([]byte, 2000)
	for {
		m, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("TCP error: %s\n", err)
		}

		if err := t.Decoder.Decode(conn, msg); err != nil {

			decoderErrorCount++
			if decoderErrorCount >= 4 {
				return
			}

			fmt.Printf("TCP error: %s\n", err)
			continue
		}

		fmt.Printf("message: %+v\n", buf[:m])
	}

}
