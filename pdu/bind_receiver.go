package pdu

import "github.com/DeathHand/smpp/protocol"

type BindReceiver struct {
	Pdu
	Header *Header
	Body   *BindBody
	Tlv    *[]Tlv
}

// NewBindReceiver creates new BIND_RECEIVER operation
func NewBindReceiver(sequence uint32, body *BindBody) *BindReceiver {
	return &BindReceiver{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.BindReceiver,
			CommandStatus:  protocol.EsmeRok,
			SequenceNumber: sequence,
		},
		Body: body,
	}
}
