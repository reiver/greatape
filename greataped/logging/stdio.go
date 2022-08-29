package logging

import (
	"contracts"
	"log"
)

type stdioLogger struct {
	verbose bool
}

func NewStdIOLogger(verbose bool) contracts.ILogger {
	return &stdioLogger{
		verbose: verbose,
	}
}

func (logger *stdioLogger) Info(args ...any) {
	log.Println(args...)
}

func (logger *stdioLogger) Debug(args ...any) {
	if !logger.verbose {
		return
	}

	log.Println(args...)
}

func (logger *stdioLogger) Error(args ...any) {
	log.Fatal(args...)
}

func (logger *stdioLogger) Fatal(args ...any) {
	log.Fatal(args...)
}
