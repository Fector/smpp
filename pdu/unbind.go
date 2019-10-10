package pdu

import "github.com/DeathHand/smpp/protocol"

type Unbind struct {
	Pdu
	Header *Header
}

// NewSubmitSmResp creates new UNBIND operation
func NewUnbind(sequence uint32) *Unbind {
	return &Unbind{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.Unbind,
			CommandStatus:  protocol.EsmeRok,
			SequenceNumber: sequence,
		},
	}
}
