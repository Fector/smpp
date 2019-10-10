package pdu

import "github.com/DeathHand/smpp/protocol"

type DataSmBody struct {
	ServiceType        string
	SourceAddrTon      uint32
	SourceAddrNpi      uint32
	SourceAddr         string
	DestAddrTon        uint32
	DestAddrNpi        uint32
	DestinationAddr    string
	EsmClass           uint32
	RegisteredDelivery uint32
	DataCoding         uint32
}

type DataSm struct {
	Header *Header
	Body   *DataSmBody
	Tlv    *[]Tlv
}

// NewDataSm creates new DATA_SM operation
func NewDataSm(sequence uint32, body *DataSmBody, tlv *[]Tlv) *DataSm {
	return &DataSm{
		Header: &Header{
			CommandLength:  0,
			CommandId:      protocol.DataSm,
			CommandStatus:  protocol.EsmeRok,
			SequenceNumber: sequence,
		},
		Body: body,
		Tlv:  tlv,
	}
}
