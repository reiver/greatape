package utility

import (
	"sync"
	"time"
)

var (
	semaphore = &sync.RWMutex{}
	lastId    int64
)

func UniqueId() int64 {
	semaphore.Lock()
	defer semaphore.Unlock()

	var id int64
	for {
		id = time.Now().UnixNano() / 1000
		if id != lastId {
			break
		}

		time.Sleep(time.Microsecond)
	}

	lastId = id
	return id
}
