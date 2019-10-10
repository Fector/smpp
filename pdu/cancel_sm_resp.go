package pdu

import "github.com/DeathHand/smpp/protocol"

type CancelSmResp struct {
	Pdu
	Header *Header
}

// NewNewCancelSm creates new CANCEL_SM_RESP operation
func NewNewCancelSm(sequence uint32, status uint32) *CancelSmResp {
	return &CancelSmResp{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.CancelSmResp,
			CommandStatus:  status,
			SequenceNumber: sequence,
		},
	}
}
