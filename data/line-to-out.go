package data;

import (
	"../interfaces"
	"sync"
)

type LineToOut struct {
	waitGroup sync.WaitGroup;
	chanLine chan interfaces.Line;
}

func (lto *LineToOut) SetWaitGroup( waitGroup *sync.WaitGroup )  {
	lto.waitGroup = *waitGroup;
}

func (lto *LineToOut) GetWaitGroup() *sync.WaitGroup {
	return &lto.waitGroup;
}

func (lto *LineToOut) SetChanLine(chanLine chan interfaces.Line)  {
	lto.chanLine = chanLine;
}
func (lto LineToOut) GetChanLine() chan interfaces.Line {
	return lto.chanLine;
}