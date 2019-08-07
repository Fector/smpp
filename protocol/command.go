package protocol

// Command ids - SMPP v3.4 - 5.1.2.1 page 110-111
const GenericNack = 0x80000000
const BindReceiver = 0x00000001
const BindReceiverResp = 0x80000001
const BindTransmitter = 0x00000002
const BindTransmitterResp = 0x80000002
const QuerySm = 0x00000003
const QuerySmResp = 0x80000003
const SubmitSm = 0x00000004
const SubmitSmResp = 0x80000004
const DeliverSm = 0x00000005
const DeliverSmResp = 0x80000005
const Unbind = 0x00000006
const UnbindResp = 0x80000006
const ReplaceSm = 0x00000007
const ReplaceSmResp = 0x80000007
const CancelSm = 0x00000008
const CancelSmResp = 0x80000008
const BindTransceiver = 0x00000009
const BindTransceiverResp = 0x80000009
const Outbind = 0x0000000B
const EnquireLink = 0x00000015
const EnquireLinkResp = 0x80000015
