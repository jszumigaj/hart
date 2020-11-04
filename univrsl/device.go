package univrsl

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/jszumigaj/hart"
	"github.com/jszumigaj/hart/status"
)

// Device implements DeviceIdentifier
type Device struct {
	status status.FieldDeviceStatus

	// embeded command data:
	cmd0Data
	// command1
	// command2
	// command3

	// command12
	// command13
	// command17
}

func (d *Device) String() string {
	return fmt.Sprintf("Id: %05d Mfr: 0x%02x Type: 0x%02x", d.DevId, d.MfrsId, d.DevType)
}

// Id is DeviceIdentifier method implementation
func (d *Device) Id() uint32 { return d.DevId }

// ManufacturerId is DeviceIdentifier method implementation
func (d *Device) ManufacturerId() byte { return d.cmd0Data.MfrsId }

// MfrsDeviceType is DeviceIdentifier method implementation
func (d *Device) MfrsDeviceType() byte { return d.DevType }

// PollAddress is DeviceIdentifier method implementation
func (d *Device) PollAddress() byte { return d.PollAddr }

// Preambles is DeviceIdentifier method implementation
func (d *Device) Preambles() byte {
	if d.Prmbles < 5 {
		d.Prmbles = 5
	}
	return d.Prmbles
}

// Status is DeviceIdentifier method implementation
func (d *Device) Status() status.FieldDeviceStatus { return d.status }

// SetStatus is DeviceIdentifier method implementation
func (d *Device) SetStatus(status status.FieldDeviceStatus) { d.status = status }

// Command0 creates command for reading HART Command #0 (Identify slave device)
func (d *Device) Command0() hart.Command { return &Command0{device: d} }

// Command1 creates command for reading HART Command #1 (Read PV)
func (d *Device) Command1() hart.Command { return &Command1{device: d} }

// Command2 creates command for reading HART Command #2 (Read current and percent of range)
func (d *Device) Command2() hart.Command { return &Command2{device: d} }

// Command3 creates command for reading HART Command #3 (Read primary variables)
func (d *Device) Command3() hart.Command { return &Command3{device: d} }

// Command12 creates command for reading HART Command #12 (Read Message)
func (d *Device) Command12() hart.Command { return &Command12{device: d} }

// Command13 creates command for reading HART Command #13 (Read tag, descriptor, date)
func (d *Device) Command13() hart.Command { return &Command13{device: d} }

// Command17 creates command for reading HART Command #12 (Write Message)
func (d *Device) Command17(message string) hart.Command {
	return &Command17{device: d, Msg: message}
}

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
