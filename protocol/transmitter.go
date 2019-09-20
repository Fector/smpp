package protocol

import (
	"github.com/DeathHand/smpp/pdu"
	"net"
)

type Transmitter struct {
	writer *Writer
	t      *chan pdu.Pdu
}

func NewTransmitter(c *net.TCPConn, t *chan pdu.Pdu) *Transmitter {
	return &Transmitter{
		writer: NewWriter(c),
		t:      t,
	}
}

func (t *Transmitter) Transmit() {

}
