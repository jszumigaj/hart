package univrsl

import (
	"time"

	"github.com/jszumigaj/hart"
)

// Command13 reads the Tag, Descriptor and Date contained within the device
type Command13 struct {
	status hart.CommandStatus

	// command data fields
	Tag        string    `json:"tag"`
	Descriptor string    `json:"descriptor"`
	Date       time.Time `json:"date"`
}

// Description properties
func (c *Command13) Description() string { return "Read tag, descriptor and date" }

// No properties
func (c *Command13) No() byte { return 13 }

// Data to send
func (c *Command13) Data() []byte { return hart.NoData }

// Status returns command status
func (c *Command13) Status() hart.CommandStatus { return c.status }

// SetData parse received data
func (c *Command13) SetData(data []byte, status hart.CommandStatus) bool {
	c.status = status

	if len(data) < 21 {
		return false
	}

	c.Tag = PackedASCII(data[:6]).String()
	c.Descriptor = PackedASCII(data[6:12]).String()

	var used bool = data[18] != 250 && data[19] != 250 && data[20] != 250
	var probablyValid bool = data[18] < 31 && data[19] < 12

	if used && probablyValid {
		c.Date = time.Date(int(data[20])+1900, time.Month(data[19]), int(data[18]), 0, 0, 0, 0, time.UTC)
	} else {
		c.Date = time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)

	}
	return true
}
