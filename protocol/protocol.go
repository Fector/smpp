package protocol

import "fmt"

// SMPP v3.4 - 3.2 page 38
const PduHeaderLength uint32 = 16

// SMPP v3.4 - 5.2.13 page 123
const ProtocolId uint32 = 0x34

// SMPP v3.4 bind mode short names
const BindModeTX string = "TX"
const BindModeRX string = "RX"
const BindModeTRX string = "TRX"

// Command ids - SMPP v3.4 - 5.1.2.1 page 110-111

// GenericNack is a GenericNack PDU header command_id
const GenericNack uint32 = 0x80000000

// BindReceiver is a BindReceiver PDU header command_id
const BindReceiver uint32 = 0x00000001

// BindReceiverResp is a BindReceiverResp PDU header command_id
const BindReceiverResp uint32 = 0x80000001

// BindTransmitter is a BindTransmitter PDU header command_id
const BindTransmitter uint32 = 0x00000002

// BindTransmitterResp is a BindTransmitterResp PDU header command_id
const BindTransmitterResp uint32 = 0x80000002

// QuerySm is a QuerySm PDU header command_id
const QuerySm uint32 = 0x00000003

// QuerySmResp is a QuerySmResp PDU header command_id
const QuerySmResp uint32 = 0x80000003

// SubmitSm is a SubmitSm PDU header command_id
const SubmitSm uint32 = 0x00000004

// SubmitSmResp is a SubmitSmResp PDU header command_id
const SubmitSmResp uint32 = 0x80000004

// DeliverSm is a DeliverSm PDU header command_id
const DeliverSm uint32 = 0x00000005

// DeliverSmResp is a DeliverSmResp PDU header command_id
const DeliverSmResp uint32 = 0x80000005

// Unbind is a unbind PDU header command_id
const Unbind uint32 = 0x00000006

// UnbindResp PDU header command_id
const UnbindResp uint32 = 0x80000006

// ReplaceSm is a ReplaceSm PDU header command_id
const ReplaceSm uint32 = 0x00000007

// ReplaceSmResp is a ReplaceSmResp PDU header command_id
const ReplaceSmResp uint32 = 0x80000007

// CancelSm is a CancelSm PDU header command_id
const CancelSm uint32 = 0x00000008

// CancelSmResp is a CancelSmResp PDU header command_id
const CancelSmResp uint32 = 0x80000008

// BindTransceiver is a BindTransceiver PDU header command_id
const BindTransceiver uint32 = 0x00000009

// BindTransceiverResp is a BindTransceiverResp PDU header command_id
const BindTransceiverResp uint32 = 0x80000009

// Outbind is a Outbind PDU header command_id
const Outbind uint32 = 0x0000000B

// EnquireLink is a EnquireLink PDU header command_id
const EnquireLink uint32 = 0x00000015

// EnquireLinkResp is a EnquireLinkResp PDU header command_id
const EnquireLinkResp uint32 = 0x80000015

// Command is command names map
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

// GetCommandName is a function to get command name by id
func GetCommandName(command uint32) string {
	if Command[command] == "" {
		return "UNKNOWN"
	}

	return Command[command]
}

//  Command status - SMPP v3.4 - 5.1.3 page 112-114

// EsmeRok is "No Error" code
const EsmeRok uint32 = 0x00000000

// EsmeRinvmsglen is a "Message Length is invalid" error code
const EsmeRinvmsglen uint32 = 0x00000001

// EsmeRinvcmdlen is a "Command Length is invalid" error code
const EsmeRinvcmdlen uint32 = 0x00000002

// EsmeRinvcmdid is a "Invalid Command ID" error code
const EsmeRinvcmdid uint32 = 0x00000003

// EsmeRinvbndsts is a "Incorrect BIND Status for given command" error code
const EsmeRinvbndsts uint32 = 0x00000004

// EsmeRalybnd is a "ESME Already in Bound State" error code
const EsmeRalybnd uint32 = 0x00000005

// EsmeRinvprtflg is a "Invalid Priority Flag" error code
const EsmeRinvprtflg uint32 = 0x00000006

// EsmeRinvregdlvflg is a "Invalid Registered Delivery Flag" error code
const EsmeRinvregdlvflg uint32 = 0x00000007

// EsmeRsyserr is a "System Error" code
const EsmeRsyserr uint32 = 0x00000008

// EsmeRinvsrcadr is a "Invalid Source Address" error code
const EsmeRinvsrcadr uint32 = 0x0000000A

// EsmeRinvdstadr is a "Invalid Dest Addr" error code
const EsmeRinvdstadr uint32 = 0x0000000B

// EsmeRinvmsgid is a "Message ID is invalid" error code
const EsmeRinvmsgid uint32 = 0x0000000C

// EsmeRbindfail is a "Bind Failed" error code
const EsmeRbindfail uint32 = 0x0000000D

// EsmeRinvpaswd is a "Invalid Password" error code
const EsmeRinvpaswd uint32 = 0x0000000E

// EsmeRinvsysid is a "Invalid System ID" error code
const EsmeRinvsysid uint32 = 0x0000000F

// EsmeRcancelfail is a "Cancel SM Failed" error code
const EsmeRcancelfail uint32 = 0x00000011

// EsmeRreplacefail is a "Replace SM Failed" error code
const EsmeRreplacefail uint32 = 0x00000013

// EsmeRmsgqful is a "Message Queue Full" error code
const EsmeRmsgqful uint32 = 0x00000014

// EsmeRinvsertyp is a "Invalid Service Type" error code
const EsmeRinvsertyp uint32 = 0x00000015

// EsmeRinvnumdests is a "Invalid number of destinations" error code
const EsmeRinvnumdests uint32 = 0x00000033

// EsmeRinvdlname is a "Invalid Distribution List name" error code
const EsmeRinvdlname uint32 = 0x00000034

// EsmeRinvdestflag is a "Destination flag (submit_multi)" error code
const EsmeRinvdestflag uint32 = 0x00000040

// EsmeRinvsubrep is a "Invalid ‘submit with replace’ request" error code
const EsmeRinvsubrep uint32 = 0x00000042

// EsmeRinvesmsubmit is a "Invalid esm_SUBMIT field data" error code
const EsmeRinvesmsubmit uint32 = 0x00000043

// EsmeRcntsubdl is a "Cannot Submit to Distribution List" error code
const EsmeRcntsubdl uint32 = 0x00000044

// EsmeRsubmitfail is a "submit_sm or submit_multi failed" error code
const EsmeRsubmitfail uint32 = 0x00000045

// EsmeRinvsrcton is a "Invalid Source address TON" error code
const EsmeRinvsrcton uint32 = 0x00000048

// EsmeRinvsrcnpi is a "Invalid Source address NPI" error code
const EsmeRinvsrcnpi uint32 = 0x00000049

// EsmeRinvdstton is a "Invalid Destination address TON" error code
const EsmeRinvdstton uint32 = 0x00000050

// EsmeRinvdstnpi is a "Invalid Destination address NPI" error code
const EsmeRinvdstnpi uint32 = 0x00000051

// EsmeRinvsystyp is a "Invalid system_type field" error code
const EsmeRinvsystyp uint32 = 0x00000053

// EsmeRinvrepflag is a "Invalid replace_if_present flag" error code
const EsmeRinvrepflag uint32 = 0x00000054

// EsmeRinvnummsgs is a "Invalid number of messages" error code
const EsmeRinvnummsgs uint32 = 0x00000055

// EsmeRthrottled is a "Throttling error (ESME has exceeded allowed message limits)" error code
const EsmeRthrottled uint32 = 0x00000058

// EsmeRinvsched is a "Invalid Scheduled Delivery Time" error code
const EsmeRinvsched uint32 = 0x00000061

