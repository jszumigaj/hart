package univrsl

import (
	"github.com/jszumigaj/hart"
)

// Command12 reads message (packedAscii)
type Command12 struct {
	status hart.CommandStatus

	// command data fields
	Msg string `json:"message"`
}

// Description properties
func (c *Command12) Description() string { return "Read message" }

// No properties
func (c *Command12) No() byte { return 12 }

// Data to send
func (c *Command12) Data() []byte { return hart.NoData }

// Status returns command status
func (c *Command12) Status() hart.CommandStatus { return c.status }

// SetData parse received data
func (c *Command12) SetData(data []byte, status hart.CommandStatus) bool {
	c.status = status

	if len(data) < 24 {
		return false
	}

	var packASCII = PackedASCII(data)
	c.Msg = packASCII.String()
	return true
}

// Message returns message
func (d *Command12) Message() string {
	return d.Msg
}
