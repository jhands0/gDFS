package internal

import (
	"log"

	"github.com/jhands0/gDFS/internal/p2p"
)

type gDFS struct{}

func (g *gDFS) Init() {
	tr := p2p.NewTCPTransport(":3000")
	log.Fatal(tr.ListenAndAccept())
}
