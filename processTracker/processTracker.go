package processTracker

import (
	"fmt"
	"time"
)

type ProcessTracker struct {
	processName string
	processState State
	maxRetries int
}

func init(){
	storage = Storage{}
}

func New(processName string) (ProcessTracker, error) {
	return NewWithRetries(processName, 3)
}

func NewWithRetries(processName string, maxRetries int) (ProcessTracker, error) {
	state, err := storage.loadOrCreate(processName)

	if err != nil {
		return ProcessTracker{}, err
	}

	tracker := ProcessTracker{ processName, state, maxRetries}

	return tracker, nil
}

func (t ProcessTracker) Step(stepName string, action func() error) error {
	step, err := t.enterStep(stepName)

	if err != nil {
		return err
	}

	if !step.stepState.succeeded {
		err = action()

		t.leaveStep(step, err)
	}

	return err
}

func (t ProcessTracker) enterStep(name string) (ProcessStep, error) {
	state := t.processState.getOrCreateStepState(name)

	if !state.succeeded {
		state.retries++

		if state.retries > t.maxRetries {
			return ProcessStep{}, fmt.Errorf("%s has tried to enter %s too many times", t.processName, name)
		}

		state.enteredAt = time.Now()
		t.processState.saveStepState(state)
	}

	return ProcessStep{name, state, t}, nil
}

func (t ProcessTracker) leaveStep(step ProcessStep, err error) {
	step.stepState.succeeded = err == nil
	step.stepState.error = err
	t.processState.saveStepState(step.stepState)
}

