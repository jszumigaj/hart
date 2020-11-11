package univrsl

import (
	"github.com/jszumigaj/hart"
)

// Command17 writes message
type Command17 struct {
	status hart.CommandStatus

	// This command inherits fields from command 12
	Command12
}

// NewCommand17 creates Command 17
func NewCommand17(message string) Command17 {
	return Command17{
		Command12: Command12{Message: message},
	}
}

// Description properties
func (c *Command17) Description() string { return "Write message" }

// No properties
func (c *Command17) No() byte { return 17 }

// Data to send
func (c *Command17) Data() []byte {
	packed := NewPackedASCII(c.Message, 24)
	return packed
}

// Status returns command status
func (c *Command17) Status() hart.CommandStatus { return c.status }

// SetData is inherited from Command12
//func (c *Command17) SetData(data []byte, status hart.CommandStatus) bool {}
