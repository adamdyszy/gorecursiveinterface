package implementation

import "log"

type Logger struct {
	keyVals  map[string]string
	disabled bool
	maxLevel int
	curLevel int
}

func NewLogger(curLevel, maxLevel int) *Logger {
	l := Logger{keyVals: make(map[string]string), maxLevel: maxLevel, curLevel: curLevel}
	if l.curLevel > l.maxLevel {
		l.disabled = true
	} else {
		l.disabled = false
	}
	return &l
}

func (l *Logger) Info(msg string) {
	if l.disabled {
		return
	}
	log.Println(l.keyVals, "LEVEL:", l.curLevel, "INFO:", msg)
}

func (l *Logger) Error(err error) {
	if l.disabled {
		return
	}
	log.Println(l.keyVals, "LEVEL:", l.curLevel, "ERROR:", err)
}

func (l *Logger) WithKeyValue(key, val string) *Logger {
	keyVals := make(map[string]string, len(l.keyVals)+1)
	for key, value := range l.keyVals {
		keyVals[key] = value
	}
	keyVals[key] = val
	return &Logger{keyVals: keyVals, disabled: l.disabled, maxLevel: l.maxLevel, curLevel: l.curLevel}
}

// adds levels to curLevel and disable if too high
func (l *Logger) V(level int) *Logger {
	newLogger := Logger{keyVals: make(map[string]string, len(l.keyVals)+1), maxLevel: l.maxLevel, curLevel: l.curLevel + level}
	for key, value := range l.keyVals {
		newLogger.keyVals[key] = value
	}
	if newLogger.curLevel > newLogger.maxLevel {
		newLogger.disabled = true
	} else {
		newLogger.disabled = false
	}
	return &newLogger
}
