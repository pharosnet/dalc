package dalc

type Log interface {
	Debugf(formatter string, args ...interface{})
}

func SetLog(v Log) {
	if v == nil {
		return
	}
	logger = v
}

var logger Log = nil

func hasLog() bool {
	return logger != nil
}
