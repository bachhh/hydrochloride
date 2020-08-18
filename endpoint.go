package hydrochloride

import "sync"

type Server struct {
	lock         *sync.RWMutex
	maxConns     uint32
	maxIdleConns uint32
	weight       uint32 // default to 1
	isDown       bool
	isStandby    bool
	// balancer will NOT route any request to standby server
	// unless ALL active server are down
}

type Servers []*Server
