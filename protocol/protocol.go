package protocol

// SMPP v3.4 - 3.2 page 38
const PduHeaderLength uint32 = 16

// SMPP v3.4 - 5.2.13 page 123
const ProtocolId uint32 = 0x34

// SMPP v3.4 bind mode short names
const BindModeTX string = "TX"
const BindModeRX string = "RX"
const BindModeTRX string = "TRX"
