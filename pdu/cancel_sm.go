package pdu

type CancelSmBody struct {
	Body
	ServiceType     string
	MessageId       string
	SourceAddrTon   uint32
	SourceAddrNpi   uint32
	SourceAddr      string
	DestAddrTon     uint32
	DestAddrNpi     uint32
	DestinationAddr string
}

type CancelSm struct {
	Pdu
	Header *Header
	Body   *CancelSmBody
}
