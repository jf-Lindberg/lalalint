/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/
package cmd

import (
	"github.com/jf-Lindberg/lalalint/lalalint"
	"github.com/jf-Lindberg/lalalint/validate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Outputs all of the linter errors found to the terminal.",
	Long:  `longer description`,
	Args:  cobra.MinimumNArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		for i := range args {
			err := validate.FileName(args[i])
			if err != nil {
				return err
			}
		}
		viper.Set("global.showErrors", viper.GetBool("commands.check.showErrors"))
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		for i := range args {
			lalalint.Check(args[i])
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
