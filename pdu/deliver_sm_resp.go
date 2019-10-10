package pdu

import "github.com/DeathHand/smpp/protocol"

type DeliverSmResp struct {
	Pdu
	Header *Header
	Body   *SmRespBody
}

// NewDataSmResp creates new DELIVER_SM_RESP operation
func NewDeliverSmResp(sequence uint32, status uint32, body *SmRespBody) *DeliverSmResp {
	return &DeliverSmResp{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.DeliverSmResp,
			CommandStatus:  status,
			SequenceNumber: sequence,
		},
		Body: body,
	}
}
