package contracts

type (
	IConfig interface {
		Get(string) string
	}
)
