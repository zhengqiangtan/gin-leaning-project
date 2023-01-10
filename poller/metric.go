package poller

import "sync/atomic"

type metric struct {
	busyWorkers uint64
}

func newMetric() *metric {
	return &metric{}
}

func (m *metric) IncBusyWorker() uint64 {
	return atomic.AddUint64(&m.busyWorkers, 1)
}

func (m *metric) DecBusyWorker() uint64 {
	return atomic.AddUint64(&m.busyWorkers, ^uint64(0))
}

func (m *metric) BusyWorkers() uint64 {
	return atomic.LoadUint64(&m.busyWorkers)
}
