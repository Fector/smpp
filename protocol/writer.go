package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/DeathHand/smpp/pdu"
	"net"
)

// Reader is a SMPP v3.4 protocol packet writer
type Writer struct {
	*bufio.Writer
}

// NewWriter is the writer constructor
func NewWriter(conn *net.TCPConn) *Writer {
	return &Writer{Writer: bufio.NewWriter(conn)}
}

// writeInt is a int type writer
func (w *Writer) writeInt(val uint32, buffer *bytes.Buffer) (int, error) {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, val)
	return buffer.Write(b)
}

// writeInt is a string type writer
func (w *Writer) writeString(val string, length int, buffer *bytes.Buffer) (int, error) {
	return buffer.WriteString(val)
}

// WriteHeader is a Pdu Header writer
func (w *Writer) WriteSmBody(body *pdu.SmBody, buffer *bytes.Buffer) (int, error) {
	return 0, nil
}

// WriteHeader is a Pdu Header writer
func (w *Writer) WriteHeader(header *pdu.Header, buffer *bytes.Buffer) (int, error) {
	n, err := w.writeInt(header.CommandLength, buffer)
	if err != nil {
		return n, err
	}
	n, err = w.writeInt(header.CommandId, buffer)
	if err != nil {
		return n, err
	}
	n, err = w.writeInt(header.CommandStatus, buffer)
	if err != nil {
		return n, err
	}
	return w.writeInt(header.SequenceNumber, buffer)
}

// WritePdu is the PDU writer
func (w *Writer) WritePdu(packet pdu.Pdu) (pdu.Pdu, error) {
	var buffer bytes.Buffer
	switch p := packet.(type) {
	case pdu.SubmitSm:
		n, err := w.WriteHeader(p.Header, &buffer)
		if err != nil {
			return nil, err
		}
		p.Header.CommandLength += uint32(n)
		n, err = w.WriteSmBody(p.Body, &buffer)
		if err != nil {
			return nil, err
		}
		p.Header.CommandLength += uint32(n)
		n, err = w.Write(buffer.Bytes())
		if uint32(n) != p.Header.CommandLength {
			return nil, err
		}
		return &p, nil
	}
	return packet, errors.New("Unknown packet type ")
}
