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

type Line struct {
	Row     int
	Content string
}

type Rules struct {
	Comments           CommentRule `mapstructure:"spaceAfterComments"`
	IndentEnvironments IndentRule  `mapstructure:"indentEnvironments"`
}

type CommentRule struct {
	Enabled bool   `mapstructure:"enabled"`
	Symbol  string `mapstructure:"symbol"`
}

type IndentRule struct {
	Enabled bool `mapstructure:"enabled"`
	Indent  int  `mapstructure:"indent"`
}

type Comment struct {
	Content Line
	CommentRule
}

type rule interface {
	lint() Line
}

var rulesCfg Rules

func (c Comment) lint() Line {
	if c.Enabled {
		line, err := LintComment(c.Content, c.Symbol)
		if err != nil {
			lintErr = append(lintErr, err)
		}
		return line
	}
	return c.Content
}

var lintErr = make([]error, 0)
var indentLevel int

func getFile(path string, filename string) *os.File {
	file, err := os.Open(path + filename)
	helper.LogFatal(err)
	return file
}

func getLine(filename string, lines chan Line, readerr chan error) {
	path := "./data/"
	file := getFile(path, filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	row := 1
	for scanner.Scan() {
		lines <- Line{row, scanner.Text()}
		row++
	}
	readerr <- scanner.Err()
}

func lintLine(line Line) Line {
	line = Comment{line, rulesCfg.Comments}.lint()
	//if rulesCfg.Comments.Enabled {
	//	var err error
	//	line, err = LintComment(line, rulesCfg.Comments.Symbol)
	//	if err != nil {
	//		lintErr = append(lintErr, err)
	//	}
	//}
	if rulesCfg.IndentEnvironments.Enabled {
		line = IndentLine(indentLevel, rulesCfg.IndentEnvironments.Indent, line)
	}
	return line
}

func printLinterErrors() {
	if viper.GetBool("showErrors") {
		if len(lintErr) > 0 {
			helper.Announce("Errors found:")
			for i := range lintErr {
				helper.PrintLintErr(lintErr[i])
			}
		}
	}
}

func unmarshalCfg() {
	err := viper.UnmarshalKey("rules", &rulesCfg)
	helper.LogFatal(err)
}

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
	start := time.Now()
	unmarshalCfg()
	lines := make(chan Line)
	readerr := make(chan error)

	helper.Announce("Checking " + input + " for errors")

	go getLine(input, lines, readerr)

loop:
	for {
		select {
		case line := <-lines:
			lintLine(line)
		case err := <-readerr:
			helper.LogFatal(err)
			break loop
		}
	}

	printLinterErrors()
	helper.Announce("Check finished")
	elapsed := time.Since(start)
	fmt.Printf("Check took %s\n", elapsed)
}

func Print(input string) {
	start := time.Now()
	unmarshalCfg()
	lines := make(chan Line)
	readerr := make(chan error)

	helper.Announce("Printing " + input)

	go getLine(input, lines, readerr)

	color.Set(color.FgGreen)
loop:
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

	color.Unset()
	if viper.GetBool("commands.print.lint") == true {
		printLinterErrors()
	}
	helper.Announce("Print done")
	elapsed := time.Since(start)
	fmt.Printf("Print took %s\n", elapsed)
}

func Write(input string, output string) {
	start := time.Now()
	unmarshalCfg()
	lines := make(chan Line)
	readerr := make(chan error)

	outputPath := "./data/output/"
	outputFile, err := os.Create(outputPath + output)
	helper.LogFatal(err)

	helper.Announce("Writing source '" + input + " to file '" + output + "'")

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
	helper.Announce("Done writing, your file is saved at '" + outputPath + output + "'")
	elapsed := time.Since(start)
	fmt.Printf("Write took %s\n", elapsed)
}
