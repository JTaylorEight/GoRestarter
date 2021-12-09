package processTracker

var storage Storage

type Storage struct {
}

func (s Storage) loadOrCreate(name string) (State, error) {
	// check actual storage for existing record and return
	return State{}, nil
}

func (s Storage) save(state State) error {
	// save record to actual storage
	return nil
}

