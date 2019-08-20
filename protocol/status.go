package protocol

import "fmt"

//  Command status - SMPP v3.4 - 5.1.3 page 112-114

// EsmeRok is "No Error" code
const EsmeRok uint32 = 0x00000000

// EsmeRinvmsglen is a "Message Length is invalid" error code
const EsmeRinvmsglen uint32 = 0x00000001

// EsmeRinvcmdlen is a "Command Length is invalid" error code
const EsmeRinvcmdlen uint32 = 0x00000002

// EsmeRinvcmdid is a "Invalid Command ID" error code
const EsmeRinvcmdid uint32 = 0x00000003

// EsmeRinvbndsts is a "Incorrect BIND Status for given command" error code
const EsmeRinvbndsts uint32 = 0x00000004

// EsmeRalybnd is a "ESME Already in Bound State" error code
const EsmeRalybnd uint32 = 0x00000005

// EsmeRinvprtflg is a "Invalid Priority Flag" error code
const EsmeRinvprtflg uint32 = 0x00000006

// EsmeRinvregdlvflg is a "Invalid Registered Delivery Flag" error code
const EsmeRinvregdlvflg uint32 = 0x00000007

// EsmeRsyserr is a "System Error" code
const EsmeRsyserr uint32 = 0x00000008

// EsmeRinvsrcadr is a "Invalid Source Address" error code
const EsmeRinvsrcadr uint32 = 0x0000000A

// EsmeRinvdstadr is a "Invalid Dest Addr" error code
const EsmeRinvdstadr uint32 = 0x0000000B

// EsmeRinvmsgid is a "Message ID is invalid" error code
const EsmeRinvmsgid uint32 = 0x0000000C

// EsmeRbindfail is a "Bind Failed" error code
const EsmeRbindfail uint32 = 0x0000000D

// EsmeRinvpaswd is a "Invalid Password" error code
const EsmeRinvpaswd uint32 = 0x0000000E

// EsmeRinvsysid is a "Invalid System ID" error code
const EsmeRinvsysid uint32 = 0x0000000F

// EsmeRcancelfail is a "Cancel SM Failed" error code
const EsmeRcancelfail uint32 = 0x00000011

// EsmeRreplacefail is a "Replace SM Failed" error code
const EsmeRreplacefail uint32 = 0x00000013

// EsmeRmsgqful is a "Message Queue Full" error code
const EsmeRmsgqful uint32 = 0x00000014

// EsmeRinvsertyp is a "Invalid Service Type" error code
const EsmeRinvsertyp uint32 = 0x00000015

// EsmeRinvnumdests is a "Invalid number of destinations" error code
const EsmeRinvnumdests uint32 = 0x00000033

// EsmeRinvdlname is a "Invalid Distribution List name" error code
const EsmeRinvdlname uint32 = 0x00000034

// EsmeRinvdestflag is a "Destination flag (submit_multi)" error code
const EsmeRinvdestflag uint32 = 0x00000040

// EsmeRinvsubrep is a "Invalid ‘submit with replace’ request" error code
const EsmeRinvsubrep uint32 = 0x00000042

// EsmeRinvesmsubmit is a "Invalid esm_SUBMIT field data" error code
const EsmeRinvesmsubmit uint32 = 0x00000043

// EsmeRcntsubdl is a "Cannot Submit to Distribution List" error code
const EsmeRcntsubdl uint32 = 0x00000044

// EsmeRsubmitfail is a "submit_sm or submit_multi failed" error code
const EsmeRsubmitfail uint32 = 0x00000045

// EsmeRinvsrcton is a "Invalid Source address TON" error code
const EsmeRinvsrcton uint32 = 0x00000048

// EsmeRinvsrcnpi is a "Invalid Source address NPI" error code
const EsmeRinvsrcnpi uint32 = 0x00000049

// EsmeRinvdstton is a "Invalid Destination address TON" error code
const EsmeRinvdstton uint32 = 0x00000050

// EsmeRinvdstnpi is a "Invalid Destination address NPI" error code
const EsmeRinvdstnpi uint32 = 0x00000051

// EsmeRinvsystyp is a "Invalid system_type field" error code
const EsmeRinvsystyp uint32 = 0x00000053

// EsmeRinvrepflag is a "Invalid replace_if_present flag" error code
const EsmeRinvrepflag uint32 = 0x00000054

// EsmeRinvnummsgs is a "Invalid number of messages" error code
const EsmeRinvnummsgs uint32 = 0x00000055

// EsmeRthrottled is a "Throttling error (ESME has exceeded allowed message limits)" error code
const EsmeRthrottled uint32 = 0x00000058

// EsmeRinvsched is a "Invalid Scheduled Delivery Time" error code
const EsmeRinvsched uint32 = 0x00000061

// EsmeRinvexpiry is a "Invalid message (Expiry time)" error code
const EsmeRinvexpiry uint32 = 0x00000062

// EsmeRinvdftmsgid is a "Predefined Message Invalid or Not Found" error code
const EsmeRinvdftmsgid uint32 = 0x00000063

// EsmeRxTAppn is a "ESME Receiver Temporary App Error Code"
const EsmeRxTAppn uint32 = 0x00000064

// EsmeRxPAppn is a "ESME Receiver Permanent App Error Code"
const EsmeRxPAppn uint32 = 0x00000065

// EsmeRxRAppn is a "ESME Receiver Reject Message Error Code"
const EsmeRxRAppn uint32 = 0x00000066

// EsmeRqueryfail is a "query_sm request failed" error code
const EsmeRqueryfail uint32 = 0x00000067

// EsmeRinvoptparstream is a "Error in the optional part of the PDU Body"
const EsmeRinvoptparstream uint32 = 0x000000C0

// EsmeRoptparnotallwd is a "Optional Parameter not allowed" error code
const EsmeRoptparnotallwd uint32 = 0x000000C1

// EsmeRinvparlen is a "Invalid Parameter Length" error code
const EsmeRinvparlen uint32 = 0x000000C2

// EsmeRmissingoptparam is a "Expected Optional Parameter missing" error code
const EsmeRmissingoptparam uint32 = 0x000000C3

// EsmeRinvoptparamval is a "Invalid Optional Parameter Value" error code
const EsmeRinvoptparamval uint32 = 0x000000C4

// EsmeRdeliveryfailure is a "Delivery Failure (data_sm_resp)" error code
const EsmeRdeliveryfailure uint32 = 0x000000FE

// EsmeRunknownerr is a "Unknown Error" code
const EsmeRunknownerr uint32 = 0x000000FF

// Status is status names map
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

// GetStatusName is a function to get command status by id
func GetStatusName(status uint32) string {
	if Status[status] == "" {
		return fmt.Sprintf("Specific status code %x", status)
	}

	return Status[status]
}
