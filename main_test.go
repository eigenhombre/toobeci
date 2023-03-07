package main

import (
	"os"
	"strings"
	"testing"
)

func TestStuff(t *testing.T) {
	var OK string = ""
	type example struct {
		input     string
		output    string
		canonical bool
		errStr    string
	}
	Case := func(input, output, errStr string) example {
		return example{input, output, false, errStr}
	}
	ECase := func(input, output, errStr string) example {
		return example{input, output, true, errStr}
	}
	S := func(ss string) string {
		rets := []string{}
		for _, s := range strings.Split(ss, " ") {
			rets = append(rets, "\t"+s)
		}
		return strings.Join(rets, "\n")
	}
	examples := []example{
		Case("1 drop", "", OK),
		ECase("\\ comments are ignored", "", OK),
		ECase("\\ . prints the 'top of the stack':", "", OK),
		ECase("1 .", "1", OK),
		ECase("\\ you can do math ...", "", OK),
		ECase("1 2 +", "", OK),
		ECase("\\ and then show the result:", "", OK),
		ECase(".", "3", OK),
		Case("1 .", "1", OK),
		Case("1 2 + .", "3", OK),
		Case("1 2 + 3 + .", "6", OK),
		ECase("10 10 / .", "1", OK),
		ECase("10 dup dup * * .", "1000", OK),
		ECase("2 3 drop .", "2", OK),
		ECase("42 emit", "*", OK),
		ECase("\\ The `emit` operator emits unicode characters:", "", OK),
		ECase("27700 emit", "水", OK),
		ECase("clr      \\ clears the stack", "", OK),
		ECase(".s       \\ shows the stack", "", OK),
		ECase("1 2 3 .s", S("3 2 1"), OK),
		ECase("swap .s  \\ swap top two items", S("2 3 1"), OK),
		ECase("rot .s   \\ rotate items", S("1 2 3"), OK),
		ECase("over .s  \\ copy & promote 2nd item", S("2 1 2 3"), OK),
		ECase("\\ Some boolean logic:", "", OK),
		ECase("1 1 and .", "1", OK),
		ECase("1 0 and .", "0", OK),
		Case("0 0 and .", "0", OK),
		Case("0 1 and .", "0", OK),
		ECase("1 1 or .", "1", OK),
		ECase("1 0 or .", "1", OK),
		Case("0 0 or .", "0", OK),
		Case("0 1 or .", "1", OK),
		ECase("\\ Default 'true' value is -1 (0b1111...):", "", OK),
		ECase("1 1 = .", "-1", OK),
		ECase("1 0 = .", "0", OK),
		ECase("3 not .", "0", OK),
		ECase("0 not .", "-1", OK),
		// ECase(": three 3 ;", "", OK),
		// ECase("three .", "3", OK),
		ECase("dakine", "", "unknown word"),
	}
	i := newInterpreter()
	// Save ECases to a file examples.fs
	// open file
	outfile := "examples.fs"
	f, err := os.Create(outfile)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	defer f.Close()
	f.WriteString("```\n" + `$ go build .
$ go install
$ toobeci
Welcome to toobeci

`)
	for _, e := range examples {
		output := ""
		out, err := i.handleInputLine(e.input)
		if err != nil && e.errStr == "" {
			t.Errorf("unexpected error: %v", err)
		} else if err == nil && e.errStr != "" {
			t.Errorf("expected error: %v", e.errStr)
		} else if err != nil && e.errStr != "" && !strings.Contains(err.Error(), e.errStr) {
			t.Errorf("expected error: %v, got %v", e.errStr, err.Error())
		} else {
			output += out
		}
		output = strings.Trim(output, "\n")
		if output != e.output && e.output != "IGNORE" {
			t.Errorf("expected '%v', got '%v'", e.output, output)
		}
		if e.canonical {
			f.WriteString("> " + e.input + "\n")
			if e.output != "" {
				f.WriteString(output + "\n")
			}
			if err != nil && err.Error() != "" {
				f.WriteString(err.Error() + "\n")
			}
		}
	}
	f.WriteString(`> ^D
Goodbye.
` + "```\n")
}
