package device

import (
	"github.com/jszumigaj/hart"
)

// Command1 implements Command interface:
type command1 struct {
	device *UniversalDevice
	status hart.CommandStatus

	// command data fields
	Pv     float32  `json:"pv"`
	PvUnit UnitCode `json:"pv_unit"`
}

// Device properties
func (c *command1) Device() hart.DeviceIdentifier { return c.device }

// Description properties
func (c *command1) Description() string { return "Read primary variable" }

// No properties
func (c *command1) No() byte { return 1 }

// Data to send
func (c *command1) Data() []byte { return hart.NoData }

// Status returns command status
func (c *command1) Status() hart.CommandStatus { return c.status }

// SetData parse received data
func (c *command1) SetData(data []byte, status hart.CommandStatus) bool {
	c.status = status

	if val, unit, ok := getFloatWithUnit(data); ok {
		c.device.Pv = val
		c.device.PvUnit = unit
		return true
	}

	return false
}

func (d *UniversalDevice) PV() (float32, UnitCode) {
	return d.Pv, d.PvUnit
}

// type Cmd1Data struct {
// 	PV      float32 `json:"some_field"`
// 	PvUnit  UnitCode
// 	Example int
// }
