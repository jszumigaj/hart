package univrsl

import (
	"fmt"

	"github.com/jszumigaj/hart/status"
)

// Device implements DeviceIdentifier
type Device struct {
	status status.FieldDeviceStatus

	// embeded command data:
	cmd0Data
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