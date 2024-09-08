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

//	func (w WrapperLogger) V(level int) WrapperLogger {
//		w.NoRecursionLogger = ????????.V(level).NoRecursionLogger
//		return w
//	}

func NewWrapperLogger[T any](logger Logger[T]) WrapperLogger {
	return WrapperLogger{
		NoRecursionLogger: logger,
		V: func(level int) WrapperLogger {
			return NewWrapperLogger(any(logger.V(level)).(Logger[T]))
		},
		WithKeyValue: func(key, val string) WrapperLogger {
			return NewWrapperLogger(any(logger.WithKeyValue(key, val)).(Logger[T]))
		},
	}
}

func SetGlobalLogger[T any](logger Logger[T]) {
	GlobalLogger = NewWrapperLogger(logger)
}
