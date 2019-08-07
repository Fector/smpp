package protocol

// Command ids - SMPP v3.4 - 5.1.2.1 page 110-111
const GenericNack uint32 = 0x80000000
const BindReceiver uint32 = 0x00000001
const BindReceiverResp uint32 = 0x80000001
const BindTransmitter uint32 = 0x00000002
const BindTransmitterResp uint32 = 0x80000002
const QuerySm uint32 = 0x00000003
const QuerySmResp uint32 = 0x80000003
const SubmitSm uint32 = 0x00000004
const SubmitSmResp uint32 = 0x80000004
const DeliverSm uint32 = 0x00000005
const DeliverSmResp uint32 = 0x80000005
const Unbind uint32 = 0x00000006
const UnbindResp uint32 = 0x80000006
const ReplaceSm uint32 = 0x00000007
const ReplaceSmResp uint32 = 0x80000007
const CancelSm uint32 = 0x00000008
const CancelSmResp uint32 = 0x80000008
const BindTransceiver uint32 = 0x00000009
const BindTransceiverResp uint32 = 0x80000009
const Outbind uint32 = 0x0000000B
const EnquireLink uint32 = 0x00000015
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

func GetCommandName(command uint32) string {
	if Command[command] == "" {
		return "UNKNOWN"
	}

	return Command[command]
}
