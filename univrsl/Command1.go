package univrsl

import (
	"github.com/jszumigaj/hart"
)

// Command1 reads Primary variable and unit
type Command1 struct {
	status hart.CommandStatus

	PV   float32  `json:"pv"`
	Unit UnitCode `json:"pv_unit"`
}

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
		c.PV = val
		c.Unit = unit
		return true
	}

	return false
}
