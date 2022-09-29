package utility

import (
	"fmt"
	"math/rand"
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

func GenerateConfirmationCode() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", 100000+rand.Intn(899999))
}
