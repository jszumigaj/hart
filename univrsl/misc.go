package univrsl

import (
	"encoding/binary"
	"math"
)

func getFloat(buf []byte) (float32, bool) {
	if len(buf) < 4 {
		return float32(math.NaN()), false
	}

	bits := binary.BigEndian.Uint32(buf[0:4])
	return math.Float32frombits(bits), true
}

func getFloatWithUnit(buf []byte) (float32, UnitCode, bool) {
	if len(buf) < 5 {
		return float32(math.NaN()), 0, false
	}

	unit := UnitCode(buf[0])
	bits := binary.BigEndian.Uint32(buf[1:5])
	return math.Float32frombits(bits), unit, true
}
