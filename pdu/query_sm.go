package pdu

import "github.com/DeathHand/smpp/protocol"

type QuerySmBody struct {
	MessageId     string
	SourceAddrTon uint32
	SourceAddrNpi uint32
	SourceAddr    string
}

type QuerySm struct {
	Pdu
	Header *Header
	Body   *QuerySmBody
}

// NewQuerySm creates new QUERY_SM operation
func NewQuerySm(sequence uint32, body *QuerySmBody) *QuerySm {
	return &QuerySm{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.QuerySm,
			CommandStatus:  protocol.EsmeRok,
			SequenceNumber: sequence,
		},
		Body: body,
	}
}
