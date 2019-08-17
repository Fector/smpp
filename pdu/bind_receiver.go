package pdu

type BindReceiver struct {
	Pdu
	Header *Header
	Body   *BindBody
	Tlv    []*Tlv
}
