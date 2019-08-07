package protocol

// SMPP v3.4 - 5.2.18 page 125
const ReplaceNo uint32 = 0x00
const ReplaceYes uint32 = 0x01

// SMPP v3.4 - 5.2.21 page 128
const NoUserDataSm uint32 = 0x00

// SMPP v3.4 - 5.2.25 page 129
const DestFlagSme uint32 = 1
const DestFlagDistlist uint32 = 2
