package logging

import "contracts"

const (
	StdIOLogger contracts.LoggerType = 0
)

func CreateLogger(componentType contracts.LoggerType) contracts.ILogger {
	switch componentType {
	case StdIOLogger:
		return New(true)
	default:
		panic("unknown_logger_type")
	}
}
