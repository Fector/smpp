package protocol

// ESM bits 1-0 - SMPP v3.4 - 5.2.12 page 121-122
const EsmSubmitModeDatagram = 0x01
const EsmSubmitModeForward = 0x02
const EsmSubmitModeStoreandforward = 0x03

// ESM bits 5-2
const EsmSubmitDefault = 0x00
const EsmSubmitBinary = 0x04
const EsmSubmitTypeEsmeDAck = 0x08
const EsmSubmitTypeEsmeUAck = 0x10
const EsmDeliverSmscReceipt = 0x04
const EsmDeliverSmeAck = 0x08
const EsmDeliverUAck = 0x10
const EsmDeliverConvAbort = 0x18
const EsmDeliverIdn = 0x20 // Intermediate delivery notification

// ESM bits 7-6
const EsmUdhiNone = 0x00
const EsmUdhi = 0x40
const EsmUdhiDlr = 0x04
const EsmReplypath = 0x80
