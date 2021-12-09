package main

import (
	"fmt"
	"github.com/JohnLTaylor/GoRestarter/database"
	"github.com/JohnLTaylor/GoRestarter/processTracker"
)

func main() {
	someOuterProcess("ABC123")
}

func someOuterProcess(id string) error {
	record := database.GetRecord(id)

	processName := fmt.Sprintf("someOuterProcess-%s", id)
	tracker, err := processTracker.New(processName)

	if err != nil {
		return err
	}

	err = tracker.Step("someInnerProcess", func() error {
		return someInnerProcess(record)
	})

	if err != nil {
		return err
	}

	err = tracker.Step("recordRefund", func() error {
		return database.SaveRecord(record)
	})

	if err != nil {
		return err
	}

	return nil
}

func someInnerProcess(record database.Record) error {
	return nil
}
