package logger

import (
	"fmt"
	"os"
)

// Verbose output
var VerboseFlag bool

// DebugFlag sets debug output
var DebugFlag bool

// QuietFlag supresses all output
var QuietFlag bool

// Debug message (if DebugFlag is enabled)
func Debug(a ...interface{}) {
	if DebugFlag {
		fmt.Println(a...)
	}
}

// Verbose message (if VerboseFlag is enabled)
func Verbose(a ...interface{}) {
	if VerboseFlag || DebugFlag {
		fmt.Println(a...)
	}
}

// Info message
func Info(a ...interface{}) {
	if !QuietFlag {
		fmt.Println(a...)
	}
}

//Error message and exit
func Error(a ...interface{}) {
	if !QuietFlag {
		fmt.Println(a...)
	}
	os.Exit(1)
}
