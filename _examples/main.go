package main

import (
	"errors"

	"github.com/webdestroya/tieredlog"
)

func main() {
	tieredlog.SetLevel(tieredlog.TraceLevel)
	tieredlog.WithField("foo", "bar").Trace("trace")
	tieredlog.WithField("foo", "bar").Debug("debug")
	tieredlog.WithField("foo", "bar").Info("info")
	tieredlog.WithField("foo", "bar").Warn("warn")
	tieredlog.WithFields(tieredlog.Fields{
		"multiple": "fields",
		"yes":      true,
	}).Info("a longer line in this particular log")
	tieredlog.IncreasePadding()
	tieredlog.WithField("foo", "bar").Info("info with increased padding")
	tieredlog.IncreasePadding()
	tieredlog.WithoutPadding().WithField("foo", "bar").Info("info without padding")
	tieredlog.WithField("foo", "bar").Info("info with a more increased padding")
	tieredlog.ResetPadding()
	tieredlog.WithError(errors.New("some error")).Error("error")
	tieredlog.WithError(errors.New("some fatal error")).Fatal("fatal")
}
