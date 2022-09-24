package helper

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"reflect"
)

func AnnounceDone(announcement string) {
	color.Set(color.FgHiGreen, color.Bold)
	fmt.Println(announcement)
	color.Unset()
}

func AnnounceStart(announcement string) {
	color.Set(color.FgHiCyan, color.Bold)
	fmt.Println(announcement)
	color.Unset()
}

func CompareType(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

func LogFatal(e error) {
	if e != nil {
		color.Set(color.Bold, color.FgHiRed)
		log.Fatal(e)
	}
}

func PrintLintErr(err error) {
	color.Set(color.FgHiRed)
	fmt.Println(err)
	color.Unset()
}
