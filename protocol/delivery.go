package protocol

// SMPP v3.4 - 5.2.17 page 124
const RegDeliveryNo = 0x00
const RegDeliverySmscBoth = 0x01 // both success and failure
const RegDeliverySmscFailed = 0x02
const RegDeliverySmeDAck = 0x04
const RegDeliverySmeUAck = 0x08
const RegDeliverySmeBoth = 0x10
const RegDeliveryIdn = 0x16 // Intermediate notification
