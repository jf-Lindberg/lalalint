/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/
package linter

import "github.com/spf13/viper"

var lintErr = make([]error, 0)

// Line Main struct for a line in a file.
// Contains Row number and a Content string.
type Line struct {
	Row     int
	Content string
}

func (line Line) Section(blankLines int) Line {
	if viper.GetBool("rules.blanklinesbeforesection.enabled") {
		nrOfLines := viper.GetInt("rules.blanklinesbeforesection.lines")
		line, err := BlankLinesBeforeSection(nrOfLines, blankLines, line)
		if err != nil {
			lintErr = append(lintErr, err)
		}
		return line
	}
	return line
}

func (line Line) Indent() Line {
	if viper.GetBool("rules.indentenvironments.enabled") {
		tabs := viper.GetInt("rules.indentenvironments.tabs")
		indentLevel := viper.GetInt("rules.indentenvironments.indentlevel")
		line, err := IndentEnvironments(indentLevel, tabs, line)
		if err != nil {
			lintErr = append(lintErr, err)
		}
		return line
	}

	return line
}

func (line Line) Newline() Line {
	if viper.GetBool("rules.newlineaftersentence.enabled") {
		line, err := NewlineAfterSentence(line)
		if err != nil {
			lintErr = append(lintErr, err)
		}
		return line
	}
	return line
}

func (line Line) Comment() Line {
	if viper.GetBool("rules.spaceaftercomments.enabled") {
		line, err := SpaceAfterComment(line)
		if err != nil {
			lintErr = append(lintErr, err)
		}
		return line
	}
	return line
}

func (line Line) Lint(blankLines int) Line {
	linted := line.Newline()
	linted = linted.Indent()
	linted = linted.Comment()
	linted = linted.Section(blankLines)
	return linted
}
