package univrsl

import (
	"github.com/jszumigaj/hart"
)

// Command2 implements Command interface:
type Command2 struct {
	status hart.CommandStatus

	// data fields:
	Curr float32 `json:"current"`
	PoR  float32 `json:"percent_of_range"`
}

// Description properties
func (c *Command2) Description() string { return "Read current and percent of range" }

// No properties
func (c *Command2) No() byte { return 2 }

// Data to send
func (c *Command2) Data() []byte { return hart.NoData }

// Status returns command status
func (c *Command2) Status() hart.CommandStatus { return c.status }

// SetData parse received data
func (c *Command2) SetData(data []byte, status hart.CommandStatus) bool {
	c.status = status

	var val float32
	var ok bool

	if val, ok = getFloat(data); ok {
		c.Curr = val
	}

	if val, ok = getFloat(data[4:]); ok {
		c.PoR = val
	}

	return ok
}

// Current returns analog output current readed by Command2
func (d *Command2) Current() float32 { return d.Curr }

// PercentOfRange returs percent of range output
func (d *Command2) PercentOfRange() float32 { return d.PoR }
