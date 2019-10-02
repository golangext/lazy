package lazy

import "sync"

type reference struct {
	initFn func() interface{}
	v      interface{}
	m      sync.Mutex
}

func NewReference(initFn func() interface{}) Reference {
	return &reference{initFn: initFn}
}

func (r *reference) Elem() interface{} {
	if r.initFn != nil {
		r.doInit()
	}
	return r.v
}

func (r *reference) doInit() {
	r.m.Lock()
	defer r.m.Unlock()
	if r.initFn != nil {
		r.v = r.initFn()
		r.initFn = nil
	}
}
