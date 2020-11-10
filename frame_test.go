package hart_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/jszumigaj/hart"
	"github.com/jszumigaj/hart/mocks"
	"github.com/jszumigaj/hart/status"
)

func TestShortHartFrame(t *testing.T) {

	empty := make([]byte, 0)
	addr := make([]byte, 1)
	frame := hart.NewFrame(5, 0x02, addr, 0x00, empty, empty)

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
	if delimiter != hart.MasterToSlaveShortFrame {
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
	frame := hart.NewFrame(5, 0x06, addr, 0x00, status, data)

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
	if delimiter != hart.SlaveToMasterShortFrame {
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
	frame := hart.NewFrame(5, 0x82, addr, 0x00, empty, empty)

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
	if delimiter != hart.MasterToSlaveLongFrame {
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
	frame := hart.NewFrame(5, 0x86, addr, 0x00, status, data)

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
	if delimiter != hart.SlaveToMasterLongFrame {
		t.Errorf("Expected delimiter = 0x86 but got 0x%x", delimiter)
	}
	if dataLen != 0x0E {
		t.Errorf("Expected dataLen = 0x00 but got 0x%x", dataLen)
	}
	if crc != 0xa6 {
		t.Errorf("Expected crc = 0xa6 but got 0x%x", crc)
	}
}

func TestDeviceStatus(t *testing.T) {

	addr := []byte{0}
	resStat := []byte{0x00, 0x40}
	data := []byte{}
	frame := hart.NewFrame(5, 0x06, addr, 0x00, resStat, data)

	if frame.DeviceStatus() != status.ConfigurationChanged {
		t.Error("Expected ConfigurationChanged flag")
	}
}

func TestCommandStatusConfCh(t *testing.T) {

	addr := []byte{0}
	resStat := []byte{0x40, 0x00}
	data := []byte{}
	frame := hart.NewFrame(5, 0x06, addr, 0x00, resStat, data)

	if frame.CommandStatus() != status.CommandNotImplemented {
		t.Error("Expected ConfigurationChanged flag")
	}
}

func TestCommunicationsErrorStatus(t *testing.T) {

	addr := []byte{0}
	resStat := []byte{0x88, 0x00}
	data := []byte{}
	frame := hart.NewFrame(5, 0x06, addr, 0x00, resStat, data)

	if frame.CommandStatus() != status.LongitudalParityError {
		t.Errorf("Expected LongitudalParityError flag but got %v", frame.CommandStatus())
	}
}

func TestPrimaryAndSecondaryMaster(t *testing.T) {
	f := hart.FrameZero

	if f.IsPrimaryMaster() {
		t.Error("Expected secondary master here")
	}

	primary := f.AsPrimaryMaster()

	if primary.IsPrimaryMaster() == false {
		t.Error("Expected primary master here")
	}

	secondary := primary.AsSecondaryMaster()

	if secondary.IsPrimaryMaster() {
		t.Error("Expected primary master here")
	}
}

func TestLongFrameFactory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//make device stub
	device := mocks.NewMockDeviceIdentifier(ctrl)
	device.EXPECT().Id().Return(uint32(0x123456))
	device.EXPECT().ManufacturerId().Return(byte(0x55))
	device.EXPECT().MfrsDeviceType().Return(byte(0xAA))
	device.EXPECT().Preambles().Return(byte(5))

	//make command stub
	cmd := mocks.NewMockCommand(ctrl)
	cmd.EXPECT().No().Return(uint8(0))
	cmd.EXPECT().Data().Return(make([]byte, 0))

	fr := hart.LongFrameFactory(device, cmd)

	// expected frame: delim, ManId, DevType, Id, cmd, len
	expected := []byte{0x82, 0x55, 0xAA, 0x12, 0x34, 0x56, 0x00, 0x00}
	if cmp.Equal(fr.Buffer()[5:13], expected) == false {
		t.Errorf("Don't expected: %02x", fr.Buffer()[5:13])
	}
}

func TestDefaultFrameFactoryForCmd0(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//make device stub
	device := mocks.NewMockDeviceIdentifier(ctrl)
	device.EXPECT().Preambles().Return(byte(5))
	device.EXPECT().PollAddress().Return(byte(1)) // let's assume that Pool address is 1

	//make command stub
	cmd := mocks.NewMockCommand(ctrl)
	cmd.EXPECT().No().Return(uint8(0)) // Command 0 should produce short command
	cmd.EXPECT().No().MaxTimes(2)      // we expected two calls
	cmd.EXPECT().Data().Return(make([]byte, 0))

	fr := hart.DefaultFrameFactory(device, cmd)

	// expected frame: delim, pollAddr, cmd, len
	expected := []byte{0x02, 0x01, 0x00, 0x00}
	if cmp.Equal(fr.Buffer()[5:9], expected) == false {
		t.Errorf("Don't expected: %02x", fr.Buffer()[5:13])
	}
}

func TestDefaultFrameFactoryForCmd1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	//make device stub
	device := mocks.NewMockDeviceIdentifier(ctrl)
	device.EXPECT().Id().Return(uint32(0x123456))
	device.EXPECT().ManufacturerId().Return(byte(0x55))
	device.EXPECT().MfrsDeviceType().Return(byte(0xAA))
	device.EXPECT().Preambles().Return(byte(5))

	//make command stub
	cmd := mocks.NewMockCommand(ctrl)
	cmd.EXPECT().No().MaxTimes(2).Return(uint8(1)) // Command 1 should produce long command
	cmd.EXPECT().Data().Return(make([]byte, 0))

	fr := hart.DefaultFrameFactory(device, cmd)

	// expected frame: delim, ManId, DevType, Id, cmd, len
	expected := []byte{0x82, 0x55, 0xAA, 0x12, 0x34, 0x56, 0x01, 0x00}
	if cmp.Equal(fr.Buffer()[5:13], expected) == false {
		t.Errorf("Don't expected: %02x", fr.Buffer()[5:13])
	}
}