// EsmeRinvexpiry is a "Invalid message (Expiry time)" error code
const EsmeRinvexpiry uint32 = 0x00000062

// EsmeRinvdftmsgid is a "Predefined Message Invalid or Not Found" error code
const EsmeRinvdftmsgid uint32 = 0x00000063

// EsmeRxTAppn is a "ESME Receiver Temporary App Error Code"
const EsmeRxTAppn uint32 = 0x00000064

// EsmeRxPAppn is a "ESME Receiver Permanent App Error Code"
const EsmeRxPAppn uint32 = 0x00000065

// EsmeRxRAppn is a "ESME Receiver Reject Message Error Code"
const EsmeRxRAppn uint32 = 0x00000066

// EsmeRqueryfail is a "query_sm request failed" error code
const EsmeRqueryfail uint32 = 0x00000067

// EsmeRinvoptparstream is a "Error in the optional part of the PDU Body"
const EsmeRinvoptparstream uint32 = 0x000000C0

// EsmeRoptparnotallwd is a "Optional Parameter not allowed" error code
const EsmeRoptparnotallwd uint32 = 0x000000C1

// EsmeRinvparlen is a "Invalid Parameter Length" error code
const EsmeRinvparlen uint32 = 0x000000C2

// EsmeRmissingoptparam is a "Expected Optional Parameter missing" error code
const EsmeRmissingoptparam uint32 = 0x000000C3

// EsmeRinvoptparamval is a "Invalid Optional Parameter Value" error code
const EsmeRinvoptparamval uint32 = 0x000000C4

// EsmeRdeliveryfailure is a "Delivery Failure (data_sm_resp)" error code
const EsmeRdeliveryfailure uint32 = 0x000000FE

// EsmeRunknownerr is a "Unknown Error" code
const EsmeRunknownerr uint32 = 0x000000FF

// Status is status names map
var Status = map[uint32]string{
	EsmeRok:              "No Error",
	EsmeRinvmsglen:       "Message Length is invalid",
	EsmeRinvcmdlen:       "Command Length is invalid",
	EsmeRinvcmdid:        "Invalid Command ID",
	EsmeRinvbndsts:       "Incorrect BIND Status for given command",
	EsmeRalybnd:          "ESME Already in Bound State",
	EsmeRinvprtflg:       "Invalid Priority Flag",
	EsmeRinvregdlvflg:    "Invalid Registered Delivery Flag",
	EsmeRsyserr:          "System Error",
	EsmeRinvsrcadr:       "Invalid Source Address",
	EsmeRinvdstadr:       "Invalid Dest Addr",
	EsmeRinvmsgid:        "Message ID is invalid",
	EsmeRbindfail:        "Bind Failed",
	EsmeRinvpaswd:        "Invalid Password",
	EsmeRinvsysid:        "Invalid System ID",
	EsmeRcancelfail:      "Cancel SM Failed",
	EsmeRreplacefail:     "Replace SM Failed",
	EsmeRmsgqful:         "Message Queue Full",
	EsmeRinvsertyp:       "Invalid Service Type",
	EsmeRinvnumdests:     "Invalid number of destinations",
	EsmeRinvdlname:       "Invalid Distribution List name",
	EsmeRinvdestflag:     "Destination flag (submit_multi)",
	EsmeRinvsubrep:       "Invalid ‘submit with replace’ request",
	EsmeRinvesmsubmit:    "Invalid esm_SUBMIT field data",
	EsmeRcntsubdl:        "Cannot Submit to Distribution List",
	EsmeRsubmitfail:      "submit_sm or submit_multi failed",
	EsmeRinvsrcton:       "Invalid Source address TON",
	EsmeRinvsrcnpi:       "Invalid Source address NPI",
	EsmeRinvdstton:       "Invalid Destination address TON",
	EsmeRinvdstnpi:       "Invalid Destination address NPI",
	EsmeRinvsystyp:       "Invalid system_type field",
	EsmeRinvrepflag:      "Invalid replace_if_present flag",
	EsmeRinvnummsgs:      "Invalid number of messages",
	EsmeRthrottled:       "Throttling error (ESME has exceeded allowed message limits)",
	EsmeRinvsched:        "Invalid Scheduled Delivery Time",
	EsmeRinvexpiry:       "Invalid message (Expiry time)",
	EsmeRinvdftmsgid:     "Predefined Message Invalid or Not Found",
	EsmeRxTAppn:          "ESME Receiver Temporary App Error Code",
	EsmeRxPAppn:          "ESME Receiver Permanent App Error Code",
	EsmeRxRAppn:          "ESME Receiver Reject Message Error Code",
	EsmeRqueryfail:       "query_sm request failed",
	EsmeRinvoptparstream: "Error in the optional part of the PDU Body.",
	EsmeRoptparnotallwd:  "Optional Parameter not allowed",
	EsmeRinvparlen:       "Invalid Parameter Length.",
	EsmeRmissingoptparam: "Expected Optional Parameter missing",
	EsmeRinvoptparamval:  "Invalid Optional Parameter Value",
	EsmeRdeliveryfailure: "Delivery Failure (data_sm_resp)",
	EsmeRunknownerr:      "Unknown Error",
}

