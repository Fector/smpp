package pdu

import "github.com/DeathHand/smpp/protocol"

type SubmitSm struct {
	Header *Header
	Body   *SmBody
	Tlv    *[]Tlv
}

// NewSubmitSm creates new SUBMIT_SM operation
func NewSubmitSm(sequence uint32, body *SmBody, tlv *[]Tlv) *SubmitSm {
	return &SubmitSm{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.SubmitSm,
			CommandStatus:  protocol.EsmeRok,
			SequenceNumber: sequence,
		},
		Body: body,
		Tlv:  tlv,
	}
}
