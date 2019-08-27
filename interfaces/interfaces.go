package interfaces;

import (
	"encoding/csv"
	"sync"
)

type Line interface {
	SetId(id int);
	GetId() int;
	SetCells(cells []string);
	GetCells() []string;
	SetRequestList(requestList RequestList);
	GetRequestList() RequestList;
}

type Request interface {
	SetFinished();
	GetFinished() bool;
	SetHash(hash string);
	GetHash() string;
	SetUrl(url string);
	GetUrl() string;
	SetStatusCode(statusCode int);
	GetStatusCode() int;
}

type RequestList interface {
	Init();
	GetInWork() int;
	DecrementInWork();
	IncrementInWork();
	SetRequest(request Request);
	GetRequest(key int) Request;
	SetRequests(requests []Request);
	GetRequests() []Request;
	GetRelation(key string) []int;
	GetRelations() map[string][]int;
}

type CheckedList interface {
	Init();
	Observation(request Request, line Line, observation Observation) bool;
	TakeOffObservation(request Request, lineToOut LineToOut, observation Observation);
}

type Observation interface {
	Init();
	Set(key string, line Line);
	Get(key string) []Line;
	Forget(key string);
}

type InProgress interface {
	SetCheckedList(checked CheckedList);
	GetCheckedList() CheckedList;
	SetObservation(observation Observation);
	GetObservation() Observation;
	ToObservation(request Request, line Line) bool;
	FromObservation(request Request, lineToOut LineToOut);
}

type LineToOut interface {
	SetWaitGroup(waitGroup *sync.WaitGroup);
	GetWaitGroup() *sync.WaitGroup;
	SetChanLine(chanLine chan Line);
	GetChanLine() chan Line;
	SetFileWriter(fileWriter FileWriter);
	GetFileWriter() FileWriter;
}

type FileWriter interface {
	SetWriter(writer *csv.Writer);
	WriteLine(line []string);
}
