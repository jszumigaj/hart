package univrsl_test

import (
	"strings"
	"testing"

	"github.com/jszumigaj/hart/status"
	"github.com/jszumigaj/hart/univrsl"
)

func TestCommand12SetData(t *testing.T) {

	sut := univrsl.Command12{}

	if sut.No() != 12 {
		t.Errorf("Unexpected No == %d", sut.No())
	}

	if sut.Description() != "Read message" {
		t.Errorf("Unexpected description: %s", sut.Description())
	}

	if len(sut.Data()) > 0 {
		t.Errorf("Unexpected data: %s", sut.Data())
	}

	data := []byte{0x50, 0x54, 0xD4, 0x80, 0xD1, 0x53, 0x4C, 0x11, 0xC5, 0x82, 0x08, 0x20, 0x82, 0x08,
		0x20, 0x82, 0x08, 0x20, 0x82, 0x08, 0x20, 0x82, 0x08, 0x20}

	sut.SetData(data, status.NoCommandSpecificError)

	if strings.TrimSpace(sut.Message) != "TEST MESSAGE" {
		t.Errorf("Unexpected Tag: %s", sut.Message)
	}
}
