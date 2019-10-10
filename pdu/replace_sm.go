package pdu

import "github.com/DeathHand/smpp/protocol"

type ReplaceSm struct {
	Pdu
	Header *Header
	Body   *SmBody
	Tlv    *[]Tlv
}

// NewReplaceSm creates new REPLACE_SM operation
func NewReplaceSm(sequence uint32, body *SmBody) *ReplaceSm {
	return &ReplaceSm{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.ReplaceSm,
			CommandStatus:  protocol.EsmeRok,
			SequenceNumber: sequence,
		},
		Body: body,
	}
}
