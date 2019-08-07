package protocol

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
