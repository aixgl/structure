package hmap

import (
	"errors"
	"sync"
)

// Independent safed map storage
type SafeMap struct {
	body  map[string]interface{}
	count uint64
	sync.RWMutex
}

// Create a new reference safemap.
func Make() *SafeMap {
	return &SafeMap{body: make(map[string]interface{})}
}

func (sm *SafeMap) Set(key string, value interface{}) error {
	sm.Lock()
	defer sm.Unlock()
	sm.body[key] = value
	sm.count++
	return nil
}

func (sm *SafeMap) Get(key string) (ret interface{}, err error) {
	sm.RLock()
	defer sm.RUnlock()
	ok := false
	ret, ok = sm.body[key]
	if !ok {
		err = errors.New("An error occurred attempting to get the value of safemap")
	}
	return
}

func (sm *SafeMap) Delete(key string) error {

	sm.Lock()
	defer sm.Unlock()
	_, ok := sm.body[key]

	if !ok {
		return errors.New("the key of map not exists;map[" + key + "]")
	}

	delete(sm.body, key)
	sm.count--
	return nil
}
