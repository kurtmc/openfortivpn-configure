package main

import "testing"

func TestConfigure(t *testing.T) {
	expected := "Configure!"
	actual := Configure()

	if expected != actual {
		t.Errorf("expected \"%s\" but got \"%s\"\n", expected, actual)
	}

}
