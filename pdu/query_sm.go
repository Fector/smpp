package pdu

type QuerySmBody struct {
	MessageId     string
	SourceAddrTon uint32
	SourceAddrNpi uint32
	SourceAddr    string
}

type QuerySm struct {
	Pdu
	Header *Header
	Body   *QuerySmBody
}
