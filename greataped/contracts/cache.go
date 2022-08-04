package contracts

type ICache interface {
	Put(string, interface{})
	Get(string) interface{}
}
