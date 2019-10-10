package pdu

import "github.com/DeathHand/smpp/protocol"

type BindTransmitterResp struct {
	Pdu
	Header *Header
	Body   *BindRespBody
	Tlv    []*Tlv
}

// NewBindTransmitterResp creates new BIND_TRANSMITTER_RESP operation
func NewBindTransmitterResp(sequence uint32, status uint32, body *BindRespBody) *BindTransmitterResp {
	return &BindTransmitterResp{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.BindTransmitterResp,
			CommandStatus:  status,
			SequenceNumber: sequence,
		},
		Body: body,
	}
}
