package main

import (
	"testing"
)

func TestStuff(t *testing.T) {
	IN := func(s ...string) []string { return s }
	examples := []struct {
		input  []string
		output string
	}{
		{IN("1", "2", "+"), ""},
		// State is saved!
		{IN("."), "3"},
		// YAH
		//		{IN("1", "2", "+", "."), "3"},
	}
	i := newInterpreter()
	for _, e := range examples {
		for _, input := range e.input {
			output, err := i.handleInputLine(input)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if output != e.output {
				t.Errorf("expected '%v', got '%v'", e.output, output)
			}
		}
	}
}
