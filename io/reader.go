package io

import (
	"bufio"
	"encoding/binary"
	"errors"
	"github.com/DeathHand/smpp/pdu"
	"github.com/DeathHand/smpp/protocol"
	"net"
	"strings"
)

type Reader struct {
	*bufio.Reader
}

func NewReader(conn *net.TCPConn) *Reader {
	return &Reader{
		Reader: bufio.NewReader(conn),
	}
}

func (r *Reader) readInt() (uint32, error) {
	p := make([]byte, 4)

	l, err := r.Read(p)

	if l < len(p) {
		return 0, errors.New("Unable to read 4 bytes for int ")
	}

	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint32(p), nil
}

func (r *Reader) readVarOctString(length int) (string, error) {
	var builder strings.Builder

	for i := 0; i < length; i++ {
		p := make([]byte, 1)

		l, err := r.Read(p)

		if l < len(p) {
			return "", errors.New("Unable to read 1 byte for string builder ")
		}

		if err != nil {
			return "", err
		}

		if p[0] == 0 {
			break
		}

		builder.Write(p)
	}

	return builder.String(), nil
}

func (r *Reader) readFixedOctString(length int) (string, error) {
	var builder strings.Builder

	for i := 0; i < length; i++ {
		p := make([]byte, 1)

		l, err := r.Read(p)

		if err != nil {
			return "", err
		}

		if l < len(p) {
			return "", errors.New("Unable to read 1 byte for string builder ")
		}

		builder.Write(p)

		if p[0] == 0 {
			break
		}
	}

	return builder.String(), nil
}

func (r *Reader) readString(length int) (string, error) {
	var builder strings.Builder

	for i := 0; i < length; i++ {
		p := make([]byte, 1)

		l, err := r.Read(p)

		if err != nil {
			return "", err
		}

		if l < len(p) {
			return "", errors.New("Unable to read 1 byte for string builder ")
		}

		builder.Write(p)
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
	case protocol.BindReceiver:
	case protocol.BindReceiverResp:
	case protocol.BindTransmitter:
	case protocol.BindTransmitterResp:
	case protocol.BindTransceiver:
	case protocol.BindTransceiverResp:
	case protocol.Unbind:
	case protocol.UnbindResp:
	case protocol.Outbind:
	case protocol.GenericNack:

	}

	return nil, nil
}

func (r *Reader) ReadPdu() (*pdu.Pdu, error) {
	return &pdu.Pdu{}, nil
}
