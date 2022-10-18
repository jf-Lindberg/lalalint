package linter

import (
	"regexp"
)

func CommentRegex() *regexp.Regexp {
	r := "^%|[^\\\\]%"
	return regexp.MustCompile(r)
}

func EnvironmentRegex() *regexp.Regexp {
	r := "(\\\\begin|\\\\end){(.*)}"
	return regexp.MustCompile(r)
}

func GetIndex(r *regexp.Regexp, line Line) []int {
	return r.FindIndex([]byte(line.Content))
}

// GetFault gets the context of the error and returns it
func GetFault(index int, line Line) string {
	var fault string
	end := len(line.Content)
	if end-index > 5 {
		fault = line.Content[index : index+5]
	} else {
		fault = line.Content[index : end-1]
	}
	return fault
}
