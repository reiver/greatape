package contracts

type (
	CacheType int

	ICache interface {
		Put(string, interface{})
		Get(string) interface{}
	}
)
