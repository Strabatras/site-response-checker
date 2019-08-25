package data;

import (
	"../interfaces"
	"sync"
)

type Writer struct {
	waitGroup  sync.WaitGroup;
	writerChan chan interfaces.Line;
}

func (w *Writer) GetWaitGroup() sync.WaitGroup {
	return w.waitGroup;
}
