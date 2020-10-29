package device

import (
	"github.com/jszumigaj/hart"
)

// command3 implements Command interface:
type command3 struct {
	device *UniversalDevice
	status hart.CommandStatus

	// command data fields
	Sv     float32  `json:"sv"`
	SvUnit UnitCode `json:"sv_unit"`
	Tv     float32  `json:"tv"`
	TvUnit UnitCode `json:"tv_unit"`
	Fv     float32  `json:"fv"`
	FvUnit UnitCode `json:"fv_unit"`
}

// Device properties
func (c *command3) Device() hart.DeviceIdentifier { return c.device }

// Description properties
func (c *command3) Description() string { return "Read PV current and dynamic variables" }

// No properties
func (c *command3) No() byte { return 3 }

// Data to send
func (c *command3) Data() []byte { return hart.NoData }

// Status returns command status
func (c *command3) Status() hart.CommandStatus { return c.status }

// SetData parse received data
func (c *command3) SetData(data []byte, status hart.CommandStatus) bool {
	c.status = status

	if len(data) < 9 {
		return false
	}

	if val, ok := getFloat(data); ok {
		c.device.Curr = val
	}

	if val, unit, ok := getFloatWithUnit(data[4:]); ok {
		c.device.Pv = val
		c.device.PvUnit = unit
	}

	if val, unit, ok := getFloatWithUnit(data[9:]); ok {
		c.device.Sv = val
		c.device.SvUnit = unit
	}

	if val, unit, ok := getFloatWithUnit(data[14:]); ok {
		c.device.Tv = val
		c.device.TvUnit = unit
	}

	if val, unit, ok := getFloatWithUnit(data[19:]); ok {
		c.device.Fv = val
		c.device.FvUnit = unit
	}

	return true
}

func (d *UniversalDevice) SV() (float32, UnitCode) {
	return d.Sv, d.SvUnit
}

func (d *UniversalDevice) TV() (float32, UnitCode) {
	return d.Tv, d.TvUnit
}

func (d *UniversalDevice) FV() (float32, UnitCode) {
	return d.Fv, d.FvUnit
}
