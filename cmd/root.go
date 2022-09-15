/*
Copyright © 2022 Filip Lindberg jakob.filip.lindberg@gmail.com
*/
package cmd

import (
	"fmt"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/jf-Lindberg/lalalint/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile   string
	inputFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lalalint",
	Short: "Lalalint is a LaTeX linter for .tex files. Basic usage: lalalint <file>.",
	Long:  `Lalalint is a LaTeX linter for .tex files.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cmd.PersistentFlags().Lookup("errors").Changed {
			viper.Set("errors", true)
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		helper.LogFatal(err)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cc.Init(&cc.Config{
		RootCmd:  rootCmd,
		Headings: cc.HiCyan + cc.Bold + cc.Underline,
		Commands: cc.HiYellow + cc.Bold,
		Example:  cc.Italic,
		ExecName: cc.Bold,
		Flags:    cc.Bold,
	})
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolP("errors", "e", false, "prints errors (default is off except for the check command)")
	viper.BindPFlag("showErrors", rootCmd.PersistentFlags().Lookup("errors"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("json")
		viper.SetConfigName(".lalalint")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
