package protocol

import (
	"github.com/DeathHand/smpp/pdu"
	"net"
	"time"
)

type Writer struct {
	*net.TCPConn
}

func NewWriter(conn *net.TCPConn) *Writer {
	return &Writer{TCPConn: conn}
}

func (w *Writer) SetBufferSize(bytes int) error {
	return w.TCPConn.SetWriteBuffer(bytes)
}

func (w *Writer) SetTimeout(t time.Time) error {
	return w.TCPConn.SetWriteDeadline(t)
}

func (w *Writer) writeInt(val uint32) error {
	return nil
}

func (w *Writer) writeString(val string) error {
	return nil
}

func (w *Writer) writeOctString(val string) error {
	return nil
}

func (w *Writer) WritePdu(pdu *pdu.Pdu) (*pdu.Pdu, error) {
	return pdu, nil
}