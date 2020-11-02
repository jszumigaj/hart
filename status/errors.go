// This package contains error types used by Master.Execute func

package status

import (
	"errors"
	"fmt"
)

// ErrNoResponse - there is no data received
var ErrNoResponse error = errors.New("No response")

// FrameDataParsingError - Frame.SetData method returns false
type FrameDataParsingError struct {
	Frame []byte
}

func (e FrameDataParsingError) Error() string {
	return fmt.Sprintf("Invalid frame data: %v", e.Frame)
}

// FrameParsingError - hart.Parse frame func returns false
type FrameParsingError struct {
	Frame []byte
}

func (e *FrameParsingError) Error() string {
	return fmt.Sprintf("Invalid hart frame: %v", e.Frame)
}
