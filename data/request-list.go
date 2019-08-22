package data;

import (
	"../interfaces"
);

type RequestList struct {
	requests  []interfaces.Request;
	relations map[string][]int;
	in_work   int;
}

func (rl *RequestList) Init() {
	rl.requests = make([]interfaces.Request, 0);
	rl.relations = make(map[string][]int);
}

func (rl RequestList) GetInWork() int {
	return rl.in_work;
}

func (rl *RequestList) DecrementInWork() {
	rl.in_work = rl.in_work - 1;
}

func (rl *RequestList) SetRequest(request interfaces.Request) {
	rl.requests = append(rl.requests, request);
	relations := rl.relations[request.GetHash()];
	index := (len(rl.requests) - 1);
	rl.relations[request.GetHash()] = append(relations, index);
}

func (rl RequestList) GetRequest(key int) interfaces.Request {
	return rl.requests[ key ];
}

func (rl *RequestList) SetRequests(requests []interfaces.Request) {
	rl.requests = requests;
}

func (rl RequestList) GetRequests() []interfaces.Request {
	return rl.requests;
}

func (rl RequestList) GetRelation(key string) []int {
	return rl.relations[key];
}

func (rl RequestList) GetRelations() map[string][]int {
	return rl.relations;
}
