package core

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/xeronith/diamante/logging"
)

type object struct {
	mutex     sync.RWMutex
	id        int64
	sortOrder float32
	keychain  map[uint64]*semaphore
}

func (object *object) SortOrder() float32 {
	return object.sortOrder
}

func (object *object) Id() int64 {
	return object.id
}

func (object *object) Lock(context uint64) {
	var key *semaphore

	func() {
		object.mutex.Lock()
		defer object.mutex.Unlock()

		if object.keychain == nil {
			object.keychain = make(map[uint64]*semaphore)
		}

		exists := false
		key, exists = object.keychain[context]
		if !exists {
			key = &semaphore{}
			object.keychain[context] = key
		}
	}()

	if !key.Acquire() {
		logging.GetDefaultLogger().Panic(fmt.Sprintf("failed_to_acquire_exclusive_lock: 0x%.8X %d %d", context, object.id, len(object.keychain)))
		panic("IDENTICAL_CONCURRENT_REQUESTS_NOT_ALLOWED")
	}
}

func (object *object) Unlock(context uint64) {
	var key *semaphore

	func() {
		object.mutex.RLock()
		defer object.mutex.RUnlock()

		if object.keychain == nil {
			logging.GetDefaultLogger().Panic("failed_to_release_exclusive_lock: empty_keychain")
			return
		}

		exists := false
		key, exists = object.keychain[context]
		if !exists {
			logging.GetDefaultLogger().Panic("failed_to_release_exclusive_lock: invalid_key")
			return
		}
	}()

	if !key.Release() {
		logging.GetDefaultLogger().Warning(fmt.Sprintf("exclusive_lock_already_released: 0x%.8X %d %d", context, object.id, len(object.keychain)))
	}
}

// -----------------------------------------------------------

type semaphore struct {
	locked int32
}

func (semaphore *semaphore) Acquire() bool {
	return atomic.CompareAndSwapInt32(&semaphore.locked, 0, 1)
}
func (semaphore *semaphore) Release() bool {
	return atomic.CompareAndSwapInt32(&semaphore.locked, 1, 0)
}

// -----------------------------------------------------------
