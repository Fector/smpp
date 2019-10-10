package pdu

import "github.com/DeathHand/smpp/protocol"

type EnquireLink struct {
	Pdu
	Header *Header
}

// NewEnquireLink creates new ENQUIRE_LINK operation
func NewEnquireLink(sequence uint32) *EnquireLink {
	return &EnquireLink{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.EnquireLink,
			CommandStatus:  protocol.EsmeRok,
			SequenceNumber: sequence,
		},
	}
}
