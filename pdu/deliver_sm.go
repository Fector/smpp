package pdu

import "github.com/DeathHand/smpp/protocol"

type DeliverSm struct {
	Header *Header
	Body   *SmBody
	Tlv    *[]Tlv
}

// NewDeliverSm creates new DELIVER_SM operation
func NewDeliverSm(sequence uint32, body *SmBody, tlv *[]Tlv) *DeliverSm {
	return &DeliverSm{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.DeliverSm,
			CommandStatus:  protocol.EsmeRok,
			SequenceNumber: sequence,
		},
		Body: body,
		Tlv:  tlv,
	}
}
