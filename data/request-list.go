package data;

import (
	"../interfaces"
);

type RequestList struct {
	requests  []interfaces.Request;
	relations map[string][]int;
}

func (rl *RequestList)Init() {
	rl.requests = make([]interfaces.Request, 0);
	rl.relations = make(map[string][]int);
}

func (rl *RequestList) SetRequest(key string, request interfaces.Request) {
	rl.requests = append(rl.requests, request);
	relations := rl.relations[key];
	index := (len(rl.requests) -1);
	rl.relations[key] = append(relations, index );
}

func (rl RequestList) GetRequest(key string) interfaces.Request {
	return rl.requests[rl.relations[key][0]];
}

func (rl RequestList) GetRequests() []interfaces.Request {
	return rl.requests;
}
