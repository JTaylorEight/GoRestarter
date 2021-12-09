package processTracker

import "time"

type StepState struct {
	name string
	succeeded bool
	error error
	createdAt time.Time
	enteredAt time.Time
	retries int
}


