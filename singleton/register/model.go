package register

import "sync"

var Global = newRegister()

func newRegister() *register {
	return &register{
		RWMutex: new(sync.RWMutex),
		store:   make(map[string]any),
	}
}

type register struct {
	*sync.RWMutex
	store map[string]any
}

func (r *register) Store(key string, value any) {
	r.Lock()
	defer r.Unlock()
	r.store[key] = value
}

func (r *register) Load(key string) (value any, ok bool) {
	r.RLock()
	defer r.RUnlock()
	value, ok = r.store[key]
	return
}
