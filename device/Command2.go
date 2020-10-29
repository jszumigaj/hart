package device

import (
	"github.com/jszumigaj/hart"
)

// Command2 implements Command interface:
type command2 struct {
	device *UniversalDevice
	status hart.CommandStatus

	// data fields:
	Curr float32 `json:"current"`
	PoR  float32 `json:"percent_of_range"`
}

// Device properties
func (c *command2) Device() hart.DeviceIdentifier { return c.device }

// Description properties
func (c *command2) Description() string { return "Read current and percent of range" }

// No properties
func (c *command2) No() byte { return 2 }

// Data to send
func (c *command2) Data() []byte { return hart.NoData }

// Status returns command status
func (c *command2) Status() hart.CommandStatus { return c.status }

// SetData parse received data
func (c *command2) SetData(data []byte, status hart.CommandStatus) bool {
	c.status = status

	var val float32
	var ok bool

	if val, ok = getFloat(data); ok {
		c.device.Curr = val
	}

	if val, ok = getFloat(data[4:]); ok {
		c.device.PoR = val
	}

	return ok
}

func (d *UniversalDevice) Current() float32 { return d.Curr }

func (d *UniversalDevice) PercentOfRange() float32 { return d.PoR }
