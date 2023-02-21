package cmd

import (
	"github.com/kinpoko/easerial/serial"
	"github.com/spf13/cobra"
)

// stringCmd represents the string command
var stringCmd = &cobra.Command{
	Use:   "string [hex string]",
	Short: "Send data from a hex string",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		hexString := args[0]

		err := serial.SendHexString(portName, baudRate, dataBits, hexString, readBytes)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	stringCmd.Flags().StringVar(&portName, "port", "/dev/ttyUSB1", "port name")
	stringCmd.Flags().IntVar(&baudRate, "baud", 9600, "baud rate")
	stringCmd.Flags().IntVar(&dataBits, "data-bits", 8, "data bits")
	stringCmd.Flags().IntVar(&readBytes, "read-bytes", 4, "read bytes")
	rootCmd.AddCommand(stringCmd)
}
