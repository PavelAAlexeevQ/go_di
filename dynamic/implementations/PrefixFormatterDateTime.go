package implementations

import (
	"fmt"
	"go_di/dynamic/interfaces"
	"time"
)

type PrefixFormatterDateTime struct {
}

func (l PrefixFormatterDateTime) GetPrefix() string {
	return fmt.Sprintf("%v: ", time.Now())
}

func ProvidePrefixFormatterDateTime() interfaces.IPrefixFormatter {
	fmt.Println("ProvidePrefixFormatterDateTime()")
	return new(PrefixFormatterDateTime)
}
