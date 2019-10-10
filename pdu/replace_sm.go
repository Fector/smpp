package pdu

type ReplaceSm struct {
	Pdu
	Header *Header
	Body   *SmBody
	Tlv    *[]Tlv
}
