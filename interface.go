package log

// Interface represents the API of both Logger and Entry.
type Interface interface {
	WithFields(Fielder) *Entry
	WithField(string, interface{}) *Entry
	WithError(error) *Entry
	WithoutPadding() *Entry
	Trace(string)
	Debug(string)
	Info(string)
	Warn(string)
	Error(string)
	Fatal(string)
	Tracef(string, ...interface{})
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
	ResetPadding()
	IncreasePadding()
	DecreasePadding()
}
