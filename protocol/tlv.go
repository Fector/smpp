package protocol

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
