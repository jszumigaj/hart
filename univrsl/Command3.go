package univrsl

import (
	"github.com/jszumigaj/hart"
)

// Command3 reads for variables with units
type Command3 struct {
	device hart.DeviceIdentifier
	status hart.CommandStatus

	// command data fields
	Current float32  `json:"Current"`
	Pv      float32  `json:"pv"`
	PvUnit  UnitCode `json:"pv_unit"`
	Sv      float32  `json:"sv"`
	SvUnit  UnitCode `json:"sv_unit"`
	Tv      float32  `json:"tv"`
	TvUnit  UnitCode `json:"tv_unit"`
	Fv      float32  `json:"fv"`
	FvUnit  UnitCode `json:"fv_unit"`
}

func NewCommand3(device hart.DeviceIdentifier) *Command3 { return &Command3{device: device} }

func (c *Command3) DeviceId() hart.DeviceIdentifier { return c.device }

// Description properties
func (c *Command3) Description() string { return "Read PV current and dynamic variables" }

// No properties
func (c *Command3) No() byte { return 3 }

// Data to send
func (c *Command3) Data() []byte { return hart.NoData }

// Status returns command status
func (c *Command3) Status() hart.CommandStatus { return c.status }

// SetData parse received data
func (c *Command3) SetData(data []byte, status hart.CommandStatus) bool {
	c.status = status

	if len(data) < 9 {
		return false
	}

	if val, ok := getFloat(data); ok {
		c.Current = val
	}

	if val, unit, ok := getFloatWithUnit(data[4:]); ok {
		c.Pv = val
		c.PvUnit = unit
	}

	if val, unit, ok := getFloatWithUnit(data[9:]); ok {
		c.Sv = val
		c.SvUnit = unit
	}

	if val, unit, ok := getFloatWithUnit(data[14:]); ok {
		c.Tv = val
		c.TvUnit = unit
	}

	if val, unit, ok := getFloatWithUnit(data[19:]); ok {
		c.Fv = val
		c.FvUnit = unit
	}

	return true
}
