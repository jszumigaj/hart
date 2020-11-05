package hart_test

import (
	"errors"
	"testing"

	"github.com/jszumigaj/hart"
	"github.com/jszumigaj/hart/univrsl"
	"github.com/jszumigaj/hart/mocks"
	"github.com/jszumigaj/hart/status"

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
	result, _ := sut.Execute(command, dev)

	AssertEqual(t, result, status.LongitudalParityError)
	AssertEqual(t, dev.Status(), status.FieldDeviceStatus(0))
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
	if result != expected {
		t.Errorf("Expected %q (%T) but got %q (%T)", expected, expected, result, result)
	}
}
