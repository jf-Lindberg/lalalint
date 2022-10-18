package linter

import (
	"github.com/spf13/viper"
	"regexp"
)

func AddBlankLines(blankLinesCfg int, line Line, arr []Line) []Line {
	row := line.Row
	for i := 0; i < blankLinesCfg; i++ {
		arr = append(arr, Line{row, ""})
		row++
	}
	arr = append(arr, Line{row, line.Content})
	/*content := strings.Repeat("\n", blankLinesCfg) + line.Content*/
	return arr
}

// NEEDS TO HAVE A CHECK IN PLACE OF WHETHER NEWLINES ALREADY EXIST

func BlankLinesBeforeSection(blankLinesCfg int, line Line) []Line {
	arr := make([]Line, 0)
	if viper.GetBool("rules.blanklinesbeforesection.enabled") && isSection(line.Content) {
		arr := AddBlankLines(blankLinesCfg, line, arr)
		return arr
	}

	arr = append(arr, line)
	return arr
}

func isSection(content string) bool {
	r := "^\\\\(?P<section>(sub)?section|chapter)"
	pattern := regexp.MustCompile(r)
	return pattern.Match([]byte(content))
}
