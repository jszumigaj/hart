package univrsl_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/jszumigaj/hart/status"
	"github.com/jszumigaj/hart/univrsl"
)

func TestCommand17(t *testing.T) {

	device := &univrsl.Device{}
	sut := univrsl.NewCommand17(device, "test message")

	if sut.No() != 17 {
		t.Errorf("Unexpected No == %d", sut.No())
	}

	if sut.Description() != "Write message" {
		t.Errorf("Unexpected description: %s", sut.Description())
	}

	expectedData := []byte{0x50, 0x54, 0xD4, 0x80, 0xD1, 0x53, 0x4C, 0x11, 0xC5, 0x82, 0x08, 0x20, 0x82, 0x08,
		0x20, 0x82, 0x08, 0x20, 0x82, 0x08, 0x20, 0x82, 0x08, 0x20}
	if !cmp.Equal(sut.Data(), expectedData) {
		t.Errorf("Unexpected data: %02x", sut.Data())
	}
}

func TestCommand17SetData(t *testing.T) {

	sut := univrsl.Command17{}

	data := []byte{0x50, 0x54, 0xD4, 0x80, 0xD1, 0x53, 0x4C, 0x11, 0xC5, 0x82, 0x08, 0x20, 0x82, 0x08,
		0x20, 0x82, 0x08, 0x20, 0x82, 0x08, 0x20, 0x82, 0x08, 0x20}

	sut.SetData(data, status.NoCommandSpecificError)

	if strings.TrimSpace(sut.Message) != "TEST MESSAGE" {
		t.Errorf("Unexpected Tag: %s", sut.Message)
	}
}
