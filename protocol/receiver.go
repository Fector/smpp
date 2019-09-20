package protocol

import (
	"github.com/DeathHand/smpp/pdu"
	"net"
)

type Receiver struct {
	R *Reader
	P *chan pdu.Pdu
}

func NewReceiver(c *net.TCPConn, r *chan pdu.Pdu) *Receiver {
	return &Receiver{
		R: NewReader(c),
		P: r,
	}
}

func (r *Receiver) Receive() {

}
