package protocol

import "fmt"

//  Command status - SMPP v3.4 - 5.1.3 page 112-114
const EsmeRok uint32 = 0x00000000              // No Error
const EsmeRinvmsglen uint32 = 0x00000001       // Message Length is invalid
const EsmeRinvcmdlen uint32 = 0x00000002       // Command Length is invalid
const EsmeRinvcmdid uint32 = 0x00000003        // Invalid Command ID
const EsmeRinvbndsts uint32 = 0x00000004       // Incorrect BIND Status for given command
const EsmeRalybnd uint32 = 0x00000005          // ESME Already in Bound State
const EsmeRinvprtflg uint32 = 0x00000006       // Invalid Priority Flag
const EsmeRinvregdlvflg uint32 = 0x00000007    // Invalid Registered Delivery Flag
const EsmeRsyserr uint32 = 0x00000008          // System Error
const EsmeRinvsrcadr uint32 = 0x0000000A       // Invalid Source Address
const EsmeRinvdstadr uint32 = 0x0000000B       // Invalid Dest Addr
const EsmeRinvmsgid uint32 = 0x0000000C        // Message ID is invalid
const EsmeRbindfail uint32 = 0x0000000D        // Bind Failed
const EsmeRinvpaswd uint32 = 0x0000000E        // Invalid Password
const EsmeRinvsysid uint32 = 0x0000000F        // Invalid System ID
const EsmeRcancelfail uint32 = 0x00000011      // Cancel SM Failed
const EsmeRreplacefail uint32 = 0x00000013     // Replace SM Failed
const EsmeRmsgqful uint32 = 0x00000014         // Message Queue Full
const EsmeRinvsertyp uint32 = 0x00000015       // Invalid Service Type
const EsmeRinvnumdests uint32 = 0x00000033     // Invalid number of destinations
const EsmeRinvdlname uint32 = 0x00000034       // Invalid Distribution List name
const EsmeRinvdestflag uint32 = 0x00000040     // Destination flag (submit_multi)
const EsmeRinvsubrep uint32 = 0x00000042       // Invalid ‘submit with replace’ request
const EsmeRinvesmsubmit uint32 = 0x00000043    // Invalid esm_SUBMIT field data
const EsmeRcntsubdl uint32 = 0x00000044        // Cannot Submit to Distribution List
const EsmeRsubmitfail uint32 = 0x00000045      // submit_sm or submit_multi failed
const EsmeRinvsrcton uint32 = 0x00000048       // Invalid Source address TON
const EsmeRinvsrcnpi uint32 = 0x00000049       // Invalid Source address NPI
const EsmeRinvdstton uint32 = 0x00000050       // Invalid Destination address TON
const EsmeRinvdstnpi uint32 = 0x00000051       // Invalid Destination address NPI
const EsmeRinvsystyp uint32 = 0x00000053       // Invalid system_type field
const EsmeRinvrepflag uint32 = 0x00000054      // Invalid replace_if_present flag
const EsmeRinvnummsgs uint32 = 0x00000055      // Invalid number of messages
const EsmeRthrottled uint32 = 0x00000058       // Throttling error (ESME has exceeded allowed message limits)
const EsmeRinvsched uint32 = 0x00000061        // Invalid Scheduled Delivery Time
const EsmeRinvexpiry uint32 = 0x00000062       // Invalid message (Expiry time)
const EsmeRinvdftmsgid uint32 = 0x00000063     // Predefined Message Invalid or Not Found
const EsmeRxTAppn uint32 = 0x00000064          // ESME Receiver Temporary App Error Code
const EsmeRxPAppn uint32 = 0x00000065          // ESME Receiver Permanent App Error Code
const EsmeRxRAppn uint32 = 0x00000066          // ESME Receiver Reject Message Error Code
const EsmeRqueryfail uint32 = 0x00000067       // query_sm request failed
const EsmeRinvoptparstream uint32 = 0x000000C0 // Error in the optional part of the PDU Body.
const EsmeRoptparnotallwd uint32 = 0x000000C1  // Optional Parameter not allowed
const EsmeRinvparlen uint32 = 0x000000C2       // Invalid Parameter Length.
const EsmeRmissingoptparam uint32 = 0x000000C3 // Expected Optional Parameter missing
const EsmeRinvoptparamval uint32 = 0x000000C4  // Invalid Optional Parameter Value
const EsmeRdeliveryfailure uint32 = 0x000000FE // Delivery Failure (data_sm_resp)
const EsmeRunknownerr uint32 = 0x000000FF      // Unknown Error

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

func GetStatusName(status uint32) string {
	if Status[status] == "" {
		return fmt.Sprintf("Specific status code %x", status)
	}

	return Status[status]
}
