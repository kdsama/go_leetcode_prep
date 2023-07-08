package rate_limiting

import (
	"errors"
	"sync"
	"time"
)

// token bucket algorithm
// each request will be given a token
// after a fixed amount of time token will be added to the bucket again.
// different buckets will be present for different requests
// for our scenario , we will just receive the endpoint.

// our functionality will can look like this

// so if it is 1000 requests per second
// we can have a bucket of 1000 entries
// every second we fill it back up
// we  would need a configuration file here where number of request, time duration is present.
// if bucket is empty , return error

var (
	errBucketFull = errors.New("bucket is full, cannot take any more requests")
)

type tokenbucket struct {
	arr        []string
	interval   time.Duration
	bucketSize int
	identifier string
	ticker     *time.Ticker
	lock       sync.Mutex
}

func NewTokenBucket(interval, bucketsize int, identifier string) *tokenbucket {

	duration := time.Duration(interval) * time.Second
	ticker := time.NewTicker(duration)

	token := &tokenbucket{[]string{}, duration, bucketsize, identifier, ticker, sync.Mutex{}}
	go token.schedule(duration)
	return token
}

func (tb *tokenbucket) schedule(duration time.Duration) {
	for {
		select {
		case <-tb.ticker.C:
			{
				tb.update()
			}
		}
	}
}

func (tb *tokenbucket) update() {
	tb.arr = []string{}
}

func (tb *tokenbucket) add(token string) error {
	tb.lock.Lock()
	// lock on read as well required as we might have multiple 999th request requesting for the 1000th source.
	if len(tb.arr)+1 > tb.bucketSize {
		return errBucketFull
	}
	tb.arr = append(tb.arr, token)
	tb.lock.Unlock()
	return nil
}
