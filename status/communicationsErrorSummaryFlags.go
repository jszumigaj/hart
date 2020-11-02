package status

import "strings"

// CommunicationsErrorSummaryFlags Communications Error Summary Flags
// This status is valid if value & 0x80 == 0x80
type CommunicationsErrorSummaryFlags byte

// Communications Error Summary Flags descriptions:
// VerticalParityError: The parity of one or more of hthe bytes received by the device was incorrect.
// OverrunError: At least one byte of data in the receive buffer of the UART was overwritten before it was read.
// FramingError: The Stop Bit of one or more bytes received by the device was not detected by the UART.
// LongitudalParityError: The Longitudal Parity calculated by the device did not match the Longitudal Parity byte at the end of the message
// Reserved: don't use
// BufferOverflow: The message was too long for the receive buffer of the device.
// Undefined: Not defined at this time
// None:  Ok
const (
	VerticalParityError   CommunicationsErrorSummaryFlags = 0x40
	OverrunError          CommunicationsErrorSummaryFlags = 0x20
	FramingError          CommunicationsErrorSummaryFlags = 0x10
	LongitudalParityError CommunicationsErrorSummaryFlags = 0x08
	Reserved              CommunicationsErrorSummaryFlags = 0x04
	BufferOverflow        CommunicationsErrorSummaryFlags = 0x02
	Undefined             CommunicationsErrorSummaryFlags = 0x01
	None                  CommunicationsErrorSummaryFlags = 0
)

var communicationsErrorFlagsDescriptions = map[CommunicationsErrorSummaryFlags]string{
	VerticalParityError:   "Vertical parity error",
	OverrunError:          "Overrun error",
	FramingError:          "Framing error",
	LongitudalParityError: "Longitudal parity error",
	Reserved:              "Reserved",
	BufferOverflow:        "Buffer Overflow error",
	Undefined:             "Undefined",
	None:                  "OK",
}

func (flag CommunicationsErrorSummaryFlags) String() string {
	if flag == None {
		return communicationsErrorFlagsDescriptions[None]
	}

	names := []string{}
	for i := 0; i < 8; i++ {
		mask := CommunicationsErrorSummaryFlags(1 << i)
		if flag.HasFlag(mask) {
			names = append(names, communicationsErrorFlagsDescriptions[mask])
		}
	}

	return strings.Join(names, ", ")
}

// HasFlag f
func (flag CommunicationsErrorSummaryFlags) HasFlag(mask CommunicationsErrorSummaryFlags) bool {
	return flag&mask == mask
}

//IsError CommandStatus inteface member
func (flag CommunicationsErrorSummaryFlags) IsError() bool {
	return flag != 0
}

//IsWarning CommandStatus inteface member
func (flag CommunicationsErrorSummaryFlags) IsWarning() bool {
	return false
}
