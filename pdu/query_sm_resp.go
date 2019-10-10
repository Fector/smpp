package pdu

import "github.com/DeathHand/smpp/protocol"

type QuerySmRespBody struct {
	MessageId    string
	FinalDate    string
	MessageState uint32
	ErrorCode    uint32
}

type QuerySmResp struct {
	Pdu
	Header *Header
	Body   *QuerySmRespBody
}

// NewQuerySmResp creates new QUERY_SM_RESP operation
func NewQuerySmResp(sequence uint32, status uint32, body *QuerySmRespBody) *QuerySmResp {
	return &QuerySmResp{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.QuerySmResp,
			CommandStatus:  status,
			SequenceNumber: sequence,
		},
		Body: body,
	}
}
