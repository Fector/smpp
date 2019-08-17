package pdu

type BindTransmitterResp struct {
	Pdu
	Header *Header
	Body   *BindRespBody
	Tlv    []*Tlv
}
