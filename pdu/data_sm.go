package pdu

type DataSmBody struct {
	ServiceType        string
	SourceAddrTon      uint32
	SourceAddrNpi      uint32
	SourceAddr         string
	DestAddrTon        uint32
	DestAddrNpi        uint32
	DestinationAddr    string
	EsmClass           uint32
	RegisteredDelivery uint32
	DataCoding         uint32
}

type DataSm struct {
	Header *Header
	Body   *DataSmBody
	Tlv    *[]Tlv
}
