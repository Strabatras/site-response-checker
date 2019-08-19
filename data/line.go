package data;

import (
	"../interfaces"
);

type Line struct {
	id          int;
	cells       []string;
	requestList interfaces.RequestList;
}

func (l *Line) SetId(id int) {
	l.id = id;
}

func (l Line) GetId() int {
	return l.id;
}

func (l *Line) SetCells(cells []string) {
	l.cells = cells;
}

func (l Line) GetCells() []string {
	return l.cells;
}

func (l *Line) SetRequestList(requestList interfaces.RequestList) {
	l.requestList = requestList;
}
func (l Line) GetRequestList() interfaces.RequestList {
	return l.requestList;
}
