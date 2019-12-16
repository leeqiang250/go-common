package common

import "sync"

type Mutex struct {
	Lock *sync.Mutex
}

type RWMutex struct {
	Lock *sync.RWMutex
}
