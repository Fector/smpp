package pdu

import "github.com/DeathHand/smpp/protocol"

type ReplaceSmResp struct {
	Pdu
	Header *Header
}

// NewReplaceSmResp creates new REPLACE_SM_RESP operation
func NewReplaceSmResp(sequence uint32, status uint32) *ReplaceSmResp {
	return &ReplaceSmResp{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.ReplaceSmResp,
			CommandStatus:  status,
			SequenceNumber: sequence,
		},
	}
}
