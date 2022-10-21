/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/
package linter

import (
	"bufio"
	"github.com/jf-Lindberg/lalalint/helper"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// getFile Gets a pointer to a file and returns it. If unsuccessful, logs error and exits with code 0
func getFile(path string, filename string) *os.File {
	file, err := os.OpenFile(path+filename, os.O_CREATE|os.O_RDWR, os.ModePerm)
	helper.LogFatal(err)
	return file
}

// getLines contains logic for reading a file concurrently.
// It returns a channel which gets fed with the contents from a specified file using bufios scanner.
func getLines(filename string) <-chan Line {
	path := viper.GetString("global.inputdirectory")
	file := getFile(path, filename)
	inputStream := make(chan Line)

	go func() {
		defer file.Close()
		defer close(inputStream)
		scanner := bufio.NewScanner(file)
		row := 1
		for scanner.Scan() {
			content := scanner.Text()
			if viper.GetBool("global.trimwhitespace") {
				content = strings.TrimSpace(content)
			}
			inputStream <- Line{row, content}
			row++
		}
		if err := scanner.Err(); err != nil {
			helper.LogFatal(err)
		}
	}()

	return inputStream
}

// overwriteFile creates a backup of the input file and writes the contents of a channel to it.
// It exits if there are any errors. If not, it renames the temporary file to the original name, essentially overwriting the input file.
func overwriteFile(lintedStream <-chan Line, input string) {
	path := viper.GetString("global.inputdirectory")
	backupName := "temp" + input
	backup, err := os.Create(path + backupName)
	if err != nil {
		helper.LogFatal(err)
	}
	writeLines(lintedStream, backup)
	err = os.Rename(path+backupName, path+input)
	if err != nil {
		helper.LogFatal(err)
	}
}

func writeLines(lintedStream <-chan Line, outputFile *os.File) {
	for line := range lintedStream {
		_, err := outputFile.WriteString(line.Content + "\n")
		helper.LogFatal(err)
	}
}
