package status

import (
	"fmt"
	"strings"
)

// CommunicationsErrorSummaryFlags Communications Error Summary Flags
// This status is valid only if MSB of the status byte is 1
type CommunicationsErrorSummaryFlags byte

// Communications Error Summary Flags
const (
	VerticalParityError   CommunicationsErrorSummaryFlags = 0x40 // The parity of one or more of hthe bytes received by the device was incorrect.
	OverrunError          CommunicationsErrorSummaryFlags = 0x20 // At least one byte of data in the receive buffer of the UART was overwritten before it was read.
	FramingError          CommunicationsErrorSummaryFlags = 0x10 // The Stop Bit of one or more bytes received by the device was not detected by the UART.
	LongitudalParityError CommunicationsErrorSummaryFlags = 0x08 // The Longitudal Parity calculated by the device did not match the Longitudal Parity byte at the end of the message
	Reserved              CommunicationsErrorSummaryFlags = 0x04 // don't use
	BufferOverflow        CommunicationsErrorSummaryFlags = 0x02 // The message was too long for the receive buffer of the device.
	Undefined             CommunicationsErrorSummaryFlags = 0x01 // Not defined at this time
	None                  CommunicationsErrorSummaryFlags = 0    // OK
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

func (flag CommunicationsErrorSummaryFlags) Error() string {
	return fmt.Sprintf("%v (0x%02x)", flag.String(), byte(flag))
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
