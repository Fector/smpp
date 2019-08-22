package protocol

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/DeathHand/smpp/pdu"
	"io"
	"net"
	"strings"
)

// Reader is a SMPP v3.4 protocol packet reader
type Reader struct {
	*bufio.Reader
}

// NewReader creates new Reader
func NewReader(conn *net.TCPConn) *Reader {
	return &Reader{Reader: bufio.NewReader(conn)}
}

// readInt reads 4-byte Integer
func (r *Reader) readInt(buffer *bytes.Buffer) (uint32, error) {
	p := make([]byte, 4)
	_, err := buffer.Read(p)
	if err != nil {
		return 0, nil
	}
	return binary.BigEndian.Uint32(p), nil
}

// readVarOctString reads various length C-octet string
func (r *Reader) readVarOctString(length int, buffer *bytes.Buffer) (string, error) {
	var builder strings.Builder
	p := make([]byte, 1)
	for i := 0; i < length; i++ {
		_, err := buffer.Read(p)
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

// readFixedOctString reads fixed length C-octet string
func (r *Reader) readFixedOctString(length int, buffer *bytes.Buffer) (string, error) {
	var builder strings.Builder
	p := make([]byte, 1)
	for i := 0; i < length; i++ {
		_, err := buffer.Read(p)
		if err != nil {
			return "", err
		}
		builder.Write(p)
		if p[0] == 0 {
			break
		}
	}
	return builder.String(), nil
}

// readString reads fixed length string
func (r *Reader) readString(length int, buffer *bytes.Buffer) (string, error) {
	var builder strings.Builder
	p := make([]byte, 1)
	for i := 0; i < length; i++ {
		_, err := buffer.Read(p)
		if err != nil {
			return "", err
		}
		builder.Write(p)
	}
	return builder.String(), nil
}

// readHeader reads PDU header struct from buffer
func (r *Reader) readHeader(buffer *bytes.Buffer) (*pdu.Header, error) {
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

// readTlv reads Tlv tags
func (r *Reader) readTlv(buffer *bytes.Buffer) (*[]pdu.Tlv, error) {
	if buffer.Len() == 0 {
		return nil, nil
	}

	return nil, nil
}

// readBindBody reads bind body for bind operations
func (r *Reader) readBindBody(buffer *bytes.Buffer) (*pdu.BindBody, error) {
	return nil, nil
}

// readSmBody reads short message body
func (r *Reader) readSmBody(buffer *bytes.Buffer) (*pdu.SmBody, error) {
	return nil, nil
}

// readSmRespBody reads sm operations response body
func (r *Reader) readSmRespBody(buffer *bytes.Buffer) (*pdu.SmRespBody, error) {
	return nil, nil
}

// ReadPduBuffer is a PDU packet reader
func (r *Reader) ReadPacket() (*[]byte, error) {
	var packet []byte
	p := make([]byte, 4)
	n, err := io.ReadFull(r, p)
	if n < len(p) {
		return nil, err
	}
	packet = append(packet, p...)
	p = make([]byte, binary.BigEndian.Uint32(p)-4)
	n, err = io.ReadFull(r, p)
	if n < len(p) {
		return nil, err
	}
	packet = append(packet, p...)
	return &packet, nil
}

// ReadPdu is a PDU buffer reader to PDU struct
func (r *Reader) ReadPdu(packet *[]byte) (pdu.Pdu, error) {
	buffer := bytes.NewBuffer(*packet)
	header, err := r.readHeader(buffer)
	if err != nil {
		return nil, err
	}
	switch header.CommandId {
	case SubmitSm:
		body, err := r.readSmBody(buffer)
		if err != nil {
			return nil, err
		}
		tlv, err := r.readTlv(buffer)
		if err != nil {
			return nil, err
		}
		return &pdu.SubmitSm{
			Header: header,
			Body:   body,
			Tlv:    tlv,
		}, nil
	case DeliverSm:
		body, err := r.readSmBody(buffer)
		if err != nil {
			return nil, err
		}
		tlv, err := r.readTlv(buffer)
		if err != nil {
			return nil, err
		}
		return &pdu.DeliverSm{
			Header: header,
			Body:   body,
			Tlv:    tlv,
		}, nil
	case SubmitSmResp:
		body, err := r.readSmRespBody(buffer)
		if err != nil {
			return nil, err
		}
		return &pdu.SubmitSmResp{
			Header: header,
			Body:   body,
		}, nil
	case DataSm:
		body, err := r.readSmBody(buffer)
		if err != nil {
			return nil, err
		}
		tlv, err := r.readTlv(buffer)
		if err != nil {
			return nil, err
		}
		return &pdu.DataSm{
			Header: header,
			Body:   body,
			Tlv:    tlv,
		}, nil
	case BindReceiver:
		body, err := r.readBindBody(buffer)
		if err != nil {
			return nil, err
		}
		tlv, err := r.readTlv(buffer)
		if err != nil {
			return nil, err
		}
		return &pdu.BindReceiver{
			Header: header,
			Body:   body,
			Tlv:    tlv,
		}, nil
	case BindTransmitter:
		body, err := r.readBindBody(buffer)
		if err != nil {
			return nil, err
		}
		tlv, err := r.readTlv(buffer)
		if err != nil {
			return nil, err
		}
		return &pdu.BindTransmitter{
			Header: header,
			Body:   body,
			Tlv:    tlv,
		}, nil
	case BindTransceiver:
		body, err := r.readBindBody(buffer)
		if err != nil {
			return nil, err
		}
		tlv, err := r.readTlv(buffer)
		if err != nil {
			return nil, err
		}
		return &pdu.BindTransceiver{
			Header: header,
			Body:   body,
			Tlv:    tlv,
		}, nil
	}
	return nil, errors.New(fmt.Sprintf("Unknown command %x", header.CommandId))
}
