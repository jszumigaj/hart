package univrsl

import (
	"github.com/jszumigaj/hart"
)

// Command1 implements Command interface:
type Command1 struct {
	device *Device
	status hart.CommandStatus

	// command data fields embedded into device
	Pv     float32  `json:"pv"`
	PvUnit UnitCode `json:"pv_unit"`
}

// Device properties
func (c *Command1) Device() hart.DeviceIdentifier { return c.device }

// Description properties
func (c *Command1) Description() string { return "Read primary variable" }

// No properties
func (c *Command1) No() byte { return 1 }

// Data to send
func (c *Command1) Data() []byte { return hart.NoData }

// Status returns command status
func (c *Command1) Status() hart.CommandStatus { return c.status }

// SetData parse received data
func (c *Command1) SetData(data []byte, status hart.CommandStatus) bool {
	c.status = status

	if val, unit, ok := getFloatWithUnit(data); ok {
		c.Pv = val
		c.PvUnit = unit
		return true
	}

	return false
}

// PV returns Primary variable value readed in Command1
func (c *Command1) PV() (float32, UnitCode) {
	return c.Pv, c.PvUnit
}
