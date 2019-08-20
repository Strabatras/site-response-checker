package interfaces;

type Line interface {
	SetId(id int);
	GetId() int;
	SetCells(cells []string);
	GetCells() []string;
	SetRequestList(requestList RequestList);
	GetRequestList() RequestList;
}

type Request interface {
	SetHash(hash string);
	GetHash() string;
	SetUrl(url string);
	GetUrl() string;
}

type RequestList interface {
	Init();
	SetRequest(request Request);
	GetRequests() []Request;
	GetRelations(key string) []int;
}

type CheckedList interface {
	Init();
	Observation( request Request, line Line, observation Observation ) bool;
}

type Observation interface {
	Init();
	Set( key string, line Line );
	Get( key string ) []Line;
}

type InProgress interface {
	SetCheckedList( checked CheckedList );
	GetCheckedList() CheckedList;
	SetObservation( observation Observation );
	GetObservation() Observation;
	ToObservation( url Request, line Line ) bool;
}