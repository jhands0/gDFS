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

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	shakeHands    HandshakeFunc
	decoder       Decoder

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
		shakeHands:    NOPHandshake,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
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

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	var err error

	peer := NewTCPPeer(conn, true)

	if err = t.shakeHands(peer); err != nil {
		fmt.Printf("TCP error: %s\n", err)
		conn.Close()
		return
	}

	// Message read loop
	decoderErrorCount := 0
	msg := &Temp{}
	for {
		if err := t.decoder.Decode(conn, msg); err != nil {

			decoderErrorCount++
			if decoderErrorCount >= 4 {
				return
			}

			fmt.Printf("TCP error: %s\n", err)
			continue
		}
	}

}
