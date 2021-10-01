package univrsl_test

import (
	"testing"

	"github.com/jszumigaj/hart/status"
	"github.com/jszumigaj/hart/univrsl"
)

func TestCommand0(t *testing.T) {

	device := univrsl.Device{}
	sut := univrsl.NewCommand0(&device)

	if sut.No() != 0 {
		t.Errorf("Unexpected No == %d", sut.No())
	}

	if sut.Description() != "Device identification" {
		t.Errorf("Unexpected description: %s", sut.Description())
	}

	if len(sut.Data()) > 0 {
		t.Errorf("Unexpected data: %s", sut.Data())
	}

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
