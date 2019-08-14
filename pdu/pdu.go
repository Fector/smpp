package pdu

type Pdu struct {
	Header *Header
	Body   *Body
	Tlv    *[]Tlv
}
