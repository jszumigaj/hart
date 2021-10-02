package univrsl

import (
	"github.com/jszumigaj/hart"
)

// Command17 writes message
type Command17 struct {
	message string
	// This command inherits fields and methods from command 12
	Command12
}

// NewCommand17 creates Command 17
func NewCommand17(device hart.DeviceIdentifier, message string) *Command17 {
	return &Command17 {
		Command12: Command12{commandBase: commandBase{device: device}},
		message: message,
	}
}

// Description property
func (c *Command17) Description() string { return "Write message" }

// No property
func (c *Command17) No() byte { return 17 }

// Data to send
func (c *Command17) Data() []byte {
	packed := NewPackedASCII(c.message, 24)
	return packed
}

// SetData is inherited from Command12
//func (c *Command17) SetData(data []byte, status hart.CommandStatus) bool {}
