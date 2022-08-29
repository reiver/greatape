package logging

import . "contracts"

const (
	StdIOLogger LoggerType = iota
)

func CreateLogger(componentType LoggerType) ILogger {
	switch componentType {
	case StdIOLogger:
		return NewStdIOLogger(true)
	default:
		panic("unknown_logger_type")
	}
}
