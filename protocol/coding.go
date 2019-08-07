package protocol

// SMPP v3.4 - 5.2.19 page 126
const DataCodingDefault = 0 //UTF-8 as internal SMSC coding
const DataCodingIa5 = 1     // IA5 (CCITT T.50)/ASCII (ANSI X3.4)
const DataCodingBinaryAlias = 2
const DataCodingIso88591 = 3 // Latin 1
const DataCodingBinary = 4
const DataCodingJis = 5
const DataCodingIso88595 = 6 // Cyrllic
const DataCodingIso88598 = 7 // Latin/Hebrew
const DataCodingUcs2 = 8     // UCS-2BE (Big Endian)
const DataCodingPictogram = 9
const DataCodingIso2022Jp = 10 // Music codes
const DataCodingKanji = 13     // Extended Kanji JIS
const DataCodingKsc5601 = 14
const DataCodingUtf16be = 15
