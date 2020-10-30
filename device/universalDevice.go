package device

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/jszumigaj/hart"
)

// UniversalDevice implements DeviceIdentifier
type UniversalDevice struct {
	status hart.FieldDeviceStatus

	// embeded command data:
	command0
	command1
	command2
	command3
}

func (d *UniversalDevice) String() string {
	return fmt.Sprintf("Id: %05d Mfr: 0x%02x Type: 0x%02x", d.DevId, d.MfrsId, d.DevType)
}

// Id is DeviceIdentifier method implementation
func (d *UniversalDevice) Id() uint32 { return d.DevId }

// ManufacturerId is DeviceIdentifier method implementation
func (d *UniversalDevice) ManufacturerId() byte { return d.MfrsId }

// MfrsDeviceType is DeviceIdentifier method implementation
func (d *UniversalDevice) MfrsDeviceType() byte { return d.DevType }

// PollAddress is DeviceIdentifier method implementation
func (d *UniversalDevice) PollAddress() byte { return d.PollAddr }

// Preambles is DeviceIdentifier method implementation
func (d *UniversalDevice) Preambles() byte {
	if d.Prmbles < 5 {
		d.Prmbles = 5
	}
	return d.Prmbles
}

// Status is DeviceIdentifier method implementation
func (d *UniversalDevice) Status() hart.FieldDeviceStatus { return d.status }

// SetStatus is DeviceIdentifier method implementation
func (d *UniversalDevice) SetStatus(status hart.FieldDeviceStatus) { d.status = status }

// Command0 creates command for reading HART Command #0
func (d *UniversalDevice) Command0() hart.Command { return &command0{device: d} }

// Command1 creates command for reading HART Command #1
func (d *UniversalDevice) Command1() hart.Command { return &command1{device: d} }

// Command2 creates command for reading HART Command #2
func (d *UniversalDevice) Command2() hart.Command { return &command2{device: d} }

// Command3 creates command for reading HART Command #3
func (d *UniversalDevice) Command3() hart.Command { return &command3{device: d} }

// UnitCode is an alias to hart.UnitCode
type UnitCode = hart.UnitCode

func getFloat(buf []byte) (float32, bool) {
	if len(buf) < 4 {
		return float32(math.NaN()), false
	}

	bits := binary.BigEndian.Uint32(buf[0:4])
	return math.Float32frombits(bits), true
}

func getFloatWithUnit(buf []byte) (float32, UnitCode, bool) {
	if len(buf) < 5 {
		return float32(math.NaN()), 0, false
	}

	unit := UnitCode(buf[0])
	bits := binary.BigEndian.Uint32(buf[1:5])
	return math.Float32frombits(bits), unit, true
}
