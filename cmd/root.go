package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var portName string
var baudRate int
var dataBits int
var readBytes int

var rootCmd = &cobra.Command{
	Use:   "easerial",
	Short: "A simple command line application for serial communication.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You must specify a subcommand. See --help more infomation")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
}
