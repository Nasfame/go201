package main_test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Run the test suite and get the exit code
	exitCode := m.Run()

	// Perform any additional cleanup or teardown tasks here

	// Exit the program with the test suite's exit code
	os.Exit(exitCode)
}

func TestAddition(t *testing.T) {
	result := 2 + 2
	expected := 4

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSubtraction(t *testing.T) {
	result := 5 - 3
	expected := 2

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}


