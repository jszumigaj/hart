package univrsl

import (
	"encoding/binary"

	"github.com/jszumigaj/hart"
)

// Command0 implements Command interface:
type command0 struct {
	device *Device
	status hart.CommandStatus

	// command data fields
	DevId                     uint32 `json:"device_id"`
	MfrsId                    byte   `json:"manufacturer_id"`
	DevType                   byte   `json:"device_type"`
	PollAddr                  byte   `json:"polling_address"`
	Prmbles                   byte   `json:"preambles"`
	HartProtocolMajorRevision byte   `json:"hart_protocol_major_revision"`
	RevisionLevel             byte   `json:"revision_level"`
	SoftwareRevisionLevel     byte   `json:"software_revision_level"`
	HardwareRevisionLevel     byte   `json:"hardware_revision_level"`
	PhisicalSignalingCode     byte   `json:"phisical_signaling_code"`
	Flags                     byte   `json:"flags"`
}

// Device properties
func (c *command0) Device() hart.DeviceIdentifier { return c.device }

// Description properties
func (c *command0) Description() string { return "Device identification" }

// No properties
func (c *command0) No() byte { return 0 }

// Data to send
func (c *command0) Data() []byte { return hart.NoData }

// Status returns command status
func (c *command0) Status() hart.CommandStatus { return c.status }

// SetData parse received data
func (c *command0) SetData(data []byte, status hart.CommandStatus) bool {
	c.status = status

	if len(data) < 12 {
		return false
	}
	if data[0] != 0xfe {
		return false
	}

	c.device.MfrsId = data[1]
	c.device.DevType = data[2]
	c.device.Prmbles = data[3]
	c.device.HartProtocolMajorRevision = data[4]
	c.device.RevisionLevel = data[5]
	c.device.SoftwareRevisionLevel = data[6]
	c.device.HardwareRevisionLevel = byte(data[7] >> 3)
	c.device.PhisicalSignalingCode = byte(data[7] & 0x07)
	c.device.Flags = data[8]
	c.device.DevId = getDeviceId(data[9:])
	return true
}

// return deviceId from 3-bytes slice
func getDeviceId(data []byte) uint32 {
	if len(data) != 3 {
		panic("only 3 bytes length device id slice is accepted")
	}
	data = append([]byte{0}, data...)
	return binary.BigEndian.Uint32(data)
}
