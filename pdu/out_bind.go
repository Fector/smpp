package pdu

type OutBindBody struct {
	SystemId string
	Password string
}

type OutBind struct {
	Pdu
	Header *Header
	Body   *OutBindBody
}
