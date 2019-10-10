package pdu

import "github.com/DeathHand/smpp/protocol"

type BindTransmitter struct {
	Pdu
	Header *Header
	Body   *BindBody
	Tlv    *[]Tlv
}

// NewBindTransmitter creates new BIND_TRANSMITTER operation
func NewBindTransmitter(sequence uint32, body *BindBody) *BindTransmitter {
	return &BindTransmitter{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.BindTransmitter,
			CommandStatus:  protocol.EsmeRok,
			SequenceNumber: sequence,
		},
		Body: body,
	}
}
