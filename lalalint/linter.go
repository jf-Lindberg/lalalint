/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/
package lalalint

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/jf-Lindberg/lalalint/helper"
	"github.com/spf13/viper"
	"os"
	"time"
)

// Line Main struct for a line in a file.
type Line struct {
	Row     int
	Content string
}

//// Config Contains unmarshalled values from JSON config.
//type Config struct {
//	Global   Global   `mapstructure:"global"`
//	Commands Commands `mapstructure:"commands"`
//	Rules    Rules    `mapstructure:"rules"`
//}
//
//type Global struct {
//	ShowErrors bool `mapstructure:"showErrors"`
//}
//
//type Commands struct {
//	Print PrintCfg `mapstructure:"print"`
//	Write WriteCfg `mapstructure:"write"`
//	Check CheckCfg `mapstructure:"check"`
//}
//
//type PrintCfg struct {
//}
//
//type WriteCfg struct {
//}
//
//type CheckCfg struct {
//}
//
//type Rules struct {
//	SpaceAfterComments SpaceAfterComments `mapstructure:"spaceAfterComments"`
//	IndentEnvironments IndentEnvironments `mapstructure:"indentEnvironments"`
//}
//
//type SpaceAfterComments struct {
//	Enabled bool   `mapstructure:"enabled"`
//	Symbol  string `mapstructure:"symbol"`
//}
//
//type IndentEnvironments struct {
//	Enabled bool `mapstructure:"enabled"`
//	Indent  int  `mapstructure:"indent"`
//}
//
//type Comment struct {
//	Content Line
//	SpaceAfterComments
//}
//
//type rule interface {
//	lint() Line
//}
//
//var rulesCfg Rules
//
//func (c Comment) lint() Line {
//	if c.Enabled {
//		line, err := LintComment(c.Content, c.Symbol)
//		if err != nil {
//			lintErr = append(lintErr, err)
//		}
//		return line
//	}
//	return c.Content
//}

// would it be possible to write the above function in a more generalized way?
// for example, send in a rule as arg, if r enabled, do lint

var lintErr = make([]error, 0)
var indentLevel int

// getFile Gets a pointer to a file and returns it. If unsuccessful, logs error and exits with code 0
func getFile(path string, filename string) *os.File {
	file, err := os.Open(path + filename)
	helper.LogFatal(err)
	return file
}

// getLine Continuously feeds the channel "lines" with Line structs. If err, feeds to "readerr" chan.
func getLine(filename string, lines chan Line, readerr chan error) {
	// Gets file and defers closing to ensure cleanup
	path := "./data/"
	file := getFile(path, filename)
	defer file.Close()
	// Initializes scanner
	scanner := bufio.NewScanner(file)
	// Initializes row number
	row := 1
	// Feeds Line structs to lines channel
	for scanner.Scan() {
		lines <- Line{row, scanner.Text()}
		row++
	}
	// Feeds any errors to readerr channel
	readerr <- scanner.Err()
}

// lintLine Calls all enabled linter rules and returns a Line struct with the fixed content
func lintLine(line Line) Line {
	// check if environment and add to / subtract from indentLevel
	//	line = Comment{line, rulesCfg.SpaceAfterComments}.lint() << might bring back
	// rulesCfg.SpaceAfterComments.Enabled << might bring back
	if viper.GetBool("rules.spaceAfterComments.enabled") {
		var err error
		line, err = LintComment(line, viper.GetString("rules.spaceAfterComments.symbol"))
		if err != nil {
			lintErr = append(lintErr, err)
		}
	}
	if viper.GetBool("rules.indentEnvironments.enabled") {
		line = IndentLine(indentLevel, viper.GetInt("rules.indentEnvironments.indent"), line)
	}
	return line
}

func printLinterErrors() {
	if viper.GetBool("global.showErrors") {
		if len(lintErr) > 0 {
			helper.AnnounceStart("Errors found:")
			for i := range lintErr {
				helper.PrintLintErr(lintErr[i])
			}
		}
	}
}

//func unmarshalCfg() {
//	// Unwraps JSON config into rulesCfg struct
//	err := viper.UnmarshalKey("rules", &rulesCfg)
//	helper.LogFatal(err)
//}

//func countLines(filename string) int {
//	path := "./data/"
//	file := getFile(path, filename)
//	scanner := bufio.NewScanner(file)
//
//	count := 0
//	for scanner.Scan() {
//		count++
//	}
//	return count
//}

func Check(input string) {
	// Initialises timer for benchmarking
	start := time.Now()
	//unmarshalCfg() << might bring back
	// Creates channels
	lines := make(chan Line)
	readerr := make(chan error)

	helper.AnnounceStart("Checking " + input + " for errors")
	// Starts goroutine to feed lines into channel
	go getLine(input, lines, readerr)

loop:
	// Continuously reads from line channel until it closes (when there's nothing being fed to it)
	for {
		select {
		case line := <-lines:
			lintLine(line)
		case err := <-readerr:
			helper.LogFatal(err)
			break loop
		}
	}

	// Finishes up - prints errors, an announcement and the time elapsed since function was called
	printLinterErrors()
	helper.AnnounceDone("✔ Check finished")
	elapsed := time.Since(start)
	fmt.Printf("Check took %s\n", elapsed)
}

func Print(input string) {
	// Initialises timer for benchmarking
	start := time.Now()
	//unmarshalCfg() << might bring back
	// Creates channels
	lines := make(chan Line)
	readerr := make(chan error)

	helper.AnnounceStart("Printing " + input)
	// Starts goroutine to feed lines into channel
	go getLine(input, lines, readerr)

	color.Set(color.FgHiWhite)
loop:
	// Continuously reads from line channel until it closes (when there's nothing being fed to it)
	for {
		select {
		case line := <-lines:
			if viper.GetBool("commands.print.lint") == true {
				line = lintLine(line)
			}
			fmt.Printf("%d) %s\n", line.Row, line.Content)
		case err := <-readerr:
			helper.LogFatal(err)
			break loop
		}
	}
	// Finishes up - prints errors, an announcement and the time elapsed since function was called
	color.Unset()
	if viper.GetBool("commands.print.lint") == true {
		printLinterErrors()
	}
	helper.AnnounceDone("✔ Print done")
	elapsed := time.Since(start)
	fmt.Printf("Print took %s\n", elapsed)
}

func Write(input string, output string) {
	start := time.Now()
	//unmarshalCfg() << might bring back
	lines := make(chan Line)
	readerr := make(chan error)

	outputPath := "./data/output/"
	outputFile, err := os.Create(outputPath + output)
	helper.LogFatal(err)

	helper.AnnounceStart("Writing source '" + input + "' to file '" + output + "'")

	go getLine(input, lines, readerr)

loop:
	for {
		select {
		case line := <-lines:
			line = lintLine(line)
			_, err := outputFile.WriteString(line.Content + "\n")
			helper.LogFatal(err)
		case err := <-readerr:
			helper.LogFatal(err)
			break loop
		}
	}

	printLinterErrors()
	helper.AnnounceDone("✔ Done writing, your file is saved at '" + outputPath + output + "'")
	elapsed := time.Since(start)
	fmt.Printf("Write took %s\n", elapsed)
}
