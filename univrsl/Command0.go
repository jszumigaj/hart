package univrsl

import (
	"encoding/binary"

	"github.com/jszumigaj/hart"
)

// contains data readed by Command #0. This struct is embeded into univrsl.Device type
type cmd0Data struct {
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

// Command0 reads device identification
type Command0 struct {
	device *Device
	status hart.CommandStatus

	// command data fields
	cmd0Data
}

func NewCommand0(device *Device) *Command0 { return &Command0{device: device} }

func (c *Command0) DeviceId() hart.DeviceIdentifier { return c.device }

// Description properties
func (c *Command0) Description() string { return "Device identification" }

// No properties
func (c *Command0) No() byte { return 0 }

// Data to send
func (c *Command0) Data() []byte { return hart.NoData }

// Status returns command status
func (c *Command0) Status() hart.CommandStatus { return c.status }

// SetData parse received data
func (c *Command0) SetData(data []byte, status hart.CommandStatus) bool {
	c.status = status

	if len(data) < 12 {
		return false
	}
	if data[0] != 0xfe {
		return false
	}

	// commands 0 has special meaning in HART protocol, becaouse it provides device identification
	// so it have to delegate all properties to the device
	// and have to embedded commands0 data structure
	c.device.MfrsId = data[1]
	c.device.DevType = data[2]
	c.device.Prmbles = data[3]
	c.device.HartProtocolMajorRevision = data[4]
	c.device.RevisionLevel = data[5]
	c.device.SoftwareRevisionLevel = data[6]
	c.device.HardwareRevisionLevel = byte(data[7] >> 3)
	c.device.PhisicalSignalingCode = byte(data[7] & 0x07)
	c.device.Flags = data[8]
	c.device.DevId = getDeviceId(data[9:12])
	return true
}

// return deviceId from 3-bytes slice
func getDeviceId(data []byte) uint32 {
	if len(data) != 3 {
		panic("DeviceId should be 3-bytes length slice")
	}
	data = append([]byte{0}, data...)
	return binary.BigEndian.Uint32(data)
}
