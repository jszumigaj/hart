package univrsl

import (
	"github.com/jszumigaj/hart"
)

// Command12 reads message (packedAscii)
type Command12 struct {
	device hart.DeviceIdentifier
	status hart.CommandStatus

	// command data fields
	Message string `json:"message"`
}

func NewCommand12(device hart.DeviceIdentifier) *Command12 { return &Command12{device: device} }

func (c *Command12) DeviceId() hart.DeviceIdentifier { return c.device }

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
	c.Message = packASCII.String()
	return true
}
