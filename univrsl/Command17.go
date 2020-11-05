package univrsl

import (
	"github.com/jszumigaj/hart"
)

// command3 implements Command interface:
type Command17 struct {
	status hart.CommandStatus

	// command data fields
	Msg string
}

// Description properties
func (c *Command17) Description() string { return "Write message" }

// No properties
func (c *Command17) No() byte { return 17 }

// Data to send
func (c *Command17) Data() []byte {
	packed := NewPackedASCII(c.Msg, 24)
	return packed
}

// Status returns command status
func (c *Command17) Status() hart.CommandStatus { return c.status }

// SetData parse received data
func (c *Command17) SetData(data []byte, status hart.CommandStatus) bool {
	c.status = status

	if len(data) < 24 {
		return false
	}

	var packASCII = PackedASCII(data)
	c.Msg = packASCII.String()
	return true
}
