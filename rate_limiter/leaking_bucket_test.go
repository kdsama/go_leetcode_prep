package rate_limiting

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestLeakingBucketAdd(t *testing.T) {
	t.Run("Test for err on exceeded requests", func(t *testing.T) {
		limits := [][]int{{1000000, 1}}
		for _, limit := range limits {

			w := NewLeakingBucket(limit[1], limit[0], "/users"+fmt.Sprint(limit[0]))

			for i := 0; i < limit[0]; i++ {
				w.add("some token")
			}

			got := w.add("some token")
			want := errBucketFull
			if want != got {
				t.Errorf("Wanted %v, but got %v for limit %v", want, got, limit[0])
			}
		}
	})

}

func BenchmarkLeakingBucketAdd(b *testing.B) {
	// test for multiple  endpoints
	// receiving multiple requests

	type req struct {
		limit    int
		endpoint string
		interval int
	}

	requests := []req{{endpoint: "/twitterfeed", limit: 1000, interval: 1}, {endpoint: "/users", limit: 200, interval: 1}, {endpoint: "/nonusers", limit: 2500, interval: 1}}
	ws := sync.WaitGroup{}
	ws.Add(len(requests))
	for _, re := range requests {
		total_req := 10000
		re := re
		go func(re req) {

			defer ws.Done()
			// re = re
			w := NewLeakingBucket(re.interval, re.limit, re.endpoint)

			for total_req >= 0 {
				num := rand.Intn(100)

				for i := 0; i < num; i++ {
					err := w.add("some token" + fmt.Sprint(i))
					if err != nil {
						b.Fatal("OOps didnot expect this here", err)
					}

				}

				time.Sleep(time.Duration(num) * time.Millisecond)
				total_req = total_req - num

			}
		}(re)
	}
	ws.Wait()

}
