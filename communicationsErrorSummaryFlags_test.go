package hart

import "testing"

func TestOneFlag(t *testing.T) {

	flag := VerticalParityError

	s := flag.String()

	if s != "Vertical parity error" {
		t.Error("Error")
	}
}

func TestTwoFlags(t *testing.T) {

	flag := VerticalParityError | LongitudalParityError

	s := flag.String()

	if s != "Longitudal parity error, Vertical parity error" {
		t.Errorf("Don't expected: %s", s)
	}
}

func TestAllFlags(t *testing.T) {

	flag := CommunicationsErrorSummaryFlags(0x7f)

	s := flag.String()

	if s != "Undefined, Buffer Overflow error, Reserved, Longitudal parity error, Framing error, Overrun error, Vertical parity error" {
		t.Errorf("Don't expected: %q", s)
	}
}
