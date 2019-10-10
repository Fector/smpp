package pdu

import "github.com/DeathHand/smpp/protocol"

type UnbindResp struct {
	Pdu
	Header *Header
}

// NewUnbindResp creates new UNBIND_RESP operation
func NewUnbindResp(sequence uint32, status uint32) *UnbindResp {
	return &UnbindResp{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.UnbindResp,
			CommandStatus:  status,
			SequenceNumber: sequence,
		},
	}
}
