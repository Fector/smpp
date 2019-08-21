package pdu

type BindTransmitter struct {
	Pdu
	Header *Header
	Body   *BindBody
	Tlv    *[]Tlv
}
