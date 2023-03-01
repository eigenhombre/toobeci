package main

import (
	"os"
	"strings"
	"testing"
)

func TestStuff(t *testing.T) {
	type example struct {
		input     []string
		output    string
		canonical bool
	}
	Case := func(params ...string) example {
		output := params[len(params)-1]
		input := params[:len(params)-1]
		return example{input, output, false}
	}
	ECase := func(params ...string) example {
		output := params[len(params)-1]
		input := params[:len(params)-1]
		return example{input, output, true}
	}
	S := func(ss string) string {
		rets := []string{}
		for _, s := range strings.Split(ss, " ") {
			rets = append(rets, "\t"+s)
		}
		return strings.Join(rets, "\n")
	}
	examples := []example{
		Case("1", "drop", ""),
		ECase("\\ comments are ignored", ""),
		ECase("\\ . prints the 'top of the stack':", ""),
		ECase("1 .", "1"),
		ECase("\\ you can do math ...", ""),
		ECase("1", "2", "+", ""),
		ECase("\\ and then show the result:", ""),
		ECase(".", "3"),
		Case("1", ".", "1"),
		Case("1", "2", "+", ".", "3"),
		Case("1", "2", "+", "3", "+", ".", "6"),
		ECase("10 10 / .", "1"),
		ECase("10 dup dup * * .", "1000"),
		ECase("2 3 drop .", "2"),
		ECase("42 emit", "*"),
		ECase("\\ Unrecognized symbols are just strings for now.", ""),
		ECase("\\ But the `emit` operator emits unicode characters:", ""),
		ECase("Unicode is fun 27700 emit", "水"),
		ECase(".", "fun"),
		ECase("clr      \\ clears the stack", ""),
		ECase(".s       \\ shows the stack", ""),
		ECase("1 2 3 .s", S("3 2 1")),
		ECase("swap .s  \\ swap top two items", S("2 3 1")),
		ECase("rot .s   \\ rotate items", S("1 2 3")),
		ECase("over .s  \\ copy & promote 2nd item", S("2 1 2 3")),
		ECase("\\ Some boolean logic:", ""),
		ECase("1 1 and .", "1"),
		ECase("1 0 and .", "0"),
		Case("0 0 and .", "0"),
		Case("0 1 and .", "0"),
		ECase("1 1 or .", "1"),
		ECase("1 0 or .", "1"),
		Case("0 0 or .", "0"),
		Case("0 1 or .", "1"),
		ECase("\\ Default 'true' value is -1 (0b1111...):", ""),
		ECase("1 1 = .", "-1"),
		ECase("1 0 = .", "0"),
		ECase("3 not .", "0"),
		ECase("0 not .", "-1"),
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
		for _, input := range e.input {
			out, err := i.handleInputLine(input)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			output += out
		}
		output = strings.Trim(output, "\n")
		if output != e.output && e.output != "IGNORE" {
			t.Errorf("expected '%v', got '%v'", e.output, output)
		}
		if e.canonical {
			f.WriteString("> " + strings.Join(e.input, " ") + "\n")
			if e.output != "" {
				f.WriteString(output + "\n")
			}
		}
	}
	f.WriteString(`> ^D
Goodbye.
` + "```\n")
}
