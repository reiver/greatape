package contracts

type ILogger interface {
	Info(...any)
	Debug(...any)
	Error(...any)
	Fatal(...any)
}
