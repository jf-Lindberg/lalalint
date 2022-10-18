/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/
package cmd

import (
	"github.com/jf-Lindberg/lalalint/linter"
	"github.com/jf-Lindberg/lalalint/validate"
	"github.com/spf13/cobra"
)

// overwriteCmd represents the overwrite command
var overwriteCmd = &cobra.Command{
	Use:   "overwrite",
	Short: "Not yet implemented",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(1)(cmd, args); err != nil {
			return err
		}
		err := validate.FileName(args[0])
		if err != nil {
			return err
		}
		return nil
	},
	/*	PreRunE: func(cmd *cobra.Command, args []string) error {
		for i := range args {
			err := validate.FileName(args[i])
			if err != nil {
				return err
			}
		}
		return nil
	},*/
	RunE: func(cmd *cobra.Command, args []string) error {
		input := args[0]
		linter.Overwrite(input)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(overwriteCmd)
}
