package protocol

import (
	"github.com/DeathHand/smpp/pdu"
	"net"
)

type Transmitter struct {
	W *Writer
	P *chan pdu.Pdu
}

func NewTransmitter(c *net.TCPConn, t *chan pdu.Pdu) *Transmitter {
	return &Transmitter{
		W: NewWriter(c),
		P: t,
	}
}

func (t *Transmitter) Transmit() {

}
