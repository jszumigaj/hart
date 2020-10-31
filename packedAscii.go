package hart

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"strings"
)

type PackedASCII []byte

// NewPackedASCII creates new PackedASCII
func NewPackedASCII(ascii string, packedLength int) PackedASCII {
	r := strings.NewReader(ascii)
	var packed = PackedASCII(make([]byte, packedLength))
	if _, err := io.Copy(packed, r); err != nil {
		panic(err)
	}
	return packed
}

func (packed PackedASCII) String() string {
	w := new(bytes.Buffer)
	if n, _ := io.Copy(w, packed); n > 0 {
		return w.String()
	}
	return ""
}

// Write writes bytes from source to the underlying data stream.
// It behaves different than normal Writer: if source ends it is padded with spaces, as is the nature of packetASCII
// It returns the number of bytes written from source (0 <= n <= len(b)). In case of padding it returns max len(src)
func (packed PackedASCII) Write(src []byte) (n int, err error) {
	n = 0  // licznik bajtów odczytanych ze źródła, inkrementowany co 4
	w := 0 // licznik bajtów wpisanych do dest, inkrementowany co 3
	a := bytes.ToUpper(src)
	srcLen := len(src)
	buf := make([]byte, 4)
	var bits uint32
	for {
		bits = bits + uint32(a[n]&0x3f)
		if n++; (n % 4) == 0 {
			//4 bytes from input are packed into 3 bytes of result integer
			binary.BigEndian.PutUint32(buf, bits)
			w += copy(packed[w:], buf[1:]) //copy 3 bytes
			if w >= len(packed) {               //checking for end
				if n < srcLen {
					return n, errors.New("Buffer too small to write all data")
				}
				return min(n, srcLen), nil
			}
		}
		// if src buf ends padd it right by space
		if n >= srcLen {
			a = append(a, ' ')
		}
		// shift result
		bits = bits << 6
	}
}

// Read reads up to len(packed) bytes into dst. It returns the number of bytes read (0 <= n <= len(dst))
func (packed PackedASCII) Read(dst []byte) (n int, err error) {
	r := 0 // bytes read from packed, incremented by 3
	n = 0  // wytes writen to dst, incremented by 4
	buf := make([]byte, 4)
	var bits uint32
	for {
		if r >= len(packed) {
			return n, io.EOF
		}
		if c := copy(buf[1:], packed[r:]); c == 3 { // copy 3 bytes of packed into buf
			r += 3
			bits = binary.BigEndian.Uint32(buf)
			for _, shift := range shifts {
				if n == len(dst) {
					return n, nil //buffer ends
				}
				dst[n] = byte((bits >> shift) & 0x3f)
				// set 6 bit if bit 5 == 0
				if (dst[n] & 0x20) == 0 {
					dst[n] |= 0x40
				}
				n++
			}
		}
	}
}

var shifts = []int{18, 12, 6, 0}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
