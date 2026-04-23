package main

import "fmt"

type LogLevel int

const (
	logWarn LogLevel = iota + 5
	logErr
	logDebug
	logInfo
	logFatal
)

const (
	Sunday int = iota + 2
	Monday
	Tuesday
)

func main() {

	fmt.Println(logWarn)
	fmt.Println(logErr)
	fmt.Println(logDebug)

}
