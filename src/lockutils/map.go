package lockutils

import "sync"

type lmap struct {
	mx sync.Mutex
	elements map[string]*sync.RWMutex
}

func NewMap() *lmap {
	_map := &lmap{}
	_map.elements = make(map[string]*sync.RWMutex)

	return _map
}

func (lm *lmap) getOrCreate(path string) *sync.RWMutex {
	lm.mx.Lock()
	if val, ok := lm.elements[path]; ok {
		lm.mx.Unlock()
		return val
	}

	lm.elements[path] = &sync.RWMutex{}
	lm.mx.Unlock()

	return lm.elements[path]
}

func (lm *lmap) RLock(path string) {
	lock := lm.getOrCreate(path)
	lock.RLock()
}

func (lm *lmap) RUnlock(path string) {
	lock := lm.getOrCreate(path)
	lock.RUnlock()
}

func (lm *lmap) Lock(path string) {
	lock := lm.getOrCreate(path)
	lock.Lock()
}

func (lm *lmap) Unlock(path string) {
	lock := lm.getOrCreate(path)
	lock.Unlock()
}