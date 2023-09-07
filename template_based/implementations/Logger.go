package implementations

import (
	"fmt"
	"go_di/do/interfaces"

	"github.com/samber/do"
)

type LoggerToScreen struct {
}

func (l LoggerToScreen) Log(a ...any) {
	fmt.Println(a...)
}

func (l LoggerToScreen) Logf(format string, a ...any) {
	fmt.Printf(format, a...)
}

func ProvideLogger(i *do.Injector) (interfaces.ILogger, error) {
	fmt.Println("ProvideLogger()")
	return LoggerToScreen{}, nil
}
