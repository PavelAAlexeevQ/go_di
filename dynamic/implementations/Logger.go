package implementations

import (
	"fmt"
	"go_di/dynamic/interfaces"

	"go.uber.org/dig"
)

type LoggerToScreen struct {
	additionalParams AdditionalParams
}

type AdditionalParams struct {
	dig.In

	Prefix        interfaces.IPrefixFormatter `name:"customizedPrefix" optional:"true"`
	PrefixDefault interfaces.IPrefixFormatter `name:"defaultPrefix"`
}

func (l LoggerToScreen) GetPrefix() string {
	var prefix string
	if l.additionalParams.Prefix != nil {
		prefix = l.additionalParams.Prefix.GetPrefix()
	} else {
		prefix = l.additionalParams.PrefixDefault.GetPrefix()
	}
	return prefix
}

func (l LoggerToScreen) Log(a ...any) {
	prefix := l.GetPrefix()
	logString := prefix + fmt.Sprint(a...)
	fmt.Println(logString)
}

func (l LoggerToScreen) Logf(format string, a ...any) {
	prefix := l.GetPrefix()
	logString := prefix + fmt.Sprintf(format, a...)
	fmt.Println(logString)
}

func ProvideLogger(additionalParams AdditionalParams) interfaces.ILogger {
	fmt.Println("ProvideLogger()")
	logger := new(LoggerToScreen)
	logger.additionalParams = additionalParams
	return logger
}
