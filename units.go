package hart

import "fmt"

// UnitCode contains HART unit codes enumeration
type UnitCode byte

// temperature
const (
	Celsius    UnitCode = 32
	Fahrenheit UnitCode = 33
	Rankine    UnitCode = 34
	Kelvin     UnitCode = 35
)

// Pressure
const (
	InH2O               UnitCode = 1
	InHg                UnitCode = 2
	FtH2O               UnitCode = 3
	mmH2O               UnitCode = 4
	mmHg                UnitCode = 5
	psi                 UnitCode = 6
	bar                 UnitCode = 7
	mbar                UnitCode = 8
	g_SqCm              UnitCode = 9
	kg_SqCm             UnitCode = 10
	Pa                  UnitCode = 11
	kPa                 UnitCode = 12
	torr                UnitCode = 13
	atm                 UnitCode = 14
	in_H2O_60_degrees_F UnitCode = 145
	m_H2O_4_degrees_C   UnitCode = 171
	MPa                 UnitCode = 237
	in_H2O_4_degrees_C  UnitCode = 238
	mm_H2O_4_degrees_C  UnitCode = 239
)

// desciptions for Stringer
var unitDescriptions = map[UnitCode]string{
	Celsius:    "°C",
	Fahrenheit: "°F",
	Rankine:    "°R",
	Kelvin:     "K",

	InH2O:               "in H₂O",
	InHg:                "in Hg",
	FtH2O:               "Ft H₂O",
	mmH2O:               "mm H₂O",
	mmHg:                "mm Hg",
	psi:                 "psi",
	bar:                 "bar",
	mbar:                "mbar",
	g_SqCm:              "G/cm²",
	kg_SqCm:             "kG/cm²",
	Pa:                  "Pa",
	kPa:                 "kPa",
	torr:                "torr",
	atm:                 "atm",
	in_H2O_60_degrees_F: "in H₂0 (60°F)",
	m_H2O_4_degrees_C:   "m H₂O (4°C)",
	MPa:                 "MPa",
	in_H2O_4_degrees_C:  "in H₂O (4°C)",
	mm_H2O_4_degrees_C:  "mm H₂O (4°C)",
}

// Stringer
func (unit UnitCode) String() string {
	if unitDescriptions[unit] == "" {
		return fmt.Sprintf("%d", unit)
	}
	return unitDescriptions[unit]
}
