package pdu

type SmBody struct {
	ServiceType          string
	SourceAddrTon        uint32
	SourceAddrNpi        uint32
	SourceAddr           string
	DestAddrTon          uint32
	DestAddrNpi          uint32
	DestinationAddr      string
	EsmClass             uint32
	ProtocolId           uint32
	PriorityFlag         uint32
	ScheduleDeliveryTime string
	ValidityPeriod       string
	RegisteredDelivery   uint32
	ReplaceIfPresentFlag uint32
	DataCoding           uint32
	SmDefaultMsgId       uint32
	SmLength             uint32
	ShortMessage         string
}
