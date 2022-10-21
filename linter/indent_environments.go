/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/
package linter

import (
	"fmt"
	"github.com/spf13/viper"
	"regexp"
	"strings"
)

type NoIndentErr struct {
	Row int
}

func (e NoIndentErr) Error() string {
	return fmt.Sprintf("environment not indented correctly on row %d", e.Row)
}

// Indent takes an input Line and adds indentation to it if needed.
// It takes three arguments: the current indentation level, the amount of tabs that each level should represent and the input Line.
// In order to avoid false positives, Indent compares the current indentation of the line with the indentation calculated from the current indentlevel * tabs.
// If they're the same, the input Line will be returned without any NoIndentErr - which means that the linter won't output any linter problem for that line.
func Indent(indentLevel int, tabs int, line Line) (Line, error) {
	indentation := indentLevel * tabs
	if indentation >= 0 {
		currentTabs := strings.Count(line.Content, "\t")
		if currentTabs == indentation {
			return line, nil
		}
		content := strings.Trim(line.Content, "\t")
		r := regexp.MustCompile("\n")
		indent := strings.Repeat("\t", indentation)
		replaced := r.ReplaceAllString(content, "\n"+indent)
		content = indent + replaced
		return Line{line.Row, content}, NoIndentErr{line.Row}
	}

	return line, nil
}

// IndentEnvironments calls Indent with the current indentLevel, but also keeps track of whether the current Line contains environments or not.
// If a line contains the same amount of beginnings and endings of environments, they cancel each other out.
// Otherwise, the difference of the two gets added / subtracted from current indentLevel.
func IndentEnvironments(indentLevel int, tabs int, line Line) (Line, error) {
	indented, err := Indent(indentLevel, tabs, line)
	if isEnvironment(line.Content) && !excluded(getEnvironment(line.Content)) {
		cIndex := GetIndex(CommentRegex(), line)
		eIndex := GetIndex(EnvironmentRegex(), line)
		if cIndex == nil || cIndex[0] > eIndex[0] {
			diff := begunEnvironments(line.Content) - endedEnvironments(line.Content)
			indentLevel += diff
			viper.Set("rules.indentenvironments.indentlevel", indentLevel)
			if diff < 0 {
				return Indent(indentLevel, tabs, line)
			}
		}
	}
	return indented, err
}

// excluded checks if the environment found is excluded in the config file.
func excluded(environment string) bool {
	excl := viper.Get("rules.indentenvironments.excluded")
	for _, v := range excl.([]interface{}) {
		if environment == v {
			return true
		}
	}
	return false
}

// begunEnvironments counts the amount of begun environments in a Line's content.
func begunEnvironments(content string) int {
	pattern := regexp.MustCompile(`(\\begin) ?({.*})`)
	matches := pattern.FindAll([]byte(content), -1)
	amount := len(matches)
	return amount
}

// endedEnvironments counts the amount of ended environments in a Line's content.
func endedEnvironments(content string) int {
	pattern := regexp.MustCompile(`(\\end) ?({.*})`)
	matches := pattern.FindAll([]byte(content), -1)
	amount := len(matches)
	return amount
}

// isBeginning returns whether the Line's content contains the beginning of an environment.
func isBeginning(content string) bool {
	pattern := regexp.MustCompile(`(\\begin) ?({.*})`)
	return pattern.Match([]byte(content))
}

// isEnd returns whether the Line's content contains the end of an environment.
func isEnd(content string) bool {
	pattern := regexp.MustCompile(`(\\end) ?({.*})`)
	return pattern.Match([]byte(content))
}

// isEnvironment returns whether the Line's content contains an environment.
func isEnvironment(content string) bool {
	return isBeginning(content) || isEnd(content)
}

// getEnvironment gets the string representation of the environment.
func getEnvironment(content string) string {
	pattern := regexp.MustCompile(`(\\begin|\\end) ?{(?P<environment>.*)}`)
	// FindStringSubMatch returns a slice of which index 0 is the entire match, index 1 is the first capturing group and
	// index 2 is the second capturing group (which holds the environment).
	environment := pattern.FindStringSubmatch(content)[2]

	return environment
}
