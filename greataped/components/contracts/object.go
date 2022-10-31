package contracts

import . "fmt"

type IObject interface {
	Stringer
	// Returns the unique identifier or 'Id' of this object instance.
	Id() int64
	// Checks whether the current instance contains any errors.
	Validate() error
	Lock(context uint64)
	Unlock(context uint64)
}
