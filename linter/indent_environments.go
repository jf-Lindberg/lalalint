/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/
package linter

import (
	"github.com/spf13/viper"
	"regexp"
	"strings"
)

// create slice with current BEGUN environments

func Indent(indentLevel int, tabs int, line Line) Line {
	r := regexp.MustCompile("\n")
	indent := strings.Repeat("\t", indentLevel*tabs)
	replaced := r.ReplaceAllString(line.Content, "\n"+indent)
	content := indent + replaced
	return Line{line.Row, content}
}

func IndentEnvironments(indentLevel int, tabs int, line Line) Line {
	if isEnvironment(line.Content) && !excluded(getEnvironment(line.Content)) {
		cIndex := GetIndex(CommentRegex(), line)
		eIndex := GetIndex(EnvironmentRegex(), line)
		if cIndex == nil || cIndex[0] > eIndex[0] {
			if isBeginning(line.Content) {
				line = Indent(indentLevel, tabs, line)
				indentLevel++
				viper.Set("rules.indentenvironments.indentlevel", indentLevel)
				return line
			}
			if isEnd(line.Content) {
				indentLevel--
				viper.Set("rules.indentenvironments.indentlevel", indentLevel)
				return Indent(indentLevel, tabs, line)
			}
		}
	}

	return Indent(indentLevel, tabs, line)
}

func excluded(environment string) bool {
	switch environment {
	case
		"document":
		return true
	}
	return false
}

func isBeginning(content string) bool {
	pattern := regexp.MustCompile("(\\\\begin)({.*})")
	return pattern.Match([]byte(content))
}

func isEnd(content string) bool {
	pattern := regexp.MustCompile("(\\\\end)({.*})")
	return pattern.Match([]byte(content))
}

func isEnvironment(content string) bool {
	return isBeginning(content) || isEnd(content)
}

// getEnvironment might be used to check whether the environment is correct or not
func getEnvironment(content string) string {
	pattern := regexp.MustCompile("(\\\\begin|\\\\end){(?P<environment>.*)}")
	// FindStringSubMatch returns a slice of which index 0 is the entire match, index 1 is the first capturing group and
	// index 2 is the second capturing group (which holds the environment).
	environment := pattern.FindStringSubmatch(content)[2]

	return environment
}
