package cache

import "sync"

var CachePWD = struct {
	sync.RWMutex
	data map[string]string
}{
	data: make(map[string]string),
}

func Set(key, value string) {
	CachePWD.Lock()
	defer CachePWD.Unlock()
	CachePWD.data[key] = value
}

func Get(key string) (string, bool) {
	CachePWD.RLock()
	defer CachePWD.RUnlock() // Fixed the defer call, it was unlocking something that wasn't locked
	val, ok := CachePWD.data[key]
	return val, ok
}
