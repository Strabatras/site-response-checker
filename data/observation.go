package data;

import (
	"../interfaces"
	"sync"
)

type Observation struct {
	data map[string][]interfaces.Line;
	mx   sync.Mutex;
}

func (o *Observation) Init() {
	o.data = make(map[string][]interfaces.Line);
}

func (o *Observation) Set(key string, line interfaces.Line) {
	o.mx.Lock();
	defer o.mx.Unlock();
	data := o.data[ key ];
	o.data[ key ] = append(data, line);
}

func (o *Observation) Get(key string) []interfaces.Line {
	o.mx.Lock();
	defer o.mx.Unlock();
	return o.data[ key ];
}

func (o *Observation) Forget(key string) {
	o.mx.Lock();
	defer o.mx.Unlock();
	delete(o.data, key);
}