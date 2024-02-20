/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/netspi/silkwasm/smuggle"
	"github.com/spf13/cobra"
)

var input string
var funcname string

var debug bool

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a new WASM Smuggle",
	Long: `Generate a WASM Decryptor to use in a HTML Smuggling page:
	
	./silkwasm gen -i evildoc.zip -f SuperSecret`,
	Run: func(cmd *cobra.Command, args []string) {
		smuggle.NewSmuggler(input, funcname, debug, false)
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	//Mandatory commands for all droppers.
	genCmd.PersistentFlags().StringVarP(&input, "in", "i", "", "input file, your payload file e.g. initialaccess.zip.")
	genCmd.MarkFlagRequired("in")

	//Optional commands for all wasm smuggles.
	genCmd.PersistentFlags().StringVarP(&funcname, "funcname", "f", "SilkWasm", "The function name to call in the wasm file.")

	genCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Show debug information and keep source files after compilation.")
}
