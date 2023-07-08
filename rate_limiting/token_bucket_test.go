package rate_limiting

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("Test for err on exceeded requests", func(t *testing.T) {
		limits := [][]int{{1000000, 1}}
		for _, limit := range limits {

			w := NewTokenBucket(limit[1], limit[0], "/users"+fmt.Sprint(limit[0]))

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
