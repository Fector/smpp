package pdu

type DeliverSm struct {
	Header *Header
	Body   *SmBody
	Tlv    *[]Tlv
}
