package univrsl

import "fmt"

// UnitCode contains HART unit codes enumeration
type UnitCode byte

// Temperature units
const (
	Celsius    UnitCode = 32
	Fahrenheit UnitCode = 33
	Rankine    UnitCode = 34
	Kelvin     UnitCode = 35
)

// Pressure units
const (
	InH2O               UnitCode = 1
	InHg                UnitCode = 2
	FtH2O               UnitCode = 3
	MmH2O               UnitCode = 4
	MmHg                UnitCode = 5
	Psi                 UnitCode = 6
	Bar                 UnitCode = 7
	Mbar                UnitCode = 8
	G_SqCm              UnitCode = 9
	Kg_SqCm             UnitCode = 10
	Pa                  UnitCode = 11
	KPa                 UnitCode = 12
	Torr                UnitCode = 13
	Atm                 UnitCode = 14
	In_H2O_60_degrees_F UnitCode = 145
	M_H2O_4_degrees_C   UnitCode = 171
	MPa                 UnitCode = 237
	In_H2O_4_degrees_C  UnitCode = 238
	Mm_H2O_4_degrees_C  UnitCode = 239
)

// desciptions for Stringer
var unitDescriptions = map[UnitCode]string{
	Celsius:             "°C",
	Fahrenheit:          "°F",
	Rankine:             "°R",
	Kelvin:              "K",
	InH2O:               "in H₂O",
	InHg:                "in Hg",
	FtH2O:               "Ft H₂O",
	MmH2O:               "mm H₂O",
	MmHg:                "mm Hg",
	Psi:                 "psi",
	Bar:                 "bar",
	Mbar:                "mbar",
	G_SqCm:              "G/cm²",
	Kg_SqCm:             "kG/cm²",
	Pa:                  "Pa",
	KPa:                 "kPa",
	Torr:                "torr",
	Atm:                 "atm",
	In_H2O_60_degrees_F: "in H₂0 (60°F)",
	M_H2O_4_degrees_C:   "m H₂O (4°C)",
	MPa:                 "MPa",
	In_H2O_4_degrees_C:  "in H₂O (4°C)",
	Mm_H2O_4_degrees_C:  "mm H₂O (4°C)",
}

// Stringer
func (unit UnitCode) String() string {
	if unitDescriptions[unit] == "" {
		return fmt.Sprintf("%d", unit)
	}
	return unitDescriptions[unit]
}
