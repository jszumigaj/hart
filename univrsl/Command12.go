package univrsl

import (
	"github.com/jszumigaj/hart"
)

// Command12 reads message (packedAscii)
type Command12 struct {
	commandBase

	// command data fields
	Message string `json:"message"`
}

func NewCommand12(device hart.DeviceIdentifier) *Command12 { 
	return &Command12{commandBase: commandBase{device: device}} 
}

// Description property
func (c *Command12) Description() string { return "Read message" }

// No property
func (c *Command12) No() byte { return 12 }

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
