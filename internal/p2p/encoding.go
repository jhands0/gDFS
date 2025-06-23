package p2p

import "io"

type Decoder interface {
	Decode(io.Reader, any) error
}

type DefaultDecoder struct{}

func (dec DefaultDecoder) Decode(r io.Reader, msg any) error {
	peekBuf := make([]byte, 1)
	if _, err := r.Read(peekBuf); err != nil {
		return err
	}

	return nil
}
