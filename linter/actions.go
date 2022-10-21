/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/
package linter

import (
	"github.com/jf-Lindberg/lalalint/helper"
	"github.com/spf13/viper"
	"os"
	"time"
)

// Check goes through all the rules and gathers which linter problems are present in the file.
// It does not format the input file or write to a new file.
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

// Overwrite contains the command logic for overwriting a file.
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

// Write contains the command logic for writing to a new file.
func Write(input string, output string) {
	start := time.Now()

	outputPath := viper.GetString("global.outputdirectory")
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
