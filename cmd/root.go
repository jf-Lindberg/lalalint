/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/
package cmd

import (
	"fmt"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/jf-Lindberg/lalalint/linter"
	"github.com/jf-Lindberg/lalalint/validate"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	cfgFile string
)

// rootCmd represents the base command, "lalalint".
var rootCmd = &cobra.Command{
	Use:   "lalalint [inputfile | directory] [outputfile] [flags]",
	Short: "Lalalint is a LaTeX linter for .tex, .bib and .tikz files. Basic usage",
	Long: `Lalalint is a static code analyzer (a.k.a. linter) for LaTeX documents - specifically for .tex, .bib and .tikz-files. 
Lalalint can output any errors found to the terminal, write a linted version of the input file(s) to a new file or overwrite the input file(s). 
It works for both individual files and for directories. The linter rules can be enabled or disabled at will via JSON.

Linting an individual file:
In order to lint a file, at least one argument must be passed - the input file. If lalalint is called without anything else, it will output how many problems it found in the file.

Linting a directory:
Lalalint can be run in directory mode without any arguments. By default, lalalint will run on the directory specified under "inputdirectory" in the configuration file.
If a path is specified as argument, it will run on that directory instead. Please note that subdirectories will NOT be linted.
`,
	Args: func(cmd *cobra.Command, args []string) error {
		// Directory takes 0 or 1 arguments
		if cmd.Flags().Lookup("directory").Changed {
			if err := cobra.RangeArgs(0, 1)(cmd, args); err != nil {
				return err
			}
			return nil
		}
		// File takes 1 or 2 arguments
		if err := cobra.RangeArgs(1, 2)(cmd, args); err != nil {
			return err
		}
		// Validate filenames
		for i := range args {
			err := validate.FileName(args[i])
			if err != nil {
				return err
			}
		}
		return nil
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		vFlag := cmd.Flags().Lookup("verbose").Changed
		// Set verbose parameter in config to true if input by user
		if vFlag {
			viper.Set("global.verbose", true)
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		dFlag := cmd.Flags().Lookup("directory").Changed
		owFlag := cmd.Flags().Lookup("overwrite").Changed

		if dFlag {
			start := time.Now()
			if len(args) == 1 {
				viper.Set("global.inputdirectory", args[0])
			}
			// Go through directory
			err := WalkDir(owFlag)
			elapsed := time.Since(start)
			fmt.Printf("%s took %s to go through\n", "Directory", elapsed)
			return err
		}
		// If overwrite flag, overwrite input file
		if owFlag {
			input := args[0]
			linter.Overwrite(input)
			return nil
		}
		// If output file specified, write to that
		if len(args) == 2 {
			input, output := args[0], args[1]
			linter.Write(input, output)
			return nil
		}
		// Default to checking the file
		input := args[0]
		linter.Check(input)
		fmt.Println("To fix the problems, please specify an output file as a second argument or use the overwrite flag.")
		return nil
	},
}

// WalkDir traverses a directory set in the configuration file.
// Depending on what flags the user has passed, it will call different actions to do on the .tex, .bib and .tikz-files in the directory.
// If no flags are passed, each file is only checked for any linter problems.
func WalkDir(owFlag bool) error {
	// Get input directory (default in config, can be overwritten through argument by user)
	root := viper.GetString("global.inputdirectory")
	if !strings.HasSuffix(root, "/") {
		root = root + "/"
	}
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// better to check if path minus root contains / and return nil if true
		// if directory or inside directory, skip it
		if d.IsDir() || strings.Count(path, "/") > 1 {
			return nil
		}
		// validate filename
		err = validate.FileName(d.Name())
		// if not validated, skip it (any other files than .tex, .bib, .tikz)
		if err != nil {
			return nil
		}
		input := d.Name()
		if owFlag {
			linter.Overwrite(input)
			return nil
		}

		linter.Check(input)
		return nil
	})
	return err
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
	rootCmd.PersistentFlags().BoolP(
		"directory",
		"d",
		false,
		"Runs the linter on all .tex, .bib and .tikz files in the input directory (changed via config). "+
			"\nAn argument can be passed to override the default directory."+
			"\nExample: lalalint -d <directory>")
	rootCmd.PersistentFlags().BoolP("overwrite",
		"o",
		false,
		"Overwrites the input file(s). "+
			"If run together with the --directory flag, all files in the directory will be overwritten.")
	rootCmd.PersistentFlags().BoolP("verbose",
		"v",
		false,
		"Prints all linter problems found to the terminal.")
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
