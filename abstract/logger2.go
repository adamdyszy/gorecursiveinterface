package abstract

type Logger2[T any] interface {
	NoRecursionLogger2
	WithKeyValue(key, val string) T
	V(level int) T
}

type NoRecursionLogger2 interface {
	Info(msg string)
	Error(err error)
}

type WrapperLogger2[T Logger2[T]] struct {
	logger T
}

func (w WrapperLogger2[T]) Info(msg string) {
	w.logger.Info(msg)
}

func (w WrapperLogger2[T]) Error(err error) {
	w.logger.Error(err)
}

func (w WrapperLogger2[T]) WithKeyValue(key, val string) WrapperLogger2[T] {
	w.logger = w.logger.WithKeyValue(key, val)
	return w
}

func (w WrapperLogger2[T]) V(level int) WrapperLogger2[T] {
	w.logger = w.logger.V(level)
	return w
}

// TODO: the problem is here, we cannot use SetGlobalLogger2
var GlobalLogger2 any

func SetGlobalLogger2[T Logger2[T]](logger T) {
	GlobalLogger2 = WrapperLogger2[T]{logger: logger}
}
