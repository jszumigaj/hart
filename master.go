package hart

import "github.com/jszumigaj/hart/status"

//FrameSender is interface used by CommandExecutor. It wraps method used to send frame.
//go:generate mockgen -destination=mocks/mock_modem.go -package=mocks . FrameSender
type FrameSender interface {
	SendFrame(rx, tx []byte) (int, error)
}

// Command is interface that wraps the basic HART command methods.
// Obiects implemented command interface can be executed by CommandExecutor
//go:generate mockgen -destination=mocks/mock_command.go -package=mocks . Command
type Command interface {
	DeviceId() DeviceIdentifier
	No() byte
	Description() string
	Status() CommandStatus
	Data() []byte
	SetData([]byte, CommandStatus) bool
}

// CommandStatus is the interface wrapped CommunicationsErrorSummaryFlags and CommandSpecificStatus as one common status
// The returned type depends of MSB bit of the first frame status byte.
// If MSB=1 this byte means CommunicationsError otherwise it is CommandStatus
// Some commands can return individual command specific status as CommandStatus type
//go:generate mockgen -destination=mocks/mock_commandStatus.go -package=mocks . CommandStatus
type CommandStatus interface {
	IsError() bool
	IsWarning() bool
	String() string
}

// DeviceIdentifier is the interface that wraps the basic Device methods
// Commands use this information to create valid HART frame
//go:generate mockgen -destination=mocks/mock_device.go -package=mocks . DeviceIdentifier
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
type FrameFactory func(Command) Frame

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
	txFrame := m.FrameFactory(command)
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
		return nil, &status.FrameParsingError{Frame: rxBuffer}
	}

	// frame is ok, set device status and command status
	result = rxFrame.CommandStatus()
	
	// checking status for communications errors
	if commError, ok := result.(status.CommunicationsErrorSummaryFlags); ok {
		return result, commError
	}

	// communication was ok, set device status and parse command data
	device := command.DeviceId()
	device.SetStatus(rxFrame.DeviceStatus())
	if ok := command.SetData(rxFrame.Data(), result); !ok {
		return result, &status.FrameDataParsingError{Frame: rxBuffer}
	}

	// everything looks good, get command specific status and return as func result
	result = command.Status()
	return result, nil
}

// ExecuteAsync executes more commands asynchronously - proof of concept:
 func (m *Master) ExecuteAsync(ch chan<- Command, commands ...Command) {

	go func() error {
		for _, cmd := range commands {
			if _, err := m.Execute(cmd); err != nil {
				return err
			}
			ch <- cmd
		}

		return nil
	}()
}
