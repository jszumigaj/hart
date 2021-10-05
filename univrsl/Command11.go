package univrsl

import (
	
)

// Command11 - Read unique identifier associated with TAG.
type Command11 struct {
	tag string
	
	// command data fields and methods inherited from Command0
	Command0
}

// NewCommand11 creates Command 11
func NewCommand11(device *Device, tag string) *Command11 {
	return &Command11 {
		Command0: Command0{device: device},
		tag: tag,
	}
}

// Description property
func (c *Command11) Description() string { return "Read unique identifier associated with TAG" }

// No property
func (c *Command11) No() byte { return 11 }

// Data to send
func (c *Command11) Data() []byte {
	packed := NewPackedASCII(c.tag, 6)
	return packed
}

// SetData is inherited from Command0
//func (c *Command11) SetData(data []byte, status hart.CommandStatus) bool {}
