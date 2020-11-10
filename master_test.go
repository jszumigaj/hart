package hart_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/jszumigaj/hart"
	"github.com/jszumigaj/hart/mocks"
	"github.com/jszumigaj/hart/status"
	"github.com/jszumigaj/hart/univrsl"

	"github.com/golang/mock/gomock"
)

func TestEverythingOkey(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// configure mock modem
	modem := mocks.NewMockFrameSender(ctrl)
	// expected Tx frame
	var expTx = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x00, 0x00, 0x00, 0x02}
	// example response frame
	resRx := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x06, 0x00, 0x00, 0x0E, 0x00, 0x40,
		0xFE, 0xBC, 0x7B, 0x05, 0x05, 0x03, 0x02, 0x10, 0x01, 0x12, 0x31, 0xE1, 0xA3}
	modem.EXPECT().SendFrame(expTx, gomock.Any()).DoAndReturn(
		func(tx, rx []byte) (int, error) {
			copy(rx, resRx)
			return len(resRx), nil
		})

	dev := &univrsl.Device{}
	command := &univrsl.Command0{Device: dev}

	//ACT
	sut := hart.NewMaster(modem)
	result, _ := sut.Execute(command, dev)

	// ASSERT
	AssertEqual(t, result, status.NoCommandSpecificError)
	AssertEqual(t, dev.ManufacturerId(), byte(0xbc))
	AssertEqual(t, dev.MfrsDeviceType(), byte(0x7b))
	AssertEqual(t, dev.Id(), uint32(1192417))
	AssertEqual(t, dev.Status(), status.ConfigurationChanged)
	AssertEqual(t, dev.Preambles(), byte(5))
}

func TestWorkingAsPrimaryMaster(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// configure mock modem
	modem := mocks.NewMockFrameSender(ctrl)
	// expected Tx frame (0x80 bit in pollAddress shoudl be set)
	var expTx = []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0x02, 0x80, 0x00, 0x00, 0x82}
	// example response frame
	resRx := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x06, 0x80, 0x00, 0x0E, 0x00, 0x40,
		0xFE, 0xBC, 0x7B, 0x05, 0x05, 0x03, 0x02, 0x10, 0x01, 0x12, 0x31, 0xE1, 0x23}
	modem.EXPECT().SendFrame(expTx, gomock.Any()).DoAndReturn(
		func(tx, rx []byte) (int, error) {
			copy(rx, resRx)
			return len(resRx), nil
		})

	dev := &univrsl.Device{}
	command := &univrsl.Command0{Device: dev}

	//ACT
	sut := hart.NewMaster(modem)
	sut.Primary = true
	result, _ := sut.Execute(command, dev)

	// ASSERT
	AssertEqual(t, result, status.NoCommandSpecificError)
}

func TestCommandStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// configure mock modem
	modem := mocks.NewMockFrameSender(ctrl)
	// example response frame with CommandNotImplemented command status and ColdStart device status
	resRx := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x86, 0xBC, 0x7B, 0x12, 0x31, 0xE1, 0x55, 0x02, 0x40, 0x20, 0xB4}
	modem.EXPECT().SendFrame(gomock.Any(), gomock.Any()).DoAndReturn(
		func(tx, rx []byte) (int, error) {
			copy(rx, resRx)
			return len(resRx), nil
		})

	dev := &univrsl.Device{}
	command := dev.Command0()

	//Act
	sut := hart.NewMaster(modem)
	result, _ := sut.Execute(command, dev)

	AssertEqual(t, result, status.CommandNotImplemented)
	AssertEqual(t, dev.Status(), status.ColdStart)
}

// CommunicationStatus tests checks master Execute behavior in case device returns LongitudalParityError communication err.
func TestCommunicationStatus(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// configure mock modem
	modem := mocks.NewMockFrameSender(ctrl)
	// example response frame - 0x88 means LongitudalParityError
	resRx := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x06, 0x00, 0x00, 0x02, 0x88, 0x00, 0x8C}
	modem.EXPECT().SendFrame(gomock.Any(), gomock.Any()).DoAndReturn(
		func(tx, rx []byte) (int, error) {
			copy(rx, resRx)
			return len(resRx), nil
		})

	dev := &univrsl.Device{}
	command := dev.Command0()

	//Act
	sut := hart.NewMaster(modem)
	result, err := sut.Execute(command, dev)

	// we expect that Execute returns comm status error in both: status and error result. Device status should be zero.
	AssertEqual(t, result, status.LongitudalParityError)
	AssertEqual(t, err, status.LongitudalParityError)
	AssertEqual(t, dev.Status(), status.OK)
}

func TestNoResponse(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// configure mock (not responding) modem
	modem := mocks.NewMockFrameSender(ctrl)
	modem.EXPECT().SendFrame(gomock.Any(), gomock.Any()).Return(0, nil)

	dev := &univrsl.Device{}
	command := dev.Command0()

	//Act
	sut := hart.NewMaster(modem)
	result, err := sut.Execute(command, dev)

	AssertEqual(t, err, status.ErrNoResponse)
	AssertEqual(t, result, nil)
}

func TestFrameParsingError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// configure mock modem
	modem := mocks.NewMockFrameSender(ctrl)
	// example response frame with bad crc
	resRx := []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x06, 0x00, 0x00, 0x02, 0x00, 0x00, 0x00}
	modem.EXPECT().SendFrame(gomock.Any(), gomock.Any()).DoAndReturn(
		func(tx, rx []byte) (int, error) {
			copy(rx, resRx)
			return len(resRx), nil
		})

	dev := &univrsl.Device{}
	command := dev.Command0()

	//Act
	sut := hart.NewMaster(modem)
	result, err := sut.Execute(command, dev)

	//Assert
	if parsErr, ok := err.(*status.FrameParsingError); !ok {
		t.Errorf("Expected FrameParsingError but got %T", err)
	} else {
		AssertEqual(t, parsErr.Frame[:len(resRx)], resRx)
	}

	AssertEqual(t, result, nil)
}

func TestErrorInSender(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// configure mock modem which return "Some problem" error
	modem := mocks.NewMockFrameSender(ctrl)
	modem.EXPECT().SendFrame(gomock.Any(), gomock.Any()).Return(0, errors.New("Some problem"))

	dev := &univrsl.Device{}
	command := dev.Command0()

	//Act
	sut := hart.NewMaster(modem)
	result, err := sut.Execute(command, dev)

	AssertEqual(t, err.Error(), "Some problem")
	AssertEqual(t, result, nil)
}

func AssertEqual(t *testing.T, result, expected interface{}) {
	if resBuf, ok := result.([]byte); ok {
		if expBuf, ok2 := expected.([]byte); ok2 {
			if !cmp.Equal(resBuf, expBuf) {
				t.Errorf("Expected %02x but got %02x", resBuf, expBuf)
			}
			return
		} else {
			t.Errorf("Can't compare %T with %T", result, expected)
			return
		}
	}
	if result != expected {
		t.Errorf("Expected %q (%T) but got %q (%T)", expected, expected, result, result)
	}
}
