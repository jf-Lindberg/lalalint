/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/
package cmd

import (
	"fmt"
	"github.com/jf-Lindberg/lalalint/configMgmt"
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Contains subcommands 'get' and 'set' for reading and writing configuration parameters",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

// getCmd gets a specific parameter from config
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets a specific parameter from config. If no arguments are passed, returns entire config",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			fmt.Println(configMgmt.GetConfig())
			return nil
		}
		fmt.Println(configMgmt.GetVal(args[0]))
		return nil
	},
}

// setCmd sets a specific parameter in config
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets a specific parameter in config",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		configMgmt.SetVal(args[0], args[1])
		fmt.Printf("%s now set to %s", args[0], args[1])
	},
}

func init() {
	configCmd.AddCommand(getCmd)
	configCmd.AddCommand(setCmd)
	rootCmd.AddCommand(configCmd)
}
