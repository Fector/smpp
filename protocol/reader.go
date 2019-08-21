package protocol

import (
	"bufio"
	"encoding/binary"
	"errors"
	"github.com/DeathHand/smpp/pdu"
	"io"
	"net"
	"strings"
	"time"
)

type Reader struct {
	*net.TCPConn
	*bufio.Reader
}

func NewReader(conn *net.TCPConn) *Reader {
	return &Reader{
		TCPConn: conn,
		Reader:  bufio.NewReader(conn),
	}
}

func (r *Reader) SetBufferSize(bytes int) error {
	return r.TCPConn.SetReadBuffer(bytes)
}

func (r *Reader) SetTimeout(t time.Time) error {
	return r.TCPConn.SetReadDeadline(t)
}

func (r *Reader) readInt() (uint32, error) {
	b := make([]byte, 4)

	l, err := io.ReadFull(r, b)

	if l < len(b) {
		return 0, errors.New("Unable to read 4 bytes for int ")
	}

	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint32(b), nil
}

func (r *Reader) readVarOctString(length int) (string, error) {
	var builder strings.Builder

	for i := 0; i < length; i++ {
		b := make([]byte, 1)

		l, err := io.ReadFull(r, b)

		if l < len(b) {
			return "", errors.New("Unable to read 1 byte for string builder ")
		}

		if err != nil {
			return "", err
		}

		if b[0] == 0 {
			break
		}

		builder.Write(b)
	}

	return builder.String(), nil
}

func (r *Reader) readFixedOctString(length int) (string, error) {
	var builder strings.Builder

	for i := 0; i < length; i++ {
		b := make([]byte, 1)

		l, err := io.ReadFull(r, b)

		if err != nil {
			return "", err
		}

		if l < len(b) {
			return "", errors.New("Unable to read 1 byte for string builder ")
		}

		builder.Write(b)

		if b[0] == 0 {
			break
		}
	}

	return builder.String(), nil
}

func (r *Reader) readString(length int) (string, error) {
	var builder strings.Builder

	for i := 0; i < length; i++ {
		b := make([]byte, 1)

		l, err := io.ReadFull(r, b)

		if err != nil {
			return "", err
		}

		if l < len(b) {
			return "", errors.New("Unable to read 1 byte for string builder ")
		}

		builder.Write(b)
	}

	return builder.String(), nil
}

func (r *Reader) ReadHeader() (*pdu.Header, error) {
	commandLength, err := r.readInt()

	if err != nil {
		return nil, err
	}

	commandId, err := r.readInt()

	if err != nil {
		return nil, err
	}

	commandStatus, err := r.readInt()

	if err != nil {
		return nil, err
	}

	sequenceNumber, err := r.readInt()

	if err != nil {
		return nil, err
	}

	return &pdu.Header{
		CommandLength:  commandLength,
		CommandId:      commandId,
		CommandStatus:  commandStatus,
		SequenceNumber: sequenceNumber,
	}, nil
}

func (r *Reader) ReadBody(header pdu.Header) (*pdu.Body, error) {
	switch header.CommandId {
	case BindReceiver:
	case BindReceiverResp:
	case BindTransmitter:
	case BindTransmitterResp:
	case BindTransceiver:
	case BindTransceiverResp:
	case Unbind:
	case UnbindResp:
	case Outbind:
	case GenericNack:

	}

	return nil, nil
}

func (r *Reader) ReadPdu() (pdu.Pdu, error) {

	var p pdu.Pdu

	return &p, nil
}
