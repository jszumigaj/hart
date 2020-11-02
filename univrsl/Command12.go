package univrsl

import (
	"github.com/jszumigaj/hart"
)

// command3 implements Command interface:
type command12 struct {
	device *Device
	status hart.CommandStatus

	// command data fields
	Msg string `json:"message"`
}

// Device properties
func (c *command12) Device() hart.DeviceIdentifier { return c.device }

// Description properties
func (c *command12) Description() string { return "Read message" }

// No properties
func (c *command12) No() byte { return 17 }

// Data to send
func (c *command12) Data() []byte { return hart.NoData }

// Status returns command status
func (c *command12) Status() hart.CommandStatus { return c.status }

// SetData parse received data
func (c *command12) SetData(data []byte, status hart.CommandStatus) bool {
	c.status = status

	if len(data) < 24 {
		return false
	}

	var packASCII = PackedASCII(data)
	c.device.Msg = packASCII.String()
	return true
}

// Message returns message
func (d *Device) Message() string {
	return d.Msg
}
