package protocol

import "fmt"

//  Command status - SMPP v3.4 - 5.1.3 page 112-114

// No Error
const EsmeRok uint32 = 0x00000000

// Message Length is invalid
const EsmeRinvmsglen uint32 = 0x00000001

// Command Length is invalid
const EsmeRinvcmdlen uint32 = 0x00000002

// Invalid Command ID
const EsmeRinvcmdid uint32 = 0x00000003

// Incorrect BIND Status for given command
const EsmeRinvbndsts uint32 = 0x00000004

// ESME Already in Bound State
const EsmeRalybnd uint32 = 0x00000005

// Invalid Priority Flag
const EsmeRinvprtflg uint32 = 0x00000006

// Invalid Registered Delivery Flag
const EsmeRinvregdlvflg uint32 = 0x00000007

// System Error
const EsmeRsyserr uint32 = 0x00000008

// Invalid Source Address
const EsmeRinvsrcadr uint32 = 0x0000000A

// Invalid Dest Addr
const EsmeRinvdstadr uint32 = 0x0000000B

// Message ID is invalid
const EsmeRinvmsgid uint32 = 0x0000000C

// Bind Failed
const EsmeRbindfail uint32 = 0x0000000D

// Invalid Password
const EsmeRinvpaswd uint32 = 0x0000000E

// Invalid System ID
const EsmeRinvsysid uint32 = 0x0000000F

// Cancel SM Failed
const EsmeRcancelfail uint32 = 0x00000011

// Replace SM Failed
const EsmeRreplacefail uint32 = 0x00000013

// Message Queue Full
const EsmeRmsgqful uint32 = 0x00000014

// Invalid Service Type
const EsmeRinvsertyp uint32 = 0x00000015

// Invalid number of destinations
const EsmeRinvnumdests uint32 = 0x00000033

// Invalid Distribution List name
const EsmeRinvdlname uint32 = 0x00000034

// Destination flag (submit_multi)
const EsmeRinvdestflag uint32 = 0x00000040

// Invalid ‘submit with replace’ request
const EsmeRinvsubrep uint32 = 0x00000042

// Invalid esm_SUBMIT field data
const EsmeRinvesmsubmit uint32 = 0x00000043

// Cannot Submit to Distribution List
const EsmeRcntsubdl uint32 = 0x00000044

// submit_sm or submit_multi failed
const EsmeRsubmitfail uint32 = 0x00000045

// Invalid Source address TON
const EsmeRinvsrcton uint32 = 0x00000048

// Invalid Source address NPI
const EsmeRinvsrcnpi uint32 = 0x00000049

// Invalid Destination address TON
const EsmeRinvdstton uint32 = 0x00000050

// Invalid Destination address NPI
const EsmeRinvdstnpi uint32 = 0x00000051

// Invalid system_type field
const EsmeRinvsystyp uint32 = 0x00000053

// Invalid replace_if_present flag
const EsmeRinvrepflag uint32 = 0x00000054

// Invalid number of messages
const EsmeRinvnummsgs uint32 = 0x00000055

// Throttling error (ESME has exceeded allowed message limits)
const EsmeRthrottled uint32 = 0x00000058

// Invalid Scheduled Delivery Time
const EsmeRinvsched uint32 = 0x00000061

// Invalid message (Expiry time)
const EsmeRinvexpiry uint32 = 0x00000062

// Predefined Message Invalid or Not Found
const EsmeRinvdftmsgid uint32 = 0x00000063

// ESME Receiver Temporary App Error Code
const EsmeRxTAppn uint32 = 0x00000064

// ESME Receiver Permanent App Error Code
const EsmeRxPAppn uint32 = 0x00000065

// ESME Receiver Reject Message Error Code
const EsmeRxRAppn uint32 = 0x00000066

// query_sm request failed
const EsmeRqueryfail uint32 = 0x00000067

// Error in the optional part of the PDU Body
const EsmeRinvoptparstream uint32 = 0x000000C0

// Optional Parameter not allowed
const EsmeRoptparnotallwd uint32 = 0x000000C1

