package data;

import (
	"../interfaces"
	"sync"
)

type LineWriter struct {
	waitGroup sync.WaitGroup;
	chanLine chan interfaces.Line;
}

func (lw *LineWriter) SetWaitGroup( waitGroup *sync.WaitGroup )  {
	lw.waitGroup = *waitGroup;
}

func (lw *LineWriter) GetWaitGroup() *sync.WaitGroup {
	return &lw.waitGroup;
}

func (lw *LineWriter) SetChanLine(chanLine chan interfaces.Line)  {
	lw.chanLine = chanLine;
}
func (lw LineWriter) GetChanLine() chan interfaces.Line {
	return lw.chanLine;
}