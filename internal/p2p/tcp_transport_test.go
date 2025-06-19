package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	tcpOption := TCPTransportOptions{
		ListenAddress: ":4000",
		ShakeHands:    NOPHandshake,
		Decoder:       DefaultDecoder{},
	}
	listenAddr := ":4000"
	tr := NewTCPTransport(tcpOption)
	assert.Equal(t, tr.ListenAddress, listenAddr)

	assert.Nil(t, tr.ListenAndAccept())
}
