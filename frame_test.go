package hart

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestShortHartFrame(t *testing.T) {

	empty := make([]byte, 0)
	addr := make([]byte, 1)
	frame := NewFrame(5, 0x02, addr, 0x00, empty, empty)

	length := frame.Length()
	if length != 10 {
		t.Errorf("Expected 10 got %d", length)
	}

	buf := frame.Buffer()

	preambles := buf[:5]
	delimiter := buf[5]
	dataLen := buf[8]
	crc := buf[9]

	if preambles[0] != 0xff || preambles[1] != 0xff || preambles[2] != 0xff || preambles[3] != 0xff || preambles[4] != 0xff {
		t.Errorf("Expected 0xff preambles but got %q", preambles)
	}
	if delimiter != MasterToSlaveShortFrame {
		t.Errorf("Expected delimiter = 0x02 but got %x", delimiter)
	}
	if dataLen != 0x00 {
		t.Errorf("Expected dataLen = 0x00 but got %x", dataLen)
	}
	if crc != 0x02 {
		t.Errorf("Expected crc = 0x02 but got %x", crc)
	}
}

func TestShortReplyHartFrame(t *testing.T) {

	addr := []byte{0}
	status := []byte{0x00, 0x40}
	data := []byte{0xFE, 0xBC, 0x7B, 0x05, 0x05, 0x03, 0x02, 0x10, 0x01, 0x12, 0x31, 0xE1}
	frame := NewFrame(5, 0x06, addr, 0x00, status, data)

	length := frame.Length()
	if length != 24 {
		t.Errorf("Expected 24 got %d", length)
	}

	buf := frame.Buffer()

	preambles := buf[:5]
	if preambles[0] != 0xff || preambles[1] != 0xff || preambles[2] != 0xff || preambles[3] != 0xff || preambles[4] != 0xff {
		t.Errorf("Expected 0xff preambles but got %q", preambles)
	}

	delimiter := buf[5]
	if delimiter != SlaveToMasterShortFrame {
		t.Errorf("Expected delimiter = 0x02 but got %x", delimiter)
	}

	dataLen := buf[8]
	if dataLen != 0x0e {
		t.Errorf("Expected dataLen = 0x0E but got %x", dataLen)
	}

	data2 := buf[11:]
	for i := range data {
		if data2[i] != data[i] {
			t.Errorf("Expected data[%d] = %x but got %x", i, data2[i], data[i])

		}
	}
	crc := buf[len(buf)-1]
	if crc != 0xa3 {
		t.Errorf("Expected crc = 0xA3 but got %x", crc)
	}
}

func TestLongHartFrame(t *testing.T) {

	empty := []byte{}
	addr := []byte{0xBC, 0x7B, 0x12, 0x31, 0xE1}
	frame := NewFrame(5, 0x82, addr, 0x00, empty, empty)

	length := frame.Length()
	if length != 14 {
		t.Errorf("Expected frame.Length() = 14 got %d", length)
	}

	buf := frame.Buffer()

	//fmt.Println(buf)

	preambles := buf[:5]
	delimiter := buf[5]
	dataLen := buf[12]
	crc := buf[13]

	if preambles[0] != 0xff || preambles[1] != 0xff || preambles[2] != 0xff || preambles[3] != 0xff || preambles[4] != 0xff {
		t.Errorf("Expected 0xff preambles but got %q", preambles)
	}
	if delimiter != MasterToSlaveLongFrame {
		t.Errorf("Expected delimiter = 0x82 but got %x", delimiter)
	}
	if dataLen != 0x00 {
		t.Errorf("Expected dataLen = 0x00 but got %x", dataLen)
	}
	if crc != 0x07 {
		t.Errorf("Expected crc = 0x07 but got %x", crc)
	}
}

func TestLongReplyHartFrame(t *testing.T) {

	addr := []byte{0xBC, 0x7B, 0x12, 0x31, 0xE1}
	status := []byte{0x00, 0x40}
	data := []byte{0xFE, 0xBC, 0x7B, 0x05, 0x05, 0x03, 0x02, 0x10, 0x01, 0x12, 0x31, 0xE1}
	frame := NewFrame(5, 0x86, addr, 0x00, status, data)

	length := frame.Length()
	if length != 28 {
		t.Errorf("Expected frame.Length() = 14 got %d", length)
	}

	buf := frame.Buffer()

	//fmt.Println(buf)

	preambles := buf[:5]
	delimiter := buf[5]
	dataLen := buf[12]
	crc := buf[27]

	if preambles[0] != 0xff || preambles[1] != 0xff || preambles[2] != 0xff || preambles[3] != 0xff || preambles[4] != 0xff {
		t.Errorf("Expected 0xff preambles but got %q", preambles)
	}
	if delimiter != SlaveToMasterLongFrame {
		t.Errorf("Expected delimiter = 0x86 but got 0x%x", delimiter)
	}
	if dataLen != 0x0E {
		t.Errorf("Expected dataLen = 0x00 but got 0x%x", dataLen)
	}
	if crc != 0xa6 {
		t.Errorf("Expected crc = 0xa6 but got 0x%x", crc)
	}
}

func TestFrameZero(t *testing.T){
	frameZero := FrameZero

	if cmp.Equal(frameZero.address, []byte{0}) == false {
		t.Error("Invalid address")
	}

	if frameZero.bytesCount() != 0 {
		t.Error("Invalid bytes count")
	}
}


func TestDeviceStatus(t *testing.T) {

	addr := []byte{0}
	status := []byte{0x00, 0x40}
	data := []byte{}
	frame := NewFrame(5, 0x06, addr, 0x00, status, data)

	if frame.DeviceStatus() != ConfigurationChanged {
		t.Error("Expected ConfigurationChanged flag")
	}
}

func TestCommandStatus(t *testing.T) {

	addr := []byte{0}
	status := []byte{0x40, 0x00}
	data := []byte{}
	frame := NewFrame(5, 0x06, addr, 0x00, status, data)

	if frame.CommandStatus() != CommandNotImplemented {
		t.Error("Expected ConfigurationChanged flag")
	}
}

func TestCommunicationsErrorStatus(t *testing.T) {

	addr := []byte{0}
	status := []byte{0x88, 0x00}
	data := []byte{}
	frame := NewFrame(5, 0x06, addr, 0x00, status, data)

	if frame.CommandStatus() != LongitudalParityError {
		t.Errorf("Expected LongitudalParityError flag but got %v", frame.CommandStatus())
	}
}