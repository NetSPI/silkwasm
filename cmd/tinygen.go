/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/netspi/silkwasm/smuggle"
	"github.com/spf13/cobra"
)

// tinygenCmd represents the tinygen command
var tinygenCmd = &cobra.Command{
	Use:   "tinygen",
	Short: "Generates WASM Decryptors for HTML Smuggling using TinyGo",
	Long: `Generate a WASM Decryptor to use in a HTML Smuggling page but using tinygo, to install tinygo visit https://tinygo.org/getting-started/,
	
	./silkwasm gen -i evildoc.zip -f SuperSecret`,
	Run: func(cmd *cobra.Command, args []string) {
		smuggle.NewSmuggler(input, funcname, debug, true)
	},
}

func init() {
	rootCmd.AddCommand(tinygenCmd)

	//Mandatory commands for all droppers.
	tinygenCmd.PersistentFlags().StringVarP(&input, "in", "i", "", "input file, shellcode.bin or exe.")
	tinygenCmd.MarkFlagRequired("in")

	//Optional commands for  wasm smuggles.
	tinygenCmd.PersistentFlags().StringVarP(&funcname, "funcname", "f", "SilkWasm", "The function name to call in the wasm file.")
	tinygenCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Show debug information and keep source files after compilation.")
}
