package pdu

import "github.com/DeathHand/smpp/protocol"

type BindTransceiver struct {
	Pdu
	Header *Header
	Body   *BindBody
	Tlv    *[]Tlv
}

// NewBindTransceiver creates new BIND_TRANSCEIVER operation
func NewBindTransceiver(sequence uint32, body *BindBody) *BindTransceiver {
	return &BindTransceiver{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.BindTransceiver,
			CommandStatus:  protocol.EsmeRok,
			SequenceNumber: sequence,
		},
		Body: body,
	}
}
