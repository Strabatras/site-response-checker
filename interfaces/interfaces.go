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
	SetUrl(url string);
	GetUrl() string;
}

type RequestList interface {
	Init();
	SetRequest(key string, request Request);
	GetRequest(key string) Request;
	GetRequests() []Request;
}
