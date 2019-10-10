package pdu

import "github.com/DeathHand/smpp/protocol"

type GenericNack struct {
	Pdu
	Header *Header
}

// NewGenericNack creates new GENERIC_NACK operation
func NewGenericNack(sequence uint32, status uint32) *GenericNack {
	return &GenericNack{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.GenericNack,
			CommandStatus:  status,
			SequenceNumber: sequence,
		},
	}
}
