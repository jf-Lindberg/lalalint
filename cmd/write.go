/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/jf-Lindberg/lalalint/linter"
	"github.com/jf-Lindberg/lalalint/validate"
	"github.com/spf13/cobra"
)

// writeCmd represents the write command
var writeCmd = &cobra.Command{
	Use:   "write <inputfile> <outputfile>",
	Short: "Writes a linted version of the input file to the output file.",
	Long:  `longer description`,
	Args:  cobra.MinimumNArgs(2),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		for i := range args {
			err := validate.FileName(args[i])
			if err != nil {
				return err
			}
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		input, output := args[0], args[1]
		linter.Write(input, output)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(writeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// writeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// writeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
