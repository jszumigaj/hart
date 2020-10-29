package hart

import (
	"encoding/binary"
)

// DefaultFrameFactory creates short frame for Command #0 and LongFrame for others
func DefaultFrameFactory(device DeviceIdentifier, command Command) Frame {
	if command.No() == 0 {
		return ShortFrameFactory(device, command)
	}

	return LongFrameFactory(device, command)
}

// ShortFrameFactory creates Master to slave short frame
func ShortFrameFactory(device DeviceIdentifier, command Command) Frame {
	pre := device.Preambles()
	addr := []byte{device.PollAddress()}
	cmd := command.No()
	data := command.Data()
	return NewFrame(pre, MasterToSlaveShortFrame, addr, cmd, NoResponseStatus, data)
}

// LongFrameFactory creates Master to slave long frame
func LongFrameFactory(device DeviceIdentifier, command Command) Frame {
	pre := device.Preambles()
	addr := getLongAddr(device)
	cmd := command.No()
	data := command.Data()
	return NewFrame(pre, MasterToSlaveLongFrame, addr, cmd, NoResponseStatus, data)
}

// builds 5-bytes length address from device identity
func getLongAddr(device DeviceIdentifier) []byte {
	idBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(idBuf, device.Id())
	return append([]byte{device.ManufacturerId(), device.MfrsDeviceType()}, idBuf[1:]...)
}
