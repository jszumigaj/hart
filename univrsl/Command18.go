package univrsl

import (
	"time"

	"github.com/jszumigaj/hart"
)

// Command18 writes the Tag, Descriptor and Date contained within the device
type Command18 struct {
	tag        string    `json:"tag"`
	descriptor string    `json:"descriptor"`
	date       time.Time `json:"date"`

	// command 18 writes data readed by command 13 and uses the same data fields
	Command13
}

// NewCommand18 creates Command18
func NewCommand18(device hart.DeviceIdentifier, tag, descriptor string, date time.Time) Command18 {
	return Command18{
		tag:        tag,
		descriptor: descriptor,
		date:       date,
		Command13: Command13 {
			device: 	device,
		}}
}

// Description properties
func (c *Command18) Description() string { return "Write tag, descriptor and date" }

// No properties
func (c *Command18) No() byte { return 18 }

// Status returns command status
func (c *Command18) Status() hart.CommandStatus { return c.status }

// Data to send
func (c *Command18) Data() []byte {
	tag := NewPackedASCII(c.tag, 6)
	descriptor := NewPackedASCII(c.descriptor, 12)
	date := []byte{
		byte(c.date.Day()),
		byte(c.date.Month()),
		byte(c.date.Year() - 1900),
	}

	var buffer = make([]byte, 21)
	copy(buffer, tag)
	copy(buffer[6:], descriptor)
	copy(buffer[18:], date)
	return buffer
}

//Command18 inherits SetData method from Command13
// func (c *Command18) SetData(data []byte, status hart.CommandStatus) bool { }