// Invalid Parameter Length
const EsmeRinvparlen uint32 = 0x000000C2

// Expected Optional Parameter missing
const EsmeRmissingoptparam uint32 = 0x000000C3

// Invalid Optional Parameter Value
const EsmeRinvoptparamval uint32 = 0x000000C4

// Delivery Failure (data_sm_resp)
const EsmeRdeliveryfailure uint32 = 0x000000FE

// Unknown Error
const EsmeRunknownerr uint32 = 0x000000FF

// Status names map
var Status = map[uint32]string{
	EsmeRok:              "No Error",
	EsmeRinvmsglen:       "Message Length is invalid",
	EsmeRinvcmdlen:       "Command Length is invalid",
	EsmeRinvcmdid:        "Invalid Command ID",
	EsmeRinvbndsts:       "Incorrect BIND Status for given command",
	EsmeRalybnd:          "ESME Already in Bound State",
	EsmeRinvprtflg:       "Invalid Priority Flag",
	EsmeRinvregdlvflg:    "Invalid Registered Delivery Flag",
	EsmeRsyserr:          "System Error",
	EsmeRinvsrcadr:       "Invalid Source Address",
	EsmeRinvdstadr:       "Invalid Dest Addr",
	EsmeRinvmsgid:        "Message ID is invalid",
	EsmeRbindfail:        "Bind Failed",
	EsmeRinvpaswd:        "Invalid Password",
	EsmeRinvsysid:        "Invalid System ID",
	EsmeRcancelfail:      "Cancel SM Failed",
	EsmeRreplacefail:     "Replace SM Failed",
	EsmeRmsgqful:         "Message Queue Full",
	EsmeRinvsertyp:       "Invalid Service Type",
	EsmeRinvnumdests:     "Invalid number of destinations",
	EsmeRinvdlname:       "Invalid Distribution List name",
	EsmeRinvdestflag:     "Destination flag (submit_multi)",
	EsmeRinvsubrep:       "Invalid ‘submit with replace’ request",
	EsmeRinvesmsubmit:    "Invalid esm_SUBMIT field data",
	EsmeRcntsubdl:        "Cannot Submit to Distribution List",
	EsmeRsubmitfail:      "submit_sm or submit_multi failed",
	EsmeRinvsrcton:       "Invalid Source address TON",
	EsmeRinvsrcnpi:       "Invalid Source address NPI",
	EsmeRinvdstton:       "Invalid Destination address TON",
	EsmeRinvdstnpi:       "Invalid Destination address NPI",
	EsmeRinvsystyp:       "Invalid system_type field",
	EsmeRinvrepflag:      "Invalid replace_if_present flag",
	EsmeRinvnummsgs:      "Invalid number of messages",
	EsmeRthrottled:       "Throttling error (ESME has exceeded allowed message limits)",
	EsmeRinvsched:        "Invalid Scheduled Delivery Time",
	EsmeRinvexpiry:       "Invalid message (Expiry time)",
	EsmeRinvdftmsgid:     "Predefined Message Invalid or Not Found",
	EsmeRxTAppn:          "ESME Receiver Temporary App Error Code",
	EsmeRxPAppn:          "ESME Receiver Permanent App Error Code",
	EsmeRxRAppn:          "ESME Receiver Reject Message Error Code",
	EsmeRqueryfail:       "query_sm request failed",
	EsmeRinvoptparstream: "Error in the optional part of the PDU Body.",
	EsmeRoptparnotallwd:  "Optional Parameter not allowed",
	EsmeRinvparlen:       "Invalid Parameter Length.",
	EsmeRmissingoptparam: "Expected Optional Parameter missing",
	EsmeRinvoptparamval:  "Invalid Optional Parameter Value",
	EsmeRdeliveryfailure: "Delivery Failure (data_sm_resp)",
	EsmeRunknownerr:      "Unknown Error",
}

// Return name of status
func GetStatusName(status uint32) string {
	if Status[status] == "" {
		return fmt.Sprintf("Specific status code %x", status)
	}

	return Status[status]
}
