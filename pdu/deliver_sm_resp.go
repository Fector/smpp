package pdu

type DeliverSmResp struct {
	Pdu
	Header *Header
	Body   *SmRespBody
}
