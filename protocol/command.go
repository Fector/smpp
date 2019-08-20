package protocol

// Command ids - SMPP v3.4 - 5.1.2.1 page 110-111

// GenericNack PDU header command_id
const GenericNack uint32 = 0x80000000

// BindReceiver PDU header command_id
const BindReceiver uint32 = 0x00000001

// BindReceiverResp PDU header command_id
const BindReceiverResp uint32 = 0x80000001

// BindTransmitter PDU header command_id
const BindTransmitter uint32 = 0x00000002

// BindTransmitterResp PDU header command_id
const BindTransmitterResp uint32 = 0x80000002

// QuerySm PDU header command_id
const QuerySm uint32 = 0x00000003

// QuerySmResp PDU header command_id
const QuerySmResp uint32 = 0x80000003

// SubmitSm PDU header command_id
const SubmitSm uint32 = 0x00000004

// SubmitSmResp PDU header command_id
const SubmitSmResp uint32 = 0x80000004

// DeliverSm PDU header command_id
const DeliverSm uint32 = 0x00000005

// DeliverSmResp PDU header command_id
const DeliverSmResp uint32 = 0x80000005

// Unbind PDU header command_id
const Unbind uint32 = 0x00000006

// UnbindResp PDU header command_id
const UnbindResp uint32 = 0x80000006

// ReplaceSm PDU header command_id
const ReplaceSm uint32 = 0x00000007

// ReplaceSmResp PDU header command_id
const ReplaceSmResp uint32 = 0x80000007

// CancelSm PDU header command_id
const CancelSm uint32 = 0x00000008

// CancelSmResp PDU header command_id
const CancelSmResp uint32 = 0x80000008

// BindTransceiver PDU header command_id
const BindTransceiver uint32 = 0x00000009

// BindTransceiverResp PDU header command_id
const BindTransceiverResp uint32 = 0x80000009

// Outbind PDU header command_id
const Outbind uint32 = 0x0000000B

// EnquireLink PDU header command_id
const EnquireLink uint32 = 0x00000015

// EnquireLinkResp PDU header command_id
const EnquireLinkResp uint32 = 0x80000015

var Command = map[uint32]string{
	GenericNack:         "GENERIC_NACK",
	BindReceiver:        "BIND_RECEIVER",
	BindReceiverResp:    "BIND_RECEIVER_RESP",
	BindTransmitter:     "BIND_TRANSMITTER",
	BindTransmitterResp: "BIND_TRANSMITTER_RESP",
	QuerySm:             "QUERY_SM",
	QuerySmResp:         "QUERY_SM_RESP",
	SubmitSm:            "SUBMIT_SM",
	SubmitSmResp:        "SUBMIT_SM_RESP",
	DeliverSm:           "DELIVER_SM",
	DeliverSmResp:       "DELIVER_SM_RESP",
	Unbind:              "UNBIND",
	UnbindResp:          "UNBIND_RESP",
	ReplaceSm:           "REPLACE_SM",
	ReplaceSmResp:       "REPLACE_SM_RESP",
	CancelSm:            "CANCEL_SM",
	CancelSmResp:        "CANCEL_SM_RESP",
	BindTransceiver:     "BIND_TRANSCEIVER",
	BindTransceiverResp: "BIND_TRANSCEIVER_RESP",
	Outbind:             "OUTBIND",
	EnquireLink:         "ENQUIRE_LINK",
	EnquireLinkResp:     "ENQUIRE_LINK_RESP",
}

// GetCommandName returns name of command
func GetCommandName(command uint32) string {
	if Command[command] == "" {
		return "UNKNOWN"
	}

	return Command[command]
}
