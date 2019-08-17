package pdu

type BindTransceiverResp struct {
	Pdu
	Header *Header
	Body   *BindRespBody
	Tlv    []*Tlv
}
