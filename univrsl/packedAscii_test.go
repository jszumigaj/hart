package univrsl

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestAsciiToPackedAscii(t *testing.T) {
	//arrange
	msg := "test"
	expected := []byte{80, 84, 212, 130, 8, 32}

	//act
	packed := NewPackedASCII(msg, 6)
	result := []byte(packed)

	//assert
	if !cmp.Equal(result, expected) {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestPackedAsciiToAscii(t *testing.T) {

	packed := PackedASCII([]byte{80, 84, 212, 130, 8, 32})
	expected := "TEST"

	result := packed.String()
	result = strings.TrimRight(result, " ")

	if result != expected {
		t.Errorf("Expected %q but got %q", expected, result)
	}
}
