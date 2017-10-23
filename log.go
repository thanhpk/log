package log

import (
	"strings"
	"os"
	"fmt"
	"runtime"
	"log"
	"github.com/fatih/color"
	"runtime/debug"
)

var (
	blue = color.New(color.FgBlue).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
	red = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
)

type Info map[string]interface{}

func getMinifiedStack() string {
	stack := ""
	for i := 3; i < 90; i++ {
		pc, fn, line, _ := runtime.Caller(i)
		f := runtime.FuncForPC(pc)

		if fn == "" {
			break
		}
		hl := false // highlight
		if strings.Contains(fn, "bitbucket.org/subiz") {
			hl = true
		}
		var split = strings.Split(fn, string(os.PathSeparator))
		var n int

		if len(split) >= 2 {
			n = len(split) - 2
		} else {
			n = len(split)
		}
		fn = strings.Join(split[n:], string(os.PathSeparator))
		if hl {
			stack += fmt.Sprintf(yellow("\t→ %s:%d ") + red("%s\n"), fn, line, f.Name())
		} else {
			stack += fmt.Sprintf(blue("\t→ %s:%d ") + red("%s\n"), fn, line, f.Name())
		}
	}
	return stack
}

func logMap(m map[string]interface{}) {
	for key, value := range m {
    printlog("%s: %s", key, fmt.Sprintf("%v", value))
	}
}

func WithStack(v ...interface{}) {
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
	log.Printf(blue("%s:%d ") + green("%s"), fn, line, message)
	color.Unset()
	stack := getMinifiedStack()
	fmt.Print(stack)
}

func LogError(info Info) {
//	info["stacktrace"] = fmt.Sprintf("%s", debug.Stack())
	logMap(info)
}

func LogPanic(info Info) {
	info["stacktrace"] = fmt.Sprintf("%s", debug.Stack())
	logMap(info)
}

// Log print anything to stdout
func printlog(f interface{}, v ...interface{}) {
	format, ok := f.(string)
	if !ok {
		v = append([]interface{}{f}, v...)
		format = strings.Repeat("%v ", len(v))
	}
	//color.Set(color.FgGreen)

	var message string
	if len(v) > 0 {
		message = fmt.Sprintf(format, v...)
	} else {
		message = format
	}

	fmt.Println(yellow("└ " + message))
	color.Unset()
}

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
	log.Printf(blue("%s:%d ") + green("%s"), fn, line, message)
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
	log.Printf(blue("%s:%d ") + green("%s"), fn, line, message)
	color.Unset()
}
