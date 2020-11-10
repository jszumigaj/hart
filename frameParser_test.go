package hart

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFrameParser(t *testing.T) {

	buffer := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x86, 0xBC, 0x7B, 0x9D, 0x34, 0xB3, 0x00, 0x0E, 0x00, 0x00,
		0xFE, 0xBC, 0x7B, 0x05, 0x05, 0x01, 0x01, 0x08, 0x01, 0x9D, 0x34, 0xB3, 0x7F}

	frame, ok := Parse(buffer)

	if ok == false {
		t.Fatal("Parse error")
	}

	if frame.delimiter != 0x86 {
		t.Errorf("Expected delimiter 0x86 but got 0x%x", frame.delimiter)
	}

	if !cmp.Equal(frame.data, buffer[15:27]) {
		t.Error("frame.Data and data in buffer are dif")
	}
}

func TestParseNotCompletedFrame(t *testing.T) {
	buffer := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x86, 0xBC, 0x7B, 0x9D, 0x34, 0xB3, 0x00, 0x0E, 0x00, 0x00,
		0xFE, 0xBC, 0x7B, 0x05, 0x05, 0x01, 0x01, 0x08, 0x01, 0x9D, 0x34}

	frame, ok := Parse(buffer)

	if ok != false {
		t.Error("Expected false but is true")
	}

	if frame != nil {
		t.Error("Expected nil")
	}

}

func TestParseBadCrcFrame(t *testing.T) {
	buffer := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x86, 0xBC, 0x7B, 0x9D, 0x34, 0xB3, 0x00, 0x0E, 0x00, 0x00,
		0xFE, 0xBC, 0x7B, 0x05, 0x05, 0x01, 0x01, 0x08, 0x01, 0x9D, 0x34, 0xB3, 0xFF}

	// parsing buffer with bad CRC returns frame but ok is false!
	frame, ok := Parse(buffer)

	if ok != false {
		t.Error("Expected false but is true")
	}

	if frame == nil {
		t.Error("Expected frame with bad crc")
	}

	if !cmp.Equal(frame.data, buffer[15:27]) {
		t.Error("frame.Data and data in buffer are dif")
	}
}

func TestFrameZero(t *testing.T) {
	frameZero := FrameZero

	if cmp.Equal(frameZero.address, []byte{0}) == false {
		t.Error("Invalid address")
	}

	if frameZero.bytesCount() != 0 {
		t.Error("Invalid bytes count")
	}
}
