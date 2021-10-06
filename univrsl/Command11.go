package univrsl

import "github.com/jszumigaj/hart"

// Command11 - Read unique identifier associated with TAG.
type Command11 struct {
	tag string
	useBroadcastAddress bool
	
	// command data fields and methods inherited from Command0
	Command0
}

// NewCommand11 creates Command 11
func NewCommand11(device *Device, tag string, useBroadcastAddress ...bool) *Command11 {
	var broadcast = true

	if len(useBroadcastAddress) > 0 {
		broadcast = useBroadcastAddress[0]
	}

	return &Command11 {
		Command0: Command0{device: device},
		tag: tag,
		useBroadcastAddress: broadcast,
	}
}

func (c *Command11) DeviceId() hart.DeviceIdentifier { 

	// in case of broadcast addres device identity should be cleared (all 5 bytes set zero)
	if c.useBroadcastAddress {
		dev := *c.device
		dev.DevId = 0
		dev.DevType = 0
		dev.MfrsId = 0
		return &dev
	}
	return c.device 
}

// Description property
func (c *Command11) Description() string { return "Read unique identifier associated with TAG" }

// No property
func (c *Command11) No() byte { return 11 }

// Data to send
func (c *Command11) Data() []byte {
	packed := NewPackedASCII(c.tag, 6)
	return packed
}

// SetData is inherited from Command0
//func (c *Command11) SetData(data []byte, status hart.CommandStatus) bool {}
