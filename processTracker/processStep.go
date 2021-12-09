package processTracker

import "time"

type ProcessStep struct {
	stepName string
	stepState StepState
	tracker ProcessTracker
}

func (s State) getOrCreateStepState(name string) StepState {
	state, prs := s.stepStates[name]

	if !prs {
		state = StepState{name, false, nil, time.Now(), time.Time{}, 0}
	}

	return state
}

func (s State) saveStepState(state StepState) error {
	s.stepStates[state.name] = state
	return storage.save(s)
}

