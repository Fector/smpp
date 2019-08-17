package pdu

type BindTransceiver struct {
	Pdu
	Header *Header
	Body   *BindBody
	Tlv    []*Tlv
}
