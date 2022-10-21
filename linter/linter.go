/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/
package linter

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jf-Lindberg/lalalint/helper"
	"github.com/spf13/viper"
	"time"
)

// doNothing does nothing.
// This is needed so that the output from lintedStream is at least read.
// This enables linter problems to be printed.
func doNothing(lintedStream <-chan Line) {
	for range lintedStream {
	}
}

// lint contains logic for linting lines concurrently.
// The actual linting is done by methods on the Line struct.
// Returns a channel which contents can be read and/or further manipulated.
func lint(inputStream <-chan Line) <-chan Line {
	// sets up chan for linted lines
	lintedStream := make(chan Line)

	go func() {
		// keeps track of how many blank lines in a row there are
		var blankLines int
		// ensures chan is closed when linting is done
		defer close(lintedStream)
		// loops through input chan (feed from file)
		for line := range inputStream {
			// calls method which calls all other methods (linter rules) of Line struct
			line = line.Lint(blankLines)
			if line.Content == "" {
				blankLines++
			} else {
				// reset counter if line not blank
				blankLines = 0
			}
			lintedStream <- line
		}
	}()
	// returns channel for use elsewhere
	return lintedStream
}

// printElapsed prints duration of time. Time started by Check/Write/Overwrite.
func printElapsed(cmd string, t time.Duration) {
	fmt.Printf("%s took %s\n", cmd, t)
}

// printLinterProblems will always output the amount of problems found during linting.
// If verbose mode is set to true, it will also output all the linter problems found in the file.
func printLinterProblems() {
	if len(lintErr) > 0 {
		if viper.GetBool("global.verbose") {
			helper.AnnounceStart("Problems detected:\n")
			color.Set(color.FgHiRed)
			for i := range lintErr {
				fmt.Printf("\t%s\n", lintErr[i])
			}
			color.Unset()
			fmt.Println()
			return
		}
	}

	// Gets amount of errors found
	amt := len(lintErr)
	s := fmt.Sprintf("%d problems found", amt)
	color.Set(color.FgHiRed)
	fmt.Println(s)
	color.Unset()
}
