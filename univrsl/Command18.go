package univrsl

import (
	"time"

	"github.com/jszumigaj/hart"
)

// Command18 writes the Tag, Descriptor and Date contained within the device
type Command18 struct {
	status hart.CommandStatus

	// command 18 writes data readed by command 13 and uses the same data fields
	Command13
}

// NewCommand18 creates Command18
func NewCommand18(tag, descriptor string, date time.Time) Command18 {
	return Command18{
		Command13: Command13{
			Tag:        tag,
			Descriptor: descriptor,
			Date:       date,
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
	tag := NewPackedASCII(c.Tag, 6)
	descriptor := NewPackedASCII(c.Descriptor, 12)
	date := []byte{
		byte(c.Date.Day()),
		byte(c.Date.Month()),
		byte(c.Date.Year() - 1900),
	}

	var buffer = make([]byte, 21)
	copy(buffer, tag)
	copy(buffer[6:], descriptor)
	copy(buffer[18:], date)
	return buffer
}

//Command18 inherits SetData method from Command13
// func (c *Command18) SetData(data []byte, status hart.CommandStatus) bool { }
