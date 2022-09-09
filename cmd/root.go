/*
Copyright © 2022 Filip Lindberg jakob.filip.lindberg@gmail.com
*/
package cmd

import (
	"fmt"
	"github.com/jf-Lindberg/lalalint/pkg/validate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile    string
	outputFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lalalint <filename> [flags]",
	Short: "Lalalint is a LaTeX linter for .tex files. Basic usage: lalalint <file>.",
	Long:  `Lalalint is a LaTeX linter for .tex files.`,
	Args:  cobra.ExactArgs(1),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		fileName := args[0]
		err := validate.FileName(fileName)
		if err != nil {
			return err
		}

		output := cmd.Flags().Lookup("write").Value
		fmt.Println(output)
		fmt.Println("GetString:", viper.GetString("write"))

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $HOME/.lalalint.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&outputFile, "write", "w", "", "Name of output file ending with .tex")
	rootCmd.Flags().BoolP("overwrite", "o", false, "Overwrites input file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".lalalint" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".lalalint")
	}

	viper.BindPFlag("write", rootCmd.Flags().Lookup("write"))

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
