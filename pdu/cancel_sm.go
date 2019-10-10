package pdu

import "github.com/DeathHand/smpp/protocol"

type CancelSmBody struct {
	Body
	ServiceType     string
	MessageId       string
	SourceAddrTon   uint32
	SourceAddrNpi   uint32
	SourceAddr      string
	DestAddrTon     uint32
	DestAddrNpi     uint32
	DestinationAddr string
}

type CancelSm struct {
	Pdu
	Header *Header
	Body   *CancelSmBody
}

// NewCancelSm creates new CANCEL_SM operation
func NewCancelSm(sequence uint32, body *CancelSmBody) *CancelSm {
	return &CancelSm{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.CancelSm,
			CommandStatus:  protocol.EsmeRok,
			SequenceNumber: sequence,
		},
		Body: body,
	}
}
