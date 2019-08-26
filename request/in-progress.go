package request;

import (
	"../interfaces"
	"sync"
)

type InProgress struct {
	checkedList interfaces.CheckedList;
	observation interfaces.Observation;
}

func (ip *InProgress) SetCheckedList(checked interfaces.CheckedList) {
	ip.checkedList = checked;
}

func (ip InProgress) GetCheckedList() interfaces.CheckedList {
	return ip.checkedList;
}

func (ip *InProgress) SetObservation(observation interfaces.Observation) {
	ip.observation = observation;
}

func (ip InProgress) GetObservation() interfaces.Observation {
	return ip.observation;
}

func (ip *InProgress) ToObservation(request interfaces.Request, line interfaces.Line) bool {
	return ip.GetCheckedList().Observation(request, line, ip.GetObservation());
}

func (ip *InProgress) FromObservation(request interfaces.Request, writer chan interfaces.Line, waitGroupWriter *sync.WaitGroup) {
	ip.GetCheckedList().TakeOffObservation(request, writer, waitGroupWriter, ip.GetObservation());
}
