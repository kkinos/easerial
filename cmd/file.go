/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"os"
	"strings"

	"github.com/kinpoko/easerial/serial"
	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file [file]",
	Short: "Send data from a text file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fileBytes, err := os.ReadFile(args[0])
		if err != nil {
			return errors.New("Faild to open file")
		}

		hexString := strings.ReplaceAll(string(fileBytes), "\n", "")
		hexString = strings.ReplaceAll(hexString, " ", "")
		hexString = strings.ReplaceAll(hexString, ",", "")

		err = serial.SendHexString(portName, baudRate, dataBits, hexString, readBytes)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	fileCmd.Flags().StringVar(&portName, "port", "/dev/ttyUSB1", "port name")
	fileCmd.Flags().IntVar(&baudRate, "baud", 9600, "baud rate")
	fileCmd.Flags().IntVar(&dataBits, "data-bits", 8, "data bits")
	fileCmd.Flags().IntVar(&readBytes, "read-bytes", 4, "read bytes")
	rootCmd.AddCommand(fileCmd)

}