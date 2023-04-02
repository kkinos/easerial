package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kinpoko/easerial/serial"
	"github.com/spf13/cobra"
)

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file [file_path]",
	Short: "Send hex data from a text file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fileBytes, err := os.ReadFile(args[0])
		if err != nil {
			return fmt.Errorf("Faild to open file `%s`", args[0])
		}

		r := strings.NewReplacer("\n", "", " ", "", ",", "")
		hexString := r.Replace(string(fileBytes))

		err = serial.SendHexString(portName, baudRate, dataBits, hexString, readBytes, readTimeoutSec)
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
	fileCmd.Flags().IntVar(&readTimeoutSec, "read-timeout-sec", 1, "read timeout sec")
	rootCmd.AddCommand(fileCmd)

}
