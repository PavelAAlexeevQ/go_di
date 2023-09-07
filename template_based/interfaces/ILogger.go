package interfaces

type ILogger interface {
	Log(a ...any)
	Logf(format string, a ...any)
}
