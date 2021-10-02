package univrsl

import (
	"github.com/jszumigaj/hart"
)


// Command1 reads Primary variable and unit
type Command1 struct {
	commandBase

	PV   float32  `json:"pv"`
	Unit UnitCode `json:"pv_unit"`
}


func NewCommand1(device hart.DeviceIdentifier) *Command1 { 
	return &Command1{commandBase: commandBase{device: device}} 
}

// Description property
func (c *Command1) Description() string { return "Read primary variable" }

// No property
func (c *Command1) No() byte { return 1 }

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
