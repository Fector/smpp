package io

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/DeathHand/smpp/pdu"
	"strings"
)

type Reader struct {
}

func (r *Reader) readInt(buffer *bytes.Buffer) (uint32, error) {
	b := make([]byte, 4)

	if buffer.Len() < len(b) {
		return 0, errors.New("Not enough length for 4-byte int ")
	}

	l, err := buffer.Read(b)

	if l < len(b) {
		return 0, errors.New("Unable to read 4 bytes for int ")
	}

	if err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint32(b), nil
}

func (r *Reader) readVarOctString(buffer *bytes.Buffer, length int) (string, error) {
	var builder strings.Builder

	for i := 0; i < length; i++ {
		b := make([]byte, 1)

		l, err := buffer.Read(b)

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

func (r *Reader) readFixedOctString(buffer *bytes.Buffer, length int) (string, error) {
	var builder strings.Builder

	for i := 0; i < length; i++ {
		b := make([]byte, 1)

		l, err := buffer.Read(b)

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

func (r *Reader) readString(buffer *bytes.Buffer, length int) (string, error) {
	var builder strings.Builder

	for i := 0; i < length; i++ {
		b := make([]byte, 1)

		l, err := buffer.Read(b)

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

func (r *Reader) ReadHeader(buffer *bytes.Buffer) (*pdu.Header, error) {
	commandLength, err := r.readInt(buffer)

	if err != nil {
		return nil, err
	}

	commandId, err := r.readInt(buffer)

	if err != nil {
		return nil, err
	}

	commandStatus, err := r.readInt(buffer)

	if err != nil {
		return nil, err
	}

	sequenceNumber, err := r.readInt(buffer)

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

func (r *Reader) ReadBody(header pdu.Header, buffer *bytes.Buffer) (*pdu.Body, error) {
	return nil, nil
}
