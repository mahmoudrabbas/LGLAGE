package main

import "fmt"

type LogLevel int

const (
	Info LogLevel = iota + 0
	Trace
	Debug
	Warn
	Error
	Fetal
)

var levels []string = []string{"INFO", "TRACE", "Debug", "Warn", "Error", "Fetal"}

func (l LogLevel) String() string {
	if l > Fetal && l < Info {
		return "Unknown"
	}

	return levels[l]
}

func printLogLevel(l LogLevel) {
	fmt.Printf("Log Level: %v and value: %v \n", l, l.String())
}

func main() {

	printLogLevel(Info)
	printLogLevel(Trace)
	printLogLevel(Debug)
	printLogLevel(Warn)
	printLogLevel(Error)
	printLogLevel(Fetal)

}
