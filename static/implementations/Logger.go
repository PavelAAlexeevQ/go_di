package implementations

import (
	"fmt"
	"go_di/static/interfaces"
)

type LoggerToScreen struct {
}

func (l LoggerToScreen) Log(a ...any) {
	fmt.Println(a...)
}

func (l LoggerToScreen) Logf(format string, a ...any) {
	fmt.Printf(format, a...)
}

func ProvideLogger() interfaces.ILogger {
	fmt.Println("ProvideLogger()")
	return new(LoggerToScreen)
}
