package status

// CommandSpecificStatus command status
// This status is valid if value & 0x80 == 0x00
type CommandSpecificStatus byte

// CommandSpecificStatus enum
const (
	NoCommandSpecificError          CommandSpecificStatus = 0
	InvalidSelection                CommandSpecificStatus = 2
	PassedParameterTooLarge         CommandSpecificStatus = 3
	PassedParameterTooSmall         CommandSpecificStatus = 4
	TooFewDataBytesReceived         CommandSpecificStatus = 5
	TransmitterSpecificCommandError CommandSpecificStatus = 6
	InWriteProtectMode              CommandSpecificStatus = 7
	WarningUpdateFailure            CommandSpecificStatus = 8
	AccessRestricted                CommandSpecificStatus = 16
	Busy                            CommandSpecificStatus = 32
	CommandNotImplemented           CommandSpecificStatus = 64
)

var commandSpecificStatusDescriptions = map[CommandSpecificStatus]string{
	NoCommandSpecificError:          "No Command-Specific Errors",
	InvalidSelection:                "Invalid Selection",
	PassedParameterTooLarge:         "Passed Parameter too Large",
	PassedParameterTooSmall:         "Passed Parameter too Small",
	TooFewDataBytesReceived:         "Too Few Data Bytes Received",
	TransmitterSpecificCommandError: "Transmitter-Specific Command Error",
	InWriteProtectMode:              "In Write Protect Mode",
	WarningUpdateFailure:            "Warning: Update Failure",
	AccessRestricted:                "Access Restricted",
	Busy:                            "Busy",
	CommandNotImplemented:           "Command Not Implemented",
}

func (status CommandSpecificStatus) String() string {
	return commandSpecificStatusDescriptions[status]
}

//IsWarning CommandStatus inteface member
func (status CommandSpecificStatus) IsWarning() bool {
	return (status >= 24 && status <= 27) ||
		(status >= 96 && status <= 111) ||
		status == 8 || status == 14 || status == 30 || status == 31 ||
		(status >= 112 && status <= 127)
}

//IsError CommandStatus inteface member
func (status CommandSpecificStatus) IsError() bool {
	return (status > 0 && status <= 7) ||
		(status >= 16 && status <= 23) ||
		(status >= 32 && status <= 64) ||
		(status >= 9 && status <= 13) ||
		status == 15 || status == 28 || status == 29 ||
		(status >= 65 && status <= 95)
}
