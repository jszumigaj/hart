package univrsl

import (
	"time"

	"github.com/jszumigaj/hart"
)

// command3 Read the Tag, Descriptor and Date contained within the device
type command13 struct {
	device *Device
	status hart.CommandStatus

	// command data fields
	Tg       string    `json:"tag"`
	Dscriptr string    `json:"descriptor"`
	Dtime    time.Time `json:"date_time"`
}

// Device properties
func (c *command13) Device() hart.DeviceIdentifier { return c.device }

// Description properties
func (c *command13) Description() string { return "Read tag, descriptor and date" }

// No properties
func (c *command13) No() byte { return 13 }

// Data to send
func (c *command13) Data() []byte { return hart.NoData }

// Status returns command status
func (c *command13) Status() hart.CommandStatus { return c.status }

// SetData parse received data
func (c *command13) SetData(data []byte, status hart.CommandStatus) bool {
	c.status = status

	if len(data) < 21 {
		return false
	}

	c.device.Tg = PackedASCII(data[:6]).String()
	c.device.Dscriptr = PackedASCII(data[6:12]).String()

	var used bool = data[18] != 250 && data[19] != 250 && data[20] != 250
	var probablyValid bool = data[18] < 31 && data[19] < 12

	if used && probablyValid {
		c.device.Dtime = time.Date(int(data[20])+1900, time.Month(data[19]), int(data[18]), 0, 0, 0, 0, time.UTC)
	} else {
		c.device.Dtime = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)

	}
	return true
}

// Tag returns Tag
func (d *Device) Tag() string { return d.Tg }

// Descriptor returns descriptor
func (d *Device) Descriptor() string { return d.Dscriptr }

// Date returns date
func (d *Device) Date() time.Time { return d.Dtime }
