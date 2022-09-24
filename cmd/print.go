/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/
package cmd

import (
	"github.com/jf-Lindberg/lalalint/helper"
	"github.com/jf-Lindberg/lalalint/linter"
	"github.com/jf-Lindberg/lalalint/validate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Cfg struct {
	Lint   bool `mapstructure:"lint"`
	Errors bool `mapstructure:"errors"`
}

var (
	PrintCfg Cfg
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "print <filename>",
	Short: "Prints the file contents with line numbers to terminal",
	Long: `Print takes the argument <filename>, reads the file and prints it line by line with line numbers to the terminal.
By default, it DOES NOT print the linted version of the file. To turn this on temporarily, you have to use the -l/--lint flag:

	lalalint print example.tex --lint

If you would like linter errors to be printed to the terminal as well, you can use the -e/--errors flag:

	lalalint print example.tex --lint --errors

The lint and error flags must be used together, the command "print example.tex -e" will show an error. 

You can also change the default behaviour by editing the config file.`,
	Args: cobra.MinimumNArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		for i := range args {
			err := validate.FileName(args[i])
			if err != nil {
				return err
			}
		}
		if errors, _ := rootCmd.PersistentFlags().GetBool("errors"); errors != false {
			cmd.MarkFlagRequired("lint")
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0]
		linter.Print(filename)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(printCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// printCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	printCmd.Flags().BoolP("lint", "l", false, "prints linted version of file and linter errors")
	err := viper.BindPFlag("commands.print.lint", printCmd.Flags().Lookup("lint"))
	helper.LogFatal(err)
}
