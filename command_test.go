package hart

import (
	"encoding/binary"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAddr(t *testing.T) {

	ID := uint32(1192417)
	buf := make([]byte, 4)

	binary.BigEndian.PutUint32(buf, ID)
	buf = buf[1:]

	if !cmp.Equal(buf, []byte{0x12, 0x31, 0xE1}) {
		t.Errorf("Got %x", buf)
	}
}
