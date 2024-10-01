package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/rs/zerolog"
)

// newLogFile returns a pointer to a log file which will be located in
// log/log_{dateTime}.json
func newLogfile(deleteEmpty bool) (*os.File, error) {
	logDir := "logs"
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return nil, fmt.Errorf("couldn't create log directory: %w", err)
	}

	if deleteEmpty {
		if err := deleteEmptyLogs(logDir); err != nil {
			return nil, fmt.Errorf("error deleting empty logs: %w", err)
		}
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("log_%s.json", timestamp)

	// Construct the full path for the log file
	logPath := filepath.Join(logDir, filename)

	logFile, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, fmt.Errorf("couldn't open log file: %w", err)
	}

	return logFile, nil
}

// newLogger creates and returns a zerolog.Logger that writes to all provided writers. If
// no writers are provided, it will write only to the console (writing to os.Stderr). The
// logger includes timestamp and caller information in each log entry.
func newLogger(writers ...io.Writer) zerolog.Logger {
	validWriters := make([]io.Writer, 0, len(writers)+1)
	for _, writer := range writers {
		if writer != nil {
			validWriters = append(validWriters, writer)
		}
	}

	// default writer
	validWriters = append(validWriters, zerolog.ConsoleWriter{
		Out:        os.Stderr,
		NoColor:    false,
		TimeFormat: time.RFC3339,
	})

	multiWriter := zerolog.MultiLevelWriter(validWriters...)

	return zerolog.New(multiWriter).With().Timestamp().Caller().Logger()
}

// deleteEmptyLogs deletes any empty file with the .json extension inside logDir
func deleteEmptyLogs(logDir string) error {
	entries, err := os.ReadDir(logDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" {
			info, err := entry.Info()
			if err != nil {
				return err
			}
			if info.Size() == 0 {
				filePath := filepath.Join(logDir, entry.Name())
				if err := os.Remove(filePath); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
