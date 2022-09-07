package log

import "fmt"

const (
	Red = uint8(91 + iota)
	Green
	Yellow
	Blue

	prefixInfo = "[INFO]"
	prefixWarn = "[WARN]"
	prefixErr  = "[ERR!]"
)

func GetColored(s string, colorCode uint8) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorCode, s)
}

func Info(s string, a ...interface{}) {
	fmt.Println(GetColored(prefixInfo, Blue), fmt.Sprintf(s, a...))
}

func Warn(s string, a ...interface{}) {
	fmt.Println(GetColored(prefixWarn, Yellow), fmt.Sprintf(s, a...))
}

func Err(s string, a ...interface{}) {
	fmt.Println(GetColored(prefixErr, Red), fmt.Sprintf(s, a...))
}
