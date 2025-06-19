package internal

import (
	"log"

	"github.com/jhands0/gDFS/internal/p2p"
)

type GDFS struct{}

func (g *GDFS) Init() {
	to := p2p.TCPTransportOptions{
		ListenAddress: ":3000",
		ShakeHands:    p2p.NOPHandshake,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(to)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
