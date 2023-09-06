package implementations

import (
	"fmt"
	"go_di/dynamic/interfaces"
)

type PrefixFormatterDefault struct {
}

func (l PrefixFormatterDefault) GetPrefix() string {
	return "> "
}

func ProvidePrefixFormatterDefault() interfaces.IPrefixFormatter {
	fmt.Println("ProvidePrefixFormatterDefault()")
	return new(PrefixFormatterDefault)
}
