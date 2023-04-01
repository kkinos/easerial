package serial

import (
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"go.bug.st/serial"
)

func SendHexString(portName string, baudRate int, dataBits int, hexString string, readBytes int) error {
	mode := &serial.Mode{
		BaudRate: baudRate,
		DataBits: dataBits,
	}

	port, err := serial.Open(portName, mode)
	if err != nil {
		return fmt.Errorf("Failed to open serial port `%s`", portName)
	}
	defer port.Close()

	port.SetReadTimeout(5 * time.Second)
	port.ResetInputBuffer()
	port.ResetOutputBuffer()

	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return errors.New("Failed to decode hex")
	}
	_, err = port.Write(bytes)
	if err != nil {
		return errors.New("Failed to send data")
	}

	buff := make([]byte, readBytes)
	_, err = port.Read(buff)
	if err != nil {
		return errors.New("Failed to read data")
	}
	fmt.Println(buff)

	return nil
}
