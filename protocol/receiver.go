package protocol

import (
	"github.com/DeathHand/smpp/pdu"
	"net"
)

type Receiver struct {
	reader *Reader
	r      *chan pdu.Pdu
}

func NewReceiver(c *net.TCPConn, r *chan pdu.Pdu) *Receiver {
	return &Receiver{
		reader: NewReader(c),
		r:      r,
	}
}

func (r *Receiver) Receive() {

}
