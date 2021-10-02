package univrsl

import (
	"github.com/jszumigaj/hart"
)

type AlarmSelectionCode byte
type TransferFunctionCode byte
type WriteProtectCode byte
type AnalogChannelFlags byte

// Command15 reads Device Output Information
type Command15 struct {
	commandBase

	// command data fields
	AlarmSelectionCode           AlarmSelectionCode   `json:"alarmSelectionCodes"`           // PV Alarm Selection Code, 8-bit unsigned integer, Refer to Table VI; Alarm Selection Codes
	TransferFunctionCode         TransferFunctionCode `json:"transferFunctionCodes"`         // PV Transfer Function Code, 8-bit unsigned integer, Refer to Table III; Transfer Function Codes
	UpperAndLowerRangeValuesUnit UnitCode             `json:"upperAndLowerRangeValuesUnits"` // PV Upeer and Lower Range Values Units Code, 8-bit unsigned integer, Refer to Table II; Unit Codes
	UpperRangeValue              float32              `json:"upperRangeValue"`               // Primary Variable Upper Range Value, IEEE 754
	LowerRangeValue              float32              `json:"lowerRangeValue"`               // Primary Variable Lower Range Value, IEEE 754
	Damping                      float32              `json:"damping"`                       // Primary Variable Damping Value, IEEE 754, Units of seconds [s]
	WriteProtecCode              WriteProtectCode     `json:"writeProtecCode"`               // Write Protec Code, 8-bit unsigned integer, Refer to Table VII; Write Protec Codes
	PrivateLabelDistributorCode  byte                 `json:"privateLabelDistributorCode"`   // HART5: Private Label Distributor Code, 8-bit unsigned integer, Refer to Table VIII; Manufacturer Identification Codes; HART7: Not used, must be set to 250
	notUsed                      byte                 `json:""`                              // Must be set to 250 by device.
	AnalogChannel                AnalogChannelFlags   `json:"analogChannel"`                 // PV Analog Channel Flags (pojawia się w HART7)
}

func NewCommand15(device hart.DeviceIdentifier) *Command15 { 
	return &Command15{commandBase: commandBase{device: device}} 
}

// Description property
func (c *Command15) Description() string { return "Read Device Output Information" }

// No property
func (c *Command15) No() byte { return 15 }

// SetData parse received data
func (c *Command15) SetData(data []byte, status hart.CommandStatus) bool {
	c.status = status

	if len(data) < 17 {
		return false
	}

	var val float32
	var ok bool

	c.AlarmSelectionCode = AlarmSelectionCode(data[0])
	c.TransferFunctionCode = TransferFunctionCode(data[1])
	c.UpperAndLowerRangeValuesUnit = UnitCode(data[2])
	if val, ok := getFloat(data[3:]); ok {
		c.UpperRangeValue = val
	}
	if val, ok = getFloat(data[7:]); ok {
		c.LowerRangeValue = val
	}
	if val, ok = getFloat(data[11:]); ok {
		c.Damping = val
	} else {
		return false
	}

	c.WriteProtecCode = WriteProtectCode(data[15])
	c.notUsed = data[16]

	// HART7:
	if len(data) >= 18 {
		c.AnalogChannel = AnalogChannelFlags(data[17])
	}

	return true
}
