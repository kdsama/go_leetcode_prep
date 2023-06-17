package randomtest

import (
	"sync"
	"testing"
	"time"
)

func TestMain(t *testing.T) {

	// Test for Errors

	// Test For Good Cases
}

type mockUserInviter struct {
	err error
}

func (mui *mockUserInviter) Invite(userEmail, firstName, lastName string) error {
	return mui.err
}

func TestScheduler(t *testing.T) {

	var wg sync.WaitGroup
	sch := NewScheduleInviter(&mockUserInviter{})
	ti := time.Now()
	wg.Add(1)
	fiveSeconds := 5 * time.Second
	go func() {
		sch.Invite("kd@gmail.com", "kshitij", "dhingra", time.Now().Add(fiveSeconds))
		wg.Done()
	}()
	wg.Wait()
	elapsed := time.Since(ti)
	delta := 100 * time.Millisecond
	if elapsed > delta+(5*time.Second) || elapsed < fiveSeconds-delta {
		t.Errorf("time expected is wrong , it took  %v", elapsed)
	}
}
