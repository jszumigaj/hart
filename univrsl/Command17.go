package univrsl

import (
	"github.com/jszumigaj/hart"
)

// command3 implements Command interface:
type command17 struct {
	device *Device
	status hart.CommandStatus

	// command data fields
	msg string
}

// Device properties
func (c *command17) Device() hart.DeviceIdentifier { return c.device }

// Description properties
func (c *command17) Description() string { return "Write message" }

// No properties
func (c *command17) No() byte { return 17 }

// Data to send
func (c *command17) Data() []byte {
	packed := NewPackedASCII(c.msg, 24)
	return packed
}

// Status returns command status
func (c *command17) Status() hart.CommandStatus { return c.status }

// SetData parse received data
func (c *command17) SetData(data []byte, status hart.CommandStatus) bool {
	c.status = status

	if len(data) < 24 {
		return false
	}

	var packASCII = PackedASCII(data)
	c.device.Msg = packASCII.String()
	return true
}
