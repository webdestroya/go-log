package main

import (
	"errors"

	"github.com/webdestroya/go-log"
)

func main() {
	log.SetLevel(log.TraceLevel)
	log.WithField("foo", "bar").Trace("trace")
	log.WithField("foo", "bar").Debug("debug")
	log.WithField("foo", "bar").Info("info")
	log.WithField("foo", "bar").Warn("warn")
	log.WithFields(log.Fields{
		"multiple": "fields",
		"yes":      true,
	}).Info("a longer line in this particular log")
	log.IncreasePadding()
	log.WithField("foo", "bar").Info("info with increased padding")
	log.IncreasePadding()
	log.WithoutPadding().WithField("foo", "bar").Info("info without padding")
	log.WithField("foo", "bar").Info("info with a more increased padding")
	log.ResetPadding()
	log.WithError(errors.New("some error")).Error("error")
	log.WithError(errors.New("some fatal error")).Fatal("fatal")
}
