package common

import "sync"

type Mutex struct {
	Lock     *sync.Mutex
	LockLock sync.Mutex
}

type RWMutex struct {
	Lock *sync.RWMutex
}