// GetStatusName is a function to get command status by id
func GetStatusName(status uint32) string {
	if Status[status] == "" {
		return fmt.Sprintf("Specific status code %x", status)
	}

	return Status[status]
}

// SMPP v3.4 - 5.2.5 page 117
const TonUnknown = 0x00
const TonInternational = 0x01
const TonNational = 0x02
const TonNetworkspecific = 0x03
const TonSubscribernumber = 0x04
const TonAlphanumeric = 0x05
const TonAbbreviated = 0x06

// SMPP v3.4 - 5.2.6 page 118
const NpiUnknown uint32 = 0x00
const NpiE164 uint32 = 0x01
const NpiData uint32 = 0x03
const NpiTelex uint32 = 0x04
const NpiE212 uint32 = 0x06
const NpiNational uint32 = 0x08
const NpiPrivate uint32 = 0x09
const NpiErmes uint32 = 0x0a
const NpiInternet uint32 = 0x0e
const NpiWapclient uint32 = 0x12

// SMPP v3.4 - 5.2.18 page 125
const ReplaceNo uint32 = 0x00
const ReplaceYes uint32 = 0x01

// SMPP v3.4 - 5.2.21 page 128
const NoUserDataSm uint32 = 0x00

// SMPP v3.4 - 5.2.25 page 129
const DestFlagSme uint32 = 1
const DestFlagDistlist uint32 = 2

// ESM bits 1-0 - SMPP v3.4 - 5.2.12 page 121-122
const EsmSubmitModeDatagram uint32 = 0x01
const EsmSubmitModeForward uint32 = 0x02
const EsmSubmitModeStoreandforward uint32 = 0x03

// ESM bits 5-2
const EsmSubmitDefault uint32 = 0x00
const EsmSubmitBinary uint32 = 0x04
const EsmSubmitTypeEsmeDAck uint32 = 0x08
const EsmSubmitTypeEsmeUAck uint32 = 0x10
const EsmDeliverSmscReceipt uint32 = 0x04
const EsmDeliverSmeAck uint32 = 0x08
const EsmDeliverUAck uint32 = 0x10
const EsmDeliverConvAbort uint32 = 0x18
const EsmDeliverIdn uint32 = 0x20 // Intermediate delivery notification

// ESM bits 7-6
const EsmUdhiNone uint32 = 0x00
const EsmUdhi uint32 = 0x40
const EsmUdhiDlr uint32 = 0x04
const EsmReplypath uint32 = 0x80

