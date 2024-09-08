package implementation

type NoLogger struct{}

func (l NoLogger) Info(msg string) {
}

func (l NoLogger) Error(err error) {
}

func (l NoLogger) WithKeyValue(key, val string) NoLogger {
	return l
}

func (l NoLogger) V(level int) NoLogger {
	return l
}
