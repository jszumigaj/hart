package hart

import (
	"errors"
	"sync"
	"time"

	"github.com/jszumigaj/serial"
)

// SerialPort struct
type SerialPort struct {
	sync.Mutex
	serial *serial.Port
}

// Open port to HART communication
func Open(portName string) (*SerialPort, error) {
	conf := &serial.Config{
		Name:           portName,
		Baud:           1200,
		Parity:         serial.ParityOdd,
		StopBits:       serial.Stop1,
		ReadTimeout:    100 * time.Millisecond,
		RTSFlowControl: true,
	}

	s, err := serial.OpenPort(conf)
	if err != nil {
		return nil, err
	}

	port := SerialPort{serial: s}
	return &port, nil
}

// Close serial port
func (p *SerialPort) Close() {
	p.serial.Close()
}

// SendFrame send txBuf and wait for response. Returns number of bytes received.
func (p *SerialPort) SendFrame(txBuf []byte, rxBuf []byte) (int, error) {
	p.Lock()
	defer p.Unlock()
	// wait for end prev transmission (CD low)
	if cd, err := waitForNoCD(p.serial, 500*time.Millisecond); !cd {
		if err != nil {
			return 0, err
		}
		return 0, errors.New("Carier detected")
	}

	// write tx data
	p.serial.Write(txBuf)

	// wait for CD (device starts reply)
	if cd, err := waitForCD(p.serial, 500*time.Millisecond); !cd {
		if err != nil {
			return 0, err
		}
		//return 0, nil
		//leave as normal from here, to purge rx buffer in case if there are some rabbish data
	}

	// read rx data
	r, err := readSerialData(p.serial, rxBuf)
	return r, err
}

func readSerialData(s *serial.Port, rxBuf []byte) (int, error) {
	var readed int = 0
	for {
		r, e := s.Read(rxBuf[readed:])
		if e != nil {
			return readed, e
		}

		if r == 0 {
			break
		}

		readed += r
	}

	return readed, nil
}

func waitForNoCD(s *serial.Port, timeOut time.Duration) (bool, error) {
	start := time.Now()
	for time.Since(start) < timeOut {
		time.Sleep(10 * time.Millisecond)
		status, e := s.GetComModemStatus()
		if e != nil {
			return false, e
		}

		//the term RLSD (Receive Line Signal Detect) is commonly referred to as the CD (Carrier Detect) line.
		if (status & serial.MS_RLSD_ON) != 0 {
			return true, nil
		}
	}

	return false, nil
}

func waitForCD(s *serial.Port, timeOut time.Duration) (bool, error) {
	start := time.Now()
	for time.Since(start) < timeOut {
		time.Sleep(10 * time.Millisecond)
		status, e := s.GetComModemStatus()
		if e != nil {
			return false, e
		}

		//the term RLSD (Receive Line Signal Detect) is commonly referred to as the CD (Carrier Detect) line.
		//CTS połączony jest z RTS
		if ((status & serial.MS_RLSD_ON) == 0) && ((status & serial.MS_CTS_ON) == 0) {
			return true, nil
		}
	}

	return false, nil
}
