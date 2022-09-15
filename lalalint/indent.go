package lalalint

import "strings"

func IndentLine(indentLevel int, indentCfg int, line Line) Line {
	content := strings.Repeat("\t", indentLevel*indentCfg) + line.Content
	return Line{line.Row, content}
}
