package hart

import "strings"

// FieldDeviceStatus provides status of device
type FieldDeviceStatus byte

// FieldDeviceStatus flags
const (
	// A hardware error or failure has been detected by the device.
	FieldDeviceMalfunction FieldDeviceStatus = 0x80

	// A write or set command has been executed.
	ConfigurationChanged FieldDeviceStatus = 0x40

	// Power has been removed and reapplied resulting in the reinstallations of the setup information.
	/// The first command to recognize this condition will automatically reset this flag.
	ColdStart FieldDeviceStatus = 0x20

	// More status information is available than can be returned in the Field Device Status command #48.
	MoreStatusAvailable FieldDeviceStatus = 0x10

	// The analog and digital analog outputs for the Primary Variable are held at the requested value.
	/// They will not respond to the applied process.
	PrimaryVariableAnalogOutputFixed FieldDeviceStatus = 0x08

	// The analog and digital analog outputs fir the Primary Variable are beyond their limits and no
	// longer represent the true applied process.
	PrimaryVariableAnalogOutputSaturated FieldDeviceStatus = 0x04

	// The process applied to a sensor other than that of the PRimary Variable, is beyond
	// the operating limits of the device.
	NonPrimaryVariableOutOfLimits FieldDeviceStatus = 0x02

	// Primary Variable Out of Limits - The process applied to the sensor for the Primary Variable
	// is beyond the operating limits of the device
	PrimaryVariableOutOfLimits FieldDeviceStatus = 0x01
)

var fieldDeviceStatusDescriptions = map[FieldDeviceStatus]string{
	FieldDeviceMalfunction:               "Field device malfunction",
	ConfigurationChanged:                 "Configuration changed",
	ColdStart:                            "Cold start",
	MoreStatusAvailable:                  "More status available",
	PrimaryVariableAnalogOutputFixed:     "Analog output current fixed",
	PrimaryVariableAnalogOutputSaturated: "Analog output saturated",
	NonPrimaryVariableOutOfLimits:        "Variable (not primary) out of limits",
	PrimaryVariableOutOfLimits:           "Primary variable out of limits",
}

func (status FieldDeviceStatus) String() string {
	names := []string{}
	for i := 0; i < 8; i++ {
		mask := FieldDeviceStatus(1 << i)
		if status.HasFlag(mask) {
			names = append(names, fieldDeviceStatusDescriptions[mask])
		}
	}

	return strings.Join(names, ", ")
}

// HasFlag f
func (status FieldDeviceStatus) HasFlag(mask FieldDeviceStatus) bool {
	return status&mask == mask
}
