/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/
package cmd

import (
	"github.com/jf-Lindberg/lalalint/linter"
	"github.com/jf-Lindberg/lalalint/validate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Outputs all of the linter errors found to the terminal.",
	Long:  `longer description`,
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}
		for i := range args {
			err := validate.FileName(args[i])
			if err != nil {
				return err
			}
		}

		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		viper.Set("global.showErrors", viper.GetBool("commands.check.showErrors"))
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		for i := range args {
			linter.Check(args[i])
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
