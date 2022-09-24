/*
Copyright © 2022 Filip Lindberg fili21@student.bth.se
*/
package linter

import "strings"

func IndentLine(indentLevel int, indentCfg int, line Line) Line {
	content := strings.Repeat("\t", indentLevel*indentCfg) + line.Content
	return Line{line.Row, content}
}
