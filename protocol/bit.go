package protocol

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
