package gumi

import "sync"

var wgpool = &sync.Pool{
	New: func() interface{} {
		return new(sync.WaitGroup)
	},
}

var mtxpool = &sync.Pool{
	New: func() interface{} {
		return new(sync.Mutex)
	},
}
