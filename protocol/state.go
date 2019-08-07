package protocol

// SMPP v3.4 - 5.2.28 page 130
const StateEnroute uint32 = 1
const StateDelivered uint32 = 2
const StateExpired uint32 = 3
const StateDeleted uint32 = 4
const StateUndeliverable uint32 = 5
const StateAccepted uint32 = 6
const StateUnknown uint32 = 7
const StateRejected uint32 = 8

var State = map[uint32]string{
	StateEnroute:       "ENROUTE",
	StateDelivered:     "DELIVRD",
	StateExpired:       "EXPIRED",
	StateDeleted:       "DELETED",
	StateUndeliverable: "UNDELIV",
	StateAccepted:      "ACCEPTD",
	StateUnknown:       "UNKNOWN",
	StateRejected:      "REJECTD",
}

func GetStateName(state uint32) string {
	if State[state] == "" {
		return State[StateUnknown]
	}

	return State[state]
}
