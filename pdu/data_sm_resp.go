package pdu

type DataSmResp struct {
	Pdu
	Header *Header
	Body   *SmRespBody
	Tlv    *[]Tlv
}
