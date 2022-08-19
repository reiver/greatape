package contracts

type (
	LoggerType int

	ILogger interface {
		Info(...any)
		Debug(...any)
		Error(...any)
		Fatal(...any)
	}
)
