package pdu

import "github.com/DeathHand/smpp/protocol"

type BindTransceiverResp struct {
	Pdu
	Header *Header
	Body   *BindRespBody
	Tlv    []*Tlv
}

// NewBindTransceiverResp creates new BIND_TRANSCEIVER_RESP operation
func NewBindTransceiverResp(sequence uint32, status uint32, body *BindRespBody) *BindTransceiverResp {
	return &BindTransceiverResp{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.BindTransceiverResp,
			CommandStatus:  status,
			SequenceNumber: sequence,
		},
		Body: body,
	}
}