// SMPP v3.4 - 5.2.19 page 126
const DataCodingDefault uint32 = 0 //UTF-8 as internal SMSC coding
const DataCodingIa5 uint32 = 1     // IA5 (CCITT T.50)/ASCII (ANSI X3.4)
const DataCodingBinaryAlias uint32 = 2
const DataCodingIso88591 uint32 = 3 // Latin 1
const DataCodingBinary uint32 = 4
const DataCodingJis uint32 = 5
const DataCodingIso88595 uint32 = 6 // Cyrllic
const DataCodingIso88598 uint32 = 7 // Latin/Hebrew
const DataCodingUcs2 uint32 = 8     // UCS-2BE (Big Endian)
const DataCodingPictogram uint32 = 9
const DataCodingIso2022Jp uint32 = 10 // Music codes
const DataCodingKanji uint32 = 13     // Extended Kanji JIS
const DataCodingKsc5601 uint32 = 14
const DataCodingUtf16be uint32 = 15

// SMPP v3.4 - 5.2.14 page 123
const PriorityFlag0 uint32 = 0x00
const PriorityFlag1 uint32 = 0x01
const PriorityFlag2 uint32 = 0x02
const PriorityFlag3 uint32 = 0x03

// SMPP v3.4 - 5.2.17 page 124
const RegDeliveryNo uint32 = 0x00
const RegDeliverySmscBoth uint32 = 0x01 // both success and failure
const RegDeliverySmscFailed uint32 = 0x02
const RegDeliverySmeDAck uint32 = 0x04
const RegDeliverySmeUAck uint32 = 0x08
const RegDeliverySmeBoth uint32 = 0x10
const RegDeliveryIdn uint32 = 0x16 // Intermediate notification

// SMPP v3.4 - 5.2.28 page 130
const StateEnroute uint32 = 1
const StateDelivered uint32 = 2
const StateExpired uint32 = 3
const StateDeleted uint32 = 4
const StateUndeliverable uint32 = 5
const StateAccepted uint32 = 6
const StateUnknown uint32 = 7
const StateRejected uint32 = 8

var State = map[uint32]string{
	StateEnroute:       "ENROUTE",
	StateDelivered:     "DELIVRD",
	StateExpired:       "EXPIRED",
	StateDeleted:       "DELETED",
	StateUndeliverable: "UNDELIV",
	StateAccepted:      "ACCEPTD",
	StateUnknown:       "UNKNOWN",
	StateRejected:      "REJECTD",
}

func GetStateName(state uint32) string {
	if State[state] == "" {
		return State[StateUnknown]
	}

	return State[state]
}

// SMPP v3.4 - 5.2.28 page 132
const DestAddrSubunitTlv uint32 = 0x0005
const DestNetworkTypeTlv uint32 = 0x0006
const DestBearerTypeTlv uint32 = 0x0007
const DestTelematicsIdTlv uint32 = 0x0008
const SourceAddrSubunitTlv uint32 = 0x000D
const SourceNetworkTypeTlv uint32 = 0x000E
const SourceBearerTypeTlv uint32 = 0x000F
const SourceTelematicsIdTlv uint32 = 0x0010
const QosTimeToLiveTlv uint32 = 0x0017
const PayloadTypeTlv uint32 = 0x0019
const AdditionalStatusInfoTextTlv uint32 = 0x001D
const ReceiptedMessageIdTlv uint32 = 0x001E
const MsMsgWaitFacilitiesTlv uint32 = 0x0030
const PrivacyIndicatorTlv uint32 = 0x0201
const SourceSubaddressTlv uint32 = 0x0202
const DestSubaddressTlv uint32 = 0x0203
const UserMessageReferenceTlv uint32 = 0x0204
const UserResponseCodeTlv uint32 = 0x0205
const SourcePortTlv uint32 = 0x020A
const DestinationPortTlv uint32 = 0x020B
const SarMsgRefNumTlv uint32 = 0x020C
const LanguageIndicatorTlv uint32 = 0x020D
const SarTotalSegmentsTlv uint32 = 0x020E
const SarSegmentSeqnumTlv uint32 = 0x020F
const ScInterfaceVersionTlv uint32 = 0x0210
const CallbackNumPresIndTlv uint32 = 0x0302
const CallbackNumAtagTlv uint32 = 0x0303
const NumberOfMessagesTlv uint32 = 0x0304
const CallbackNumTlv uint32 = 0x0381
const DpfResultTlv uint32 = 0x0420
const SetDpfTlv uint32 = 0x0421
const MsAvailabilityStatusTlv uint32 = 0x0422
const NetworkErrorCodeTlv uint32 = 0x0423
const MessagePayloadTlv uint32 = 0x0424
const DeliveryFailureReasonTlv uint32 = 0x0425
const MoreMessagesToSendTlv uint32 = 0x0426
const MessageStateTlv uint32 = 0x0427
const UssdServiceOpTlv uint32 = 0x0501
const DisplayTimeTlv uint32 = 0x1201
const SmsSignalTlv uint32 = 0x1203
const MsValidityTlv uint32 = 0x1204
const AlertOnMessageDeliveryTlv uint32 = 0x130C
const ItsReplyTypeTlv uint32 = 0x1380
const ItsSessionInfoTlv uint32 = 0x1383

