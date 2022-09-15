package helper

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

func LogFatal(e error) {
	if e != nil {
		color.Set(color.Bold, color.FgHiRed)
		log.Fatal(e)
	}
}

func Announce(announcement string) {
	color.Set(color.FgHiWhite, color.Bold)
	fmt.Println(announcement)
	color.Unset()
}

func PrintLintErr(err error) {
	color.Set(color.FgHiRed)
	fmt.Println(err)
	color.Unset()
}
