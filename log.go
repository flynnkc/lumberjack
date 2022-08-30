package lumberjack

import (
	"errors"
	"fmt"
	"io"
	"log"
)

// If log level is greater than this something is wrong
const lumberjackMaxLogLevel int8 = 3

// MakeLogs generates a set of logs for different levels. Each level enables all
// sets of logs for each level lower than the selected level.
// 0 FATAL -- Abandon all hope, log the event, and put an end to the farce
// 1 ERROR -- Errors encountered during runtime
// 2 INFO -- General information
// 3 DEBUG -- An enhanced set of log collection to be used during debugging
func MakeLogs(level int, writer io.Writer) (*Woodpile, error) {
	if level < 0 || int8(level) > lumberjackMaxLogLevel {
		return nil, errors.New(fmt.Sprintf("Invalid log level setting %v", level))
	}

	w := &Woodpile{
		Fatal: log.New(writer, "FATAL: ", 7),
		level: int8(level),
	}
	if level > 0 {
		w.Error = log.New(writer, "ERROR: ", 7)
	}
	if level > 1 {
		w.Info = log.New(writer, "INFO: ", 7)
	}
	if level > 2 {
		w.Debug = log.New(writer, "DEBUG: ", 7)
	}

	return w, nil
}
