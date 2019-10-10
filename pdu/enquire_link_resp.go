package pdu

import "github.com/DeathHand/smpp/protocol"

type EnquireLinkResp struct {
	Pdu
	Header *Header
}

// NewEnquireLinkResp creates new ENQUIRE_LINK_RESP operation
func NewEnquireLinkResp(sequence uint32) *EnquireLinkResp {
	return &EnquireLinkResp{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.EnquireLinkResp,
			CommandStatus:  protocol.EsmeRok,
			SequenceNumber: sequence,
		},
	}
}
