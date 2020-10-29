package hart

// CommandStatus is the interface wrapped CommunicationsErrorSummaryFlags and CommandSpecificStatus as one common status
// The returned type depends of MSB bit of the first frame status byte.
// If MSB=1 this byte means CommunicationsError otherwise it is CommandStatus
// Some commands can return individual command specific status as CommandStatus type
type CommandStatus interface {
	IsError() bool
	IsWarning() bool
}
