// This file contains error types use by CommandExecutor func

package hart

import (
	"errors"
	"fmt"
)

// ErrNoResponse - there is no data received
var ErrNoResponse error = errors.New("No response")

// FrameDataParsingError - Frame.SetData method returns false
type FrameDataParsingError struct {
	frame Frame
}

func (e FrameDataParsingError) Error() string {
	return fmt.Sprintf("Invalid frame data: %v", e.frame)
}

// FrameParsingError - hart.Parse frame func returns false
type FrameParsingError struct {
	frame Frame
}

func (e *FrameParsingError) Error() string {
	return fmt.Sprintf("Invalid hart frame: %v", e.frame)
}
