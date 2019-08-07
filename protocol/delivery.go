package protocol

// SMPP v3.4 - 5.2.17 page 124
const RegDeliveryNo uint32 = 0x00
const RegDeliverySmscBoth uint32 = 0x01 // both success and failure
const RegDeliverySmscFailed uint32 = 0x02
const RegDeliverySmeDAck uint32 = 0x04
const RegDeliverySmeUAck uint32 = 0x08
const RegDeliverySmeBoth uint32 = 0x10
const RegDeliveryIdn uint32 = 0x16 // Intermediate notification
