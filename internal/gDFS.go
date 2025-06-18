package internal

import (
	"log"

	"github.com/jhands0/gDFS/internal/p2p"
)

type GDFS struct{}

func (g *GDFS) Init() {
	tr := p2p.NewTCPTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
