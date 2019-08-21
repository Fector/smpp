package pdu

type DataSm struct {
	Header *Header
	Body   *SmBody
	Tlv    *[]Tlv
}
