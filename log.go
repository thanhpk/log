package log

import (
	"strings"
	"os"
	"fmt"
	"runtime"
	"time"
	"github.com/fatih/color"
)

var (
	blue = color.New(color.FgBlue).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
	red = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
)

var starttime = time.Now()

// Logf log with format to stdout
func Logf(format string, v ...interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	var split = strings.Split(fn, string(os.PathSeparator))
	var n int
	if len(split) >= 2 {
		n = len(split) - 2
	} else {
		n = len(split)
	}
	fn = strings.Join(split[n:], string(os.PathSeparator))
	message := fmt.Sprintf(format, v...)
	fmt.Printf(getCurrentTimeString() + blue("%s:%d ") + green("%s") + "\n", fn, line, message)
	color.Unset()
}

// Log print anything to stdout
func Log(v ...interface{}) {
	//color.Set(color.FgGreen)
	_, fn, line, _ := runtime.Caller(1)
	var split = strings.Split(fn, string(os.PathSeparator))
	var n int
	if len(split) >= 2 {
		n = len(split) - 2
	} else {
		n = len(split)
	}
	fn = strings.Join(split[n:], string(os.PathSeparator))
	format := strings.Repeat("%v ", len(v))
	message := fmt.Sprintf(format, v...)

	fmt.Printf(getCurrentTimeString() + blue("%s:%d ") + green("%s") + "\n", fn, line, message)
	color.Unset()
}

func getCurrentTimeString() string {
	now := time.Now()
	m := fmt.Sprintf("%d", int(now.Month()))
	if m == "10" {
		m = "O"
	} else if m == "11" {
		m = "N"
	} else if m == "12" {
		m = "D"
	}

	var ds string
	d := now.Day()
	if d > 9 {
		ds = string(rune(d + 87))
	} else {
		ds = string(rune(d + 48))
	}
	return m + ds + fmt.Sprintf("%d %d:%d:%d ", int(time.Since(starttime).Seconds()),
		now.Hour(), now.Minute(), now.Second())
}
