package abstract

type WrapperLogger2[T Logger[T]] struct {
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

func SetGlobalLogger2[T Logger[T]](logger T) {
	GlobalLogger2 = WrapperLogger2[T]{logger: logger}
}
