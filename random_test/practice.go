// Here is the test requirements:

// Our current system allows our customer support to invite new users for the App, but a new requirement was made to us to allow scheduling those invites.

// The problem we pose is to reuse the existing user invite interface and do a scheduler that follow the scheduler invite interface.

// We expect idiomatic go code, object-oriented as much as possible no external API like http or graphql is expected, nor main executable is required, only unit tests that prove it works the happy path.

// If you need storage for object simply use in-memory approach.

// For the purpose of the test do not consider that the scheduler is going to run on multiple boxes.

// Have fun!
// This is the skeleton you have to work from:

// package KultraTest

// import "time"

// type UsersInviter interface {
//    Invite(userEmail, firstName, lastName string) error
// }

//	type UsersScheduledInviter interface {
//	   Invite(userEmail, firstName, lastName string, when time.Time) error
//	}
//
// This is not the test set by me but I got from the company via email and I have copied it here AS-IS no modification
package randomtest

import (
	"sync"
	"time"
)

const (
	THREADCOUNT = 2
)

type UsersInviter interface {
	Invite(userEmail, firstName, lastName string) error
}

type UsersScheduledInviter interface {
	Invite(userEmail, firstName, lastName string, when time.Time) error
}

type Scheduler struct {
	ui UsersInviter

	mut sync.Mutex
}

type Request struct {
	userEmail string
	firstName string
	lastName  string
	when      time.Time
}

var Requests = []Request{}

func NewScheduleInviter(ui UsersInviter) *Scheduler {
	return &Scheduler{ui, sync.Mutex{}}
}
func (s *Scheduler) Invite(userEmail, firstName, lastName string, when time.Time) error {

	time.Sleep(time.Since(when))
	s.mut.Lock()
	Requests = append(Requests, Request{userEmail, firstName, lastName, when})
	s.mut.Unlock()
	return s.ui.Invite(userEmail, firstName, lastName)
}

// invite after time x
