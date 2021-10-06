package univrsl_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jszumigaj/hart"
	"github.com/jszumigaj/hart/status"
	"github.com/jszumigaj/hart/univrsl"
)

func TestCommand11(t *testing.T) {

	device := &univrsl.Device{}
	sut := univrsl.NewCommand11(device, "MyTag")

	if sut.No() != 11 {
		t.Errorf("Unexpected No == %d", sut.No())
	}

	if sut.Description() != "Read unique identifier associated with TAG" {
		t.Errorf("Unexpected description: %s", sut.Description())
	}

	expectedData := []byte{0x35, 0x95, 0x01, 0x1e, 0x08, 0x20}
	if !cmp.Equal(sut.Data(), expectedData) {
		t.Errorf("Unexpected data: 0x%02x", sut.Data())
	}
}

func TestCommand11SetData(t *testing.T) {

	device := univrsl.Device{}
	sut := univrsl.NewCommand11(&device, "tag")

	data := []byte{0xFE, 0xBC, 0x7B, 0x05, 0x05, 0x01, 0x01, 0x09, 0x01, 0x9D, 0x34, 0xB3, 0x05, 0x08, 0x00, 0x7C, 0x00, 0x00, 0xBC, 0x00, 0xBC, 0x01}

	sut.SetData(data, status.NoCommandSpecificError)

	if device.ManufacturerId() != 0xbc {
		t.Errorf("Unexpected ManId: %x", device.ManufacturerId())
	}

	if device.MfrsDeviceType() != 0x7b {
		t.Errorf("Unexpected DevType: %x", device.MfrsDeviceType())
	}

	if device.Id() != 10302643 {
		t.Errorf("Unexpected DevId: %x", device.Id())
	}

	if device.Preambles() != 5 {
		t.Errorf("Unexpected preambles no: %d", device.Preambles())
	}

	if device.HartProtocolMajorRevision != 5 {
		t.Errorf("Unexpected HartProtocolMajorRevision: %d", device.HartProtocolMajorRevision)
	}

	if device.RevisionLevel != 1 {
		t.Errorf("Unexpected RevisionLevel: %d", device.RevisionLevel)
	}

	if device.SoftwareRevisionLevel != 1 {
		t.Errorf("Unexpected SoftwareRevisionLevel: %d", device.SoftwareRevisionLevel)
	}

	if device.HardwareRevisionLevel != 1 {
		t.Errorf("Unexpected HardwareRevisionLevel: %d", device.HardwareRevisionLevel)
	}

	if device.PhisicalSignalingCode != 1 {
		t.Errorf("Unexpected PhisicalSignalingCode: %d", device.PhisicalSignalingCode)
	}

	if device.Flags != 1 {
		t.Errorf("Unexpected Flags: %d", device.Flags)
	}
}

func TestCheckBoradcastAddress(t *testing.T) {
	device := &univrsl.Device{}
	device.MfrsId = 0xbc
	device.DevType = 0x7b
	device.DevId = 0x123456

	// default useBroadcastAddres = true
	sut := univrsl.NewCommand11(device, "TAG")

	frame := hart.LongFrameFactory(sut)
	tx := frame.Buffer()
	addr := tx[6:11]
	for i, v := range addr {
		if v != 0 {
			t.Errorf("Expected all zeros in addr, but get addr[%v]=%v", i, v)
		}
	}
}

func TestCheckLongAddress(t *testing.T) {
	device := &univrsl.Device{}
	device.MfrsId = 0xbc
	device.DevType = 0x7b
	device.DevId = 0x123456

	// useBroadcastAddres = false
	sut := univrsl.NewCommand11(device, "TAG", false)

	frame := hart.LongFrameFactory(sut)
	tx := frame.Buffer()
	addr := tx[6:11]
	for i, v := range addr {
		if v == 0 {
			t.Errorf("Not expected zeros in addr, but get addr[%v]=%v", i, v)
		}
	}
}
