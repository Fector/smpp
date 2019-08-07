package pdu

type Header struct {
	CommandLength  uint
	CommandId      uint
	CommandStatus  uint
	SequenceNumber uint
}
