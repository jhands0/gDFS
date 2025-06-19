package p2p

import "errors"

var ErrInvalidHandshake = errors.New("invalid handshake")

type HandshakeFunc func(Peer) error

func NOPHandshake(Peer) error { return nil }
