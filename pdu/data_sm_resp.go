package pdu

import "github.com/DeathHand/smpp/protocol"

type DataSmResp struct {
	Pdu
	Header *Header
	Body   *SmRespBody
	Tlv    *[]Tlv
}

// NewDataSmResp creates new DATA_SM_RESP operation
func NewDataSmResp(sequence uint32, status uint32, body *SmRespBody, tlv *[]Tlv) *DataSmResp {
	return &DataSmResp{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.DataSmResp,
			CommandStatus:  status,
			SequenceNumber: sequence,
		},
		Body: body,
		Tlv:  tlv,
	}
}
