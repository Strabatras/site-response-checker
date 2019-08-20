package data;

import (
	"../interfaces"
	"sync"
)

type CheckedList struct {
	data map[string]interfaces.Request;
	mx   sync.Mutex;
}

func (cl *CheckedList) Init() {
	cl.data = make(map[string]interfaces.Request);
}

func (cl *CheckedList) Set(request interfaces.Request) {
	cl.mx.Lock();
	defer cl.mx.Unlock();
	cl.data[ request.GetHash() ] = request;
}

func (cl *CheckedList) Get(key string) interfaces.Request {
	cl.mx.Lock();
	defer cl.mx.Unlock();
	return cl.data[ key ];
}

// возвращает true если ссылка ранее проверялась
func (cl *CheckedList) Observation(request interfaces.Request, line interfaces.Line, observation interfaces.Observation) bool {
	cl.mx.Lock();
	defer cl.mx.Unlock();
	if _, ok := cl.data[ request.GetHash() ]; ok {
		observation.Set(request.GetHash(), line);
		return true;
	}
	cl.data[ request.GetHash() ] = request;
	return false;
}
