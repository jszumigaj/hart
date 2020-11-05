package univrsl

import (
	"github.com/jszumigaj/hart"
)

// Command2 reads current and percent of range
type Command2 struct {
	status hart.CommandStatus

	// data fields:
	Current        float32 `json:"current"`
	PercentOfRange float32 `json:"percent_of_range"`
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
		c.Current = val
	}

	if val, ok = getFloat(data[4:]); ok {
		c.PercentOfRange = val
	}

	return ok
}
