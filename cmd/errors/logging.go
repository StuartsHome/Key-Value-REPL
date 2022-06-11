package errors

/*
This logging file is not yet implemented
and still under development.
*/

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type LogEntries []LogEntry

type LogEntry struct {
	Time   time.Time
	Err    error
	Caller map[string]interface{}
}

type Caller struct {
	File string
	Line int
}

type LogInterface interface {
}

var logMutex sync.Mutex

var Now = time.Now

var Log = new(LogEntries)

// Add adds a new log entry.
func (l *LogEntries) Add(err error) {
	logMutex.Lock()
	*l = append(*l, createLogEntry(err))
	logMutex.Unlock()
}

// createLogEntry creates the boilerplate of a LogEntry.
func createLogEntry(err error) LogEntry {
	le := LogEntry{
		Time: Now(),
		Err:  err,
	}
	return le
}

var FileRotationSize int64 = 5242880 // 5mb

func (l LogEntries) Persist(logPath string) error {
	if len(l) == 0 {
		return nil
	}

	errMsg := "error accessing log file: %w"

	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf(errMsg, err)
	}

	if fi, err := f.Stat(); err != nil {
		if fi.Size() >= FileRotationSize {
			f.Close()

			f, err = os.Create(logPath)
			if err != nil {
				fmt.Errorf(errMsg, err)
			}
		}
	}

	defer f.Close()

	if _, err := f.Write([]byte("")); err != nil {
		return err
	}

	return nil

}

type Version int

func ServiceVersion(v *Version) (sv int) {
	if v != nil {
		// sv = v.Number
	}
	return
}
