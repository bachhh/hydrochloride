package hydrochloride

import (
	"sync"
	"sync/atomic"
)

// Balancer an interface to pick the next endpoints to send request
// balancer should manage their own internal state, including which endpoints are failling
type Balancer interface {
	Next() (endpoint *Server, err error)
	List() (endpoints Servers, err error)
}

// wrrbalancer weighted round-robin balancer
type wrrbalancer struct {
	endpoints       Servers
	endpointWeights map[uint32]uint32
	curPos          uint32        // current position to return on the call to Next()
	locker          *sync.RWMutex // to lock endpoints updates
}

// WeightedRoundRobinBalancer return weighted round robin balancer
func WeightedRoundRobinBalancer(stuffs ...interface{}) Balancer {
	return &wrrbalancer{
		locker: &sync.RWMutex{},
	}
}

func (w *wrrbalancer) Next() (endpoint *Server, err error) {
	w.locker.RLock()
	defer w.locker.RUnlock()

	nextPos := atomic.AddUint32(&w.curPos, 1)

	endpoint := w.endpoints[nextPos]
	return
}

func (w *wrrbalancer) List() (endpoints Servers, err error) {
	return
}
