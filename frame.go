// Package hart provides HART protocol core functions
package hart

import "github.com/jszumigaj/hart/status"

// Delimiter constants
const (
	MasterToSlaveShortFrame = 0x02
	MasterToSlaveLongFrame  = 0x82
	SlaveToMasterShortFrame = 0x06
	SlaveToMasterLongFrame  = 0x86
)

// Frame contains HART frame data and methods to communicate through
type Frame struct {
	preambles      int
	delimiter      byte
	address        []byte
	command        byte
	responseStatus []byte
	data           []byte
}

// FrameZero is default ShortFrame with Command0 data (identification)
var FrameZero = Frame{
	preambles: 5,
	delimiter: MasterToSlaveShortFrame,
	address:   []byte{0},
}

// EmptyResponseStatus provides empty response status
var EmptyResponseStatus = []byte{}

// NoData provides empty buffer as data
var NoData = []byte{}

// NewFrame constructor. By default it creates SecondaryMaster frame. Use f.AsPrimaryMaster() method to change it.
func NewFrame(preambles byte, delimiter byte, address []byte, command byte, responseStatus []byte, data []byte) Frame {
	address[0] &= 0x7f
	return Frame{
		preambles:      int(preambles),
		delimiter:      delimiter,
		address:        address,
		command:        command,
		responseStatus: responseStatus,
		data:           data,
	}
}

// Parse buffer to HartFrame. Returns HartFrame and ok status (==true).
// If buffer continas invalid data returns nil, false
// If buffer contains valid data but only CRC is invalid - returns frame, false
func Parse(buffer []byte) (*Frame, bool) {
	index := 0
	frame := Frame{}

	// parse preambles
	for _, b := range buffer {
		if b == 0xff {
			frame.preambles++
			index++
		} else {
			break
		}
	}

	if frame.preambles < 2 || frame.preambles > 20 || frame.preambles == len(buffer) {
		return nil, false
	}

	// check buffer length
	if len(buffer) < index+2 {
		return nil, false
	}

	// delimiter
	frame.delimiter = buffer[index]
	index++

	// check delimiter and parse address
	switch frame.delimiter {
	case SlaveToMasterShortFrame, MasterToSlaveShortFrame:
		frame.address = buffer[index : index+1]
		index++
	case SlaveToMasterLongFrame, MasterToSlaveLongFrame:
		frame.address = buffer[index : index+5]
		index += 5
	default:
		return nil, false
	}

	// check buffer length here
	if len(buffer) < index+3 {
		return nil, false
	}

	// parse command
	frame.command = buffer[index]
	index++

	// calculate response status length
	responseStatusLen := 0
	switch frame.delimiter {
	case SlaveToMasterShortFrame, SlaveToMasterLongFrame:
		responseStatusLen = 2
	}

	// parse bytes count
	dataLength := int(buffer[index]) - responseStatusLen
	index++

	// check buffer length again here...
	if len(buffer) < index+responseStatusLen+dataLength+1 {
		return nil, false
	}

	// parse response status
	frame.responseStatus = buffer[index : index+responseStatusLen]
	index += responseStatusLen

	// parse data
	frame.data = buffer[index : index+dataLength]
	index += dataLength

	// check crc
	crc := calcCrc(buffer[frame.preambles:index])

	if crc != buffer[index] {
		return &frame, false //frame with bad CRC
	}

	// return frame, success
	return &frame, true
}

// Buffer frame
func (f *Frame) Buffer() []byte {
	buf := make([]byte, f.Length())
	// copy preambles
	index := 0
	for index < f.preambles {
		buf[index] = 0xff
		index++
	}
	// copy delimiter
	buf[index] = f.delimiter
	index++
	// copy address
	index += copy(buf[index:], f.address)
	// copy command
	buf[index] = f.command
	index++
	// copy data length
	buf[index] = f.bytesCount()
	index++
	// copy response status
	index += copy(buf[index:], f.responseStatus)
	// copy data
	index += copy(buf[index:], f.data)
	// copy crc
	buf[index] = calcCrc(buf[f.preambles:])
	return buf
}

// Length of frame buffer
func (f *Frame) Length() int {
	return f.preambles + 1 /*delimiter*/ + len(f.address) + 1 /*command*/ + 1 /*Bytes count*/ + int(f.bytesCount()) + 1
}

// CommandStatus returns command status -OR- return communications error summary flags if MSB is 1
func (f *Frame) CommandStatus() CommandStatus {

	if len(f.responseStatus) == 0 {
		return CommandStatus(status.None)
	}

	response := f.responseStatus[0]

	if response&0x80 == 0x80 {
		return status.CommunicationsErrorSummaryFlags(response & 0x7f)
	}
	return status.CommandSpecificStatus(response)
}

// DeviceStatus returns field device status flags
func (f *Frame) DeviceStatus() status.FieldDeviceStatus {

	if len(f.responseStatus) < 2 {
		return status.FieldDeviceStatus(0)
	}

	return status.FieldDeviceStatus(f.responseStatus[1])
}

// Data returns frame data
func (f *Frame) Data() []byte {
	return f.data
}

// IsPrimaryMaster returns true if frame is addressed to/from Secondary Master
func (f *Frame) IsPrimaryMaster() bool {
	return f.address[0]&0x80 == 0x80
}

// AsSecondaryMaster clear primary bit in address
func (f *Frame) AsSecondaryMaster() *Frame {
	f.address[0] &= 0x7f
	return f
}

// AsPrimaryMaster set primary bit in address
func (f *Frame) AsPrimaryMaster() *Frame {
	f.address[0] |= 0x80
	return f
}

func (f *Frame) bytesCount() byte {
	return byte(len(f.responseStatus) + len(f.data))
}

func calcCrc(buffer []byte) byte {
	var crc byte
	for _, b := range buffer {
		crc ^= b
	}
	return crc
}
