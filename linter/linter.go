/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/
package linter

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/jf-Lindberg/lalalint/helper"
	"github.com/spf13/viper"
	"os"
	"strings"
	"time"
)

// Line Main struct for a line in a file.
type Line struct {
	Row     int
	Content string
}

var lintErr = make([]error, 0)

func doNothing(lintedStream <-chan Line) {
	for _ = range lintedStream {
	}
}

// getFile Gets a pointer to a file and returns it. If unsuccessful, logs error and exits with code 0
func getFile(path string, filename string) *os.File {
	file, err := os.OpenFile(path+filename, os.O_CREATE|os.O_RDWR, os.ModePerm)
	helper.LogFatal(err)
	return file
}

// getLine
func getLines(filename string) <-chan Line {
	path := "./data/"
	file := getFile(path, filename)
	inputStream := make(chan Line)

	go func() {
		defer file.Close()
		defer close(inputStream)
		scanner := bufio.NewScanner(file)
		row := 1
		for scanner.Scan() {
			content := strings.TrimSpace(scanner.Text())
			inputStream <- Line{row, content}
			row++
		}
		if err := scanner.Err(); err != nil {
			helper.LogFatal(err)
		}
	}()

	return inputStream
}

// lint is the main function for calling all the linter rules.
// needs a LOT of refactoring, totally spaghetti at the moment
func lint(inputStream <-chan Line) <-chan Line {
	var err error
	var prev Line
	// sets up chan for linted lines
	lintedStream := make(chan Line)

	go func() {
		// ensures chan is closed when linting is done
		defer close(lintedStream)
		// loops through input chan (feed from file)
		for line := range inputStream {
			for _, newline := range NewlineAfterSentence(line) {
				// "This is a line\nIt continues here"
				line = newline.line
				if err = newline.err; err != nil {
					lintErr = append(lintErr, err)
				}
				if viper.GetBool("rules.spaceaftercomments.enabled") {
					line, err = SpaceAfterComment(line)
				}
				if err != nil {
					lintErr = append(lintErr, err)
				}
				if viper.GetBool("rules.indentenvironments.enabled") {
					tabs := viper.GetInt("rules.indentenvironments.tabs")
					indentLevel := viper.GetInt("rules.indentenvironments.indentlevel")
					line = IndentEnvironments(indentLevel, tabs, line)
				}
				// kanske hålla koll på hur många blanka rader innan så det funkar även med 2-3 som setting
				if viper.GetBool("rules.blanklinesbeforesection.enabled") && prev.Content != "" {
					for _, linestruct := range BlankLinesBeforeSection(viper.GetInt("rules.blanklinesbeforesection.lines"), line) {
						lintedStream <- linestruct
					}
					prev = line
					continue
				}
				prev = line
				lintedStream <- line
			}
		}
	}()

	return lintedStream
}

func overwriteFile(lintedStream <-chan Line, input string) {
	path := "./data/"
	backupName := "temp" + input
	backup, err := os.Create(path + backupName)
	if err != nil {
		helper.LogFatal(err)
	}
	for line := range lintedStream {
		_, err := backup.WriteString(line.Content + "\n")
		helper.LogFatal(err)
	}
	err = os.Rename(path+backupName, path+input)
	if err != nil {
		helper.LogFatal(err)
	}
}

func printElapsed(cmd string, t time.Duration) {
	fmt.Printf("%s took %s\n", cmd, t)
}

func printLines(lintedStream <-chan Line) {
	color.Set(color.FgHiWhite)
	for line := range lintedStream {
		fmt.Printf("%s\n", line.Content)
	}
	color.Unset()
}

func printLinterProblems() {
	if viper.GetBool("global.showErrors") {
		if len(lintErr) > 0 {
			helper.AnnounceStart("Problems found:\n")
			color.Set(color.FgHiRed)
			for i := range lintErr {
				fmt.Printf("\t%s\n", lintErr[i])
			}
			color.Unset()
			fmt.Println()
		}
	}
}

func writeLines(lintedStream <-chan Line, outputFile *os.File) {
	for line := range lintedStream {
		_, err := outputFile.WriteString(line.Content + "\n")
		helper.LogFatal(err)
	}
}

func Check(input string) {
	// Initialises timer for benchmarking
	start := time.Now()

	helper.AnnounceStart("Checking " + input + " for errors")
	lineStream := getLines(input)
	lintedStream := lint(lineStream)

	doNothing(lintedStream)

	// Finishes up - prints errors, an announcement and the time elapsed since function was called
	printLinterProblems()
	helper.AnnounceDone("✔ Check finished")
	elapsed := time.Since(start)
	printElapsed("Check", elapsed)
}

func Print(input string) {
	// Initialises timer for benchmarking
	start := time.Now()
	helper.AnnounceStart("Printing " + input)
	// Starts goroutine to feed lines into channel
	lineStream := getLines(input)
	lintedStream := lint(lineStream)

	printLines(lintedStream)

	// Finishes up - prints errors, an announcement and the time elapsed since function was called
	color.Unset()
	if viper.GetBool("commands.print.lint") {
		printLinterProblems()
	}
	helper.AnnounceDone("✔ Print done")
	elapsed := time.Since(start)
	printElapsed("Print", elapsed)
}

func Overwrite(input string) {
	start := time.Now()

	helper.AnnounceStart("Overwriting " + input)

	lineStream := getLines(input)
	lintedStream := lint(lineStream)

	overwriteFile(lintedStream, input)

	printLinterProblems()
	elapsed := time.Since(start)
	printElapsed("Overwrite", elapsed)
}

func Write(input string, output string) {
	start := time.Now()

	outputPath := "./data/output/"
	outputFile, err := os.Create(outputPath + output)
	helper.LogFatal(err)

	helper.AnnounceStart("Writing source '" + input + "' to file '" + output + "'")

	lineStream := getLines(input)
	lintedStream := lint(lineStream)

	writeLines(lintedStream, outputFile)

	printLinterProblems()
	helper.AnnounceDone("✔ Done writing, your file is saved at '" + outputPath + output + "'")
	elapsed := time.Since(start)
	printElapsed("Write", elapsed)
}
