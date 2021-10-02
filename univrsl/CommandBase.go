package univrsl

import (
	"github.com/jszumigaj/hart"
)

type commandBase struct {
	device hart.DeviceIdentifier
	status hart.CommandStatus
}

func (c *commandBase) DeviceId() hart.DeviceIdentifier { return c.device }

// Status returns command status
func (c *commandBase) Status() hart.CommandStatus { return c.status }

// Data to send
func (c *commandBase) Data() []byte { return hart.NoData }
