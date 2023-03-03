package sync

import (
	"fmt"
	"sync"
)

var syncMap sync.Map
var mutex sync.Mutex
var rwMutex sync.RWMutex
var once sync.Once
var wg sync.WaitGroup
var syncPool sync.Pool

func Mutex() {
	mutex.Lock()
	defer mutex.Unlock()
}

func PrintOnce() {
	once.Do(func() {
		fmt.Println("")
	})
}