var Tlv = map[uint32]string{
	DestAddrSubunitTlv:          "dest_addr_subunit",
	DestNetworkTypeTlv:          "dest_network_type",
	DestBearerTypeTlv:           "dest_bearer_type",
	DestTelematicsIdTlv:         "dest_telematics_id",
	SourceAddrSubunitTlv:        "source_addr_subunit",
	SourceNetworkTypeTlv:        "source_network_type",
	SourceBearerTypeTlv:         "source_bearer_type",
	SourceTelematicsIdTlv:       "source_telematics_id",
	QosTimeToLiveTlv:            "qos_time_to_live",
	PayloadTypeTlv:              "payload_type",
	AdditionalStatusInfoTextTlv: "additional_status_info_text",
	ReceiptedMessageIdTlv:       "receipted_message_id",
	MsMsgWaitFacilitiesTlv:      "ms_msg_wait_facilities",
	PrivacyIndicatorTlv:         "privacy_indicator",
	SourceSubaddressTlv:         "source_subaddress",
	DestSubaddressTlv:           "dest_subaddress",
	UserMessageReferenceTlv:     "user_message_reference",
	UserResponseCodeTlv:         "user_response_code",
	SourcePortTlv:               "source_port",
	DestinationPortTlv:          "destination_port",
	SarMsgRefNumTlv:             "sar_msg_ref_num",
	LanguageIndicatorTlv:        "language_indicator",
	SarTotalSegmentsTlv:         "sar_total_segments",
	SarSegmentSeqnumTlv:         "sar_segment_seqnum",
	ScInterfaceVersionTlv:       "sc_interface_version",
	CallbackNumPresIndTlv:       "callback_num_pres_ind",
	CallbackNumAtagTlv:          "callback_num_atag",
	NumberOfMessagesTlv:         "number_of_messages",
	CallbackNumTlv:              "callback_num",
	DpfResultTlv:                "dpf_result",
	SetDpfTlv:                   "set_dpf",
	MsAvailabilityStatusTlv:     "ms_availability_status",
	NetworkErrorCodeTlv:         "network_error_code",
	MessagePayloadTlv:           "message_payload",
	DeliveryFailureReasonTlv:    "delivery_failure_reason",
	MoreMessagesToSendTlv:       "more_messages_to_send",
	MessageStateTlv:             "message_state",
	UssdServiceOpTlv:            "ussd_service_op",
	DisplayTimeTlv:              "display_time",
	SmsSignalTlv:                "sms_signal",
	MsValidityTlv:               "ms_validity",
	AlertOnMessageDeliveryTlv:   "alert_on_message_delivery",
	ItsReplyTypeTlv:             "its_reply_type",
	ItsSessionInfoTlv:           "its_session_info",
}

func GetTlvName(tag uint32) string {
	if Tlv[tag] == "" {
		return "UNKNOWN"
	}

	return Tlv[tag]
}
