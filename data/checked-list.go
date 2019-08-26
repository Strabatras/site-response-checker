package data;

import (
	"../interfaces"
	"fmt"
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

// возвращает true если ссылка проверялась ранее
//
func (cl *CheckedList) Observation(request interfaces.Request, line interfaces.Line, observation interfaces.Observation) bool {
	cl.mx.Lock();
	defer cl.mx.Unlock();
	// ссылка проверялась ранее
	if _request, ok := cl.data[ request.GetHash() ]; ok {
		// запрос был проверен
		if (_request.GetFinished()) {
			requests := line.GetRequestList().GetRequests();
			// все связи по проверенному запросу получают статус закончено
			// декремент счетчика 'в работе'
			for _, relation := range line.GetRequestList().GetRelation(_request.GetHash()) {
				requests[relation] = _request;
				line.GetRequestList().DecrementInWork();
			}
			line.GetRequestList().SetRequests(requests);
		} else {
			observation.Set(request.GetHash(), line);
		}
		return true;
	}
	cl.data[ request.GetHash() ] = request;
	return false;
}

func lineWriter(writer chan interfaces.Line, waitGroupWriter *sync.WaitGroup)  {
	defer waitGroupWriter.Done();
	for {
		line, more := <-writer
		if more {
			fmt.Println("LINE => ", line);
		} else {
			return;
		}
	}
}

func (cl *CheckedList) TakeOffObservation(request interfaces.Request, writer chan interfaces.Line, waitGroupWriter *sync.WaitGroup, observation interfaces.Observation) {
	lines := observation.Get(request.GetHash());
	observation.Forget(request.GetHash());
	for _, line := range lines {
		requests := line.GetRequestList().GetRequests();
		// все связи по проверенному запросу получают статус закончено
		// декремент счетчика 'в работе'
		for _, relation := range line.GetRequestList().GetRelation(request.GetHash()) {
			requests[relation] = request;
			line.GetRequestList().DecrementInWork();
		}
		line.GetRequestList().SetRequests(requests);
		if ( line.GetRequestList().GetInWork() == 0 ) {
			waitGroupWriter.Add(1);
			go lineWriter(writer, waitGroupWriter);
			writer <- line;
		}
	}
}