package pdu

type BindBody struct {
	Body
	SystemId         string
	Password         string
	SystemType       string
	InterfaceVersion uint32
	AddrTon          uint32
	AddrNpi          uint32
	AddressRange     string
}
