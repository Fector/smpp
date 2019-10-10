package pdu

import "github.com/DeathHand/smpp/protocol"

type SubmitSmResp struct {
	Header *Header
	Body   *SmRespBody
}

// NewSubmitSmResp creates new SUBMIT_SM_RESP operation
func NewSubmitSmResp(sequence uint32, status uint32, body *SmRespBody) *SubmitSmResp {
	return &SubmitSmResp{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.SubmitSmResp,
			CommandStatus:  status,
			SequenceNumber: sequence,
		},
		Body: body,
	}
}
