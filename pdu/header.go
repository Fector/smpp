package pdu

type Header struct {
	CommandLength  uint32
	CommandId      uint32
	CommandStatus  uint32
	SequenceNumber uint32
}
