package pdu

import "github.com/DeathHand/smpp/protocol"

type BindReceiverResp struct {
	Pdu
	Header *Header
	Body   *BindRespBody
	Tlv    []*Tlv
}

// NewBindReceiverResp creates new BIND_RECEIVER_RESP operation
func NewBindReceiverResp(sequence uint32, status uint32, body *BindRespBody) *BindReceiverResp {
	return &BindReceiverResp{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.BindReceiverResp,
			CommandStatus:  status,
			SequenceNumber: sequence,
		},
		Body: body,
	}
}
