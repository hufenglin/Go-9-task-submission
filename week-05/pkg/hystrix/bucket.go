package hystrix

import (
	"sync"
	"time"
)

type Bucket struct {
	sync.RWMutex
	// Total request
	Total int
	// Failed times
	Failed    int
	Timestamp time.Time
}

func NewBucket() *Bucket {
	return &Bucket{
		Timestamp: time.Now(),
	}
}

// Record the result
func (b *Bucket) Record(result bool) {
	b.Lock()
	defer b.Unlock()
	if !result {
		b.Failed++
	}
	b.Total++
}
