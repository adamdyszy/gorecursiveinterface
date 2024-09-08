package abstract

type Logger[T any] interface {
	NoRecursionLogger
	WithKeyValue(key, val string) T
	V(level int) T
}

type NoRecursionLogger interface {
	Info(msg string)
	Error(err error)
}

var GlobalLogger WrapperLogger

type WrapperLogger struct {
	NoRecursionLogger
	V            func(level int) WrapperLogger
	WithKeyValue func(key, val string) WrapperLogger
}

func SetGlobalLogger[T NoRecursionLogger](logger Logger[T]) {
	NewWrapperLogger := make([]func(logger Logger[T]) WrapperLogger, 1)
	NewWrapperLogger[0] = func(logger Logger[T]) WrapperLogger {
		return WrapperLogger{
			NoRecursionLogger: logger,
			V: func(level int) WrapperLogger {
				return NewWrapperLogger[0](any(logger.V(level)).(Logger[T]))
			},
			WithKeyValue: func(key, val string) WrapperLogger {
				return NewWrapperLogger[0](any(logger.WithKeyValue(key, val)).(Logger[T]))
			},
		}
	}
	GlobalLogger = NewWrapperLogger[0](logger)
}
