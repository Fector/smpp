package pdu

type QuerySmRespBody struct {
	MessageId    string
	FinalDate    string
	MessageState uint32
	ErrorCode    uint32
}

type QuerySmResp struct {
	Pdu
	Header *Header
	Body   *QuerySmRespBody
}
