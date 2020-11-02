package hart

import "github.com/jszumigaj/hart/status"

// FrameSender is interface used by CommandExecutor. It wraps method used to send frame.
// go:generate mockgen -destination=mocks/mock_modem.go -package=mocks . FrameSender
type FrameSender interface {
	SendFrame(rx, tx []byte) (int, error)
}

// Command is interface that wraps the basic HART command methods.
// Obiects implemented command interface can be executed by CommandExecutor
type Command interface {
	No() byte
	Description() string
	Device() DeviceIdentifier
	Status() CommandStatus
	Data() []byte
	SetData([]byte, CommandStatus) bool
}

// DeviceIdentifier is the interface that wraps the basic Device methods
// Commands use this information to create valid HART frame
type DeviceIdentifier interface {
	ManufacturerId() byte
	MfrsDeviceType() byte
	Id() uint32
	PollAddress() byte
	Preambles() byte
	Status() status.FieldDeviceStatus
	SetStatus(status.FieldDeviceStatus)
}

// FrameFactory is the func used as factory to create frames by the executor.
// Client can use one of predefined hart.ShortFrameFactory or hart.LongFrameFactory
type FrameFactory func(DeviceIdentifier, Command) Frame

// Master executes command. Set Primary property to true to send frame as Primary master
type Master struct {
	modem        FrameSender
	FrameFactory FrameFactory
	Primary      bool
}

// NewMaster creates Master object
func NewMaster(modem FrameSender) *Master {
	return &Master{
		modem:        modem,
		FrameFactory: DefaultFrameFactory,
	}
}

// Execute method executes HART command
func (m *Master) Execute(command Command) (CommandStatus, error) {
	device := command.Device()
	txFrame := m.FrameFactory(device, command)
	if m.Primary {
		txFrame.AsPrimaryMaster()
	}
	rxBuffer := make([]byte, 128)
	var result CommandStatus = nil

	count, err := m.modem.SendFrame(txFrame.Buffer(), rxBuffer)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, status.ErrNoResponse
	}

	// try parse frame
	var rxFrame *Frame
	var ok bool
	if rxFrame, ok = Parse(rxBuffer); !ok {
		return nil, &status.FrameParsingError{rxBuffer}
	}

	// frame is ok, set device status and command status
	result = rxFrame.CommandStatus()
	// checking status for communications errors
	switch result.(type) {
	case status.CommunicationsErrorSummaryFlags:
		return result, nil
	}

	// communication was ok, set device status and parse command data
	device.SetStatus(rxFrame.DeviceStatus())
	if ok := command.SetData(rxFrame.Data(), result); !ok {
		return result, &status.FrameDataParsingError{rxBuffer}
	}

	// everything looks good, get command specific status and return as func result
	result = command.Status()
	return result, nil
}

// ExecuteAsync executes more commands asynchronously - proof of concept:
func (m *Master) ExecuteAsync(ch chan<- Command, commands ...Command) error {

	//panic("Not implemented!")

	for _, cmd := range commands {
		if _, err := m.Execute(cmd); err != nil {
			return err
		}
		ch <- cmd
	}

	return nil
}
