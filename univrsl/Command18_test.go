package univrsl_test

import (
	"strings"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"

	"github.com/jszumigaj/hart/status"
	"github.com/jszumigaj/hart/univrsl"
)

func TestCommand18(t *testing.T) {

	device := &univrsl.Device{}
	date, _ := time.Parse("2006-01-02", "2020-11-12")
	sut := univrsl.NewCommand18(device, "tag", "descriptor", date)

	if sut.No() != 18 {
		t.Errorf("Unexpected No == %d", sut.No())
	}

	if sut.Description() != "Write tag, descriptor and date" {
		t.Errorf("Unexpected description: %s", sut.Description())
	}

	expectedData := []byte{0x50, 0x11, 0xE0, 0x82, 0x08, 0x20, 0x10, 0x54, 0xC3, 0x48, 0x94, 0x14, 0x3D, 0x28,
		0x20, 0x82, 0x08, 0x20, 12, 11, 120}
	if !cmp.Equal(sut.Data(), expectedData) {
		t.Errorf("Unexpected data: %02x", sut.Data())
	}
}

func TestCommand18SetData(t *testing.T) {

	sut := univrsl.Command18{}

	if sut.No() != 18 {
		t.Errorf("Unexpected No == %d", sut.No())
	}

	data := []byte{0x50, 0x11, 0xE0, 0x82, 0x08, 0x20, 0x10, 0x54, 0xC3, 0x48, 0x94, 0x14, 0x3D, 0x28,
		0x20, 0x82, 0x08, 0x20, 12, 11, 120}

	sut.SetData(data, status.NoCommandSpecificError)

	if strings.TrimSpace(sut.Tag) != "TAG" {
		t.Errorf("Unexpected Tag: %s", sut.Tag)
	}
	if strings.TrimSpace(sut.Descriptor) != "DESCRIPTOR" {
		t.Errorf("Unexpected Descriptor: %s", sut.Descriptor)
	}

	date, _ := time.Parse("2006-01-02", "2020-11-12")
	if sut.Date != date {
		t.Errorf("Unexpected Date: %v", sut.Date)
	}
}
