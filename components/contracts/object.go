package contracts

import . "fmt"

type IObject interface {
	Stringer
	// Id returns the unique identifier or 'Id' of this object instance.
	Id() int64
	SortOrder() float32
	// Validate checks whether the current instance contains any errors.
	Validate() error
	Lock(context uint64)
	Unlock(context uint64)
}
