package cmd

import (
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"go.bug.st/serial"
)

func sendData(portName string, baudRate int, dataBits int, hexString string, readBytes int) error {
	mode := &serial.Mode{
		BaudRate: baudRate,
		DataBits: dataBits,
	}

	port, err := serial.Open(portName, mode)
	if err != nil {
		fmt.Println(err)
		return errors.New("Failed to open serial port")

	}
	defer port.Close()

	port.SetReadTimeout(1 * time.Second)
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
