package hart

import "strings"

type PackedASCII []byte

type ASCII []byte

func NewPackedAscii(s string, packedLength int) PackedASCII {
	unpackedLength := packedLength * 8 / 6
	ascii := strings.ToUpper(s) + strings.Repeat(" ", unpackedLength-len(s))
	unpacked := []byte(ascii[:unpackedLength])

	var packed PackedASCII = make([]byte, packedLength)
	pack(unpacked, packed)
	return PackedASCII(packed)
}

func (packed PackedASCII) String() string {
	var unpackedBytes = len(packed) * 8 / 6
	var ascii = make([]byte, unpackedBytes)

	unpack(packed, ascii)
	return string(ascii)
}

func pack(ascii ASCII, packed PackedASCII) {
	var bits = len(ascii) * 6
	for i := 0; i < bits; i++ {
		bit := ascii.getASCIIBit(i)
		packed.setPackedBit(i, bit)
	}
}

func unpack(packed PackedASCII, ascii ASCII) {
	var bits = len(ascii) * 6
	for i := 0; i < bits; i++ {
		var bit = packed.getPackedBit(i)
		ascii.setASCIIBit(i, bit)
	}

	// set 6 bit (complement of bit 5)
	for i := 0; i < len(ascii); i++ {
		if (ascii[i] & 0x20) == 0 {
			ascii[i] |= 0x40
		}
	}

}

var shifts = []int{2, 3, 4, 5, 6, 7, 4, 5, 6, 7, 0, 1, 6, 7, 0, 1, 2, 3, 0, 1, 2, 3, 4, 5}
var tabIdxs = []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 2, 2, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2}

func (p PackedASCII) getPackedBit(index int) bool {
	var i = index % 24
	var shift = shifts[i]
	var tabIdx = tabIdxs[i] + ((index / 24) * 3)
	var b = (byte)((p[tabIdx] >> shift) & 0x01)
	return b == 0x01
}

func (p PackedASCII) setPackedBit(index int, bit bool) {
	var i = index % 24
	var shift = shifts[i]
	var tabIdx = tabIdxs[i] + ((index / 24) * 3)
	var mask = (byte)(1 << shift)
	if bit {
		p[tabIdx] |= mask
	} else {
		p[tabIdx] &= (byte)(mask ^ 0xff)
	}
}

func (ascii ASCII) getASCIIBit(index int) bool {
	tabIdx, shift := index/6, index%6
	b := byte((ascii[tabIdx] >> shift) & 0x01)
	return b == 0x01
}

func (ascii ASCII) setASCIIBit(index int, bit bool) {
	tabIdx, shift := index/6, index%6
	mask := byte(1 << shift)
	if bit {
		ascii[tabIdx] |= mask
	} else {
		ascii[tabIdx] &= byte(mask ^ 0xff)
	}
}
