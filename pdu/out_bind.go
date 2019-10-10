package pdu

import "github.com/DeathHand/smpp/protocol"

type OutBindBody struct {
	SystemId string
	Password string
}

type OutBind struct {
	Pdu
	Header *Header
	Body   *OutBindBody
}

// NewOutBind creates new OUTBIND operation
func NewOutBind(sequence uint32, body *OutBindBody) *OutBind {
	return &OutBind{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.Outbind,
			CommandStatus:  protocol.EsmeRok,
			SequenceNumber: sequence,
		},
		Body: body,
	}
}
