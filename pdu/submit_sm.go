package pdu

type SubmitSm struct {
	Header *Header
	Body   *SmBody
	Tlv    *[]Tlv
}
