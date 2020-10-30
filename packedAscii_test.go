package hart

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAsciiToPackedAscii(t *testing.T) {
	//arrange
	msg := "test"
	expected := []byte{80, 84, 212, 130, 8, 32, 130, 0}

	//act
	packedA := NewPackedAscii(msg, 8)
	result := []byte(packedA)
	
	//assert
	if !cmp.Equal(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestPackedAsciiToAscii(t *testing.T) {

	packed := PackedASCII([]byte{80, 84, 212, 130, 8, 32, 130, 0})
	expected := "TEST"

	result := packed.String()
	result = strings.TrimRight(result, " ")

    if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}
