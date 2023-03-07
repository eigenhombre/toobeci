package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type stackElement interface {
	String() string
	Equals(stackElement) bool
}

type stack struct {
	elements []stackElement
}

type intElement int32

// type stringElement string

func (i intElement) String() string {
	return fmt.Sprintf("%d", i)
}

func (i intElement) Equals(e stackElement) bool {
	switch t := e.(type) {
	case intElement:
		return i == t
	}
	return false
}

// func (s stringElement) String() string {
// 	return string(s)
// }

// func (s stringElement) Equals(e stackElement) bool {
// 	switch t := e.(type) {
// 	case stringElement:
// 		return s == t
// 	}
// 	return false
// }

type dictionary map[string]func(*stack) (*stack, string, error)

type interpreter struct {
	s    *stack
	dict dictionary
}

func newInterpreter() *interpreter {
	return &interpreter{
		s:    &stack{},
		dict: initialDict, // builtin.go
	}
}

const True = intElement(-1)
const False = intElement(0)

func (s *stack) String() string {
	ret := ""
	// reverse order so top is at top:
	for i := len(s.elements) - 1; i >= 0; i-- {
		ret += fmt.Sprintf("\t%v\n", s.elements[i])
	}
	return ret
}

func (s *stack) push(e stackElement) {
	s.elements = append(s.elements, e)
}

func (s *stack) pop() (stackElement, error) {
	if len(s.elements) == 0 {
		return nil, fmt.Errorf("empty stack")
	}
	l := len(s.elements)
	e := s.elements[l-1]
	s.elements = s.elements[:l-1]
	return e, nil
}

func parse(input string) []string {
	halves := strings.Split(input, "\\")
	// discard comments:
	nonComment := halves[0]
	trimmed := strings.Trim(nonComment, " \t\n")
	return strings.Split(trimmed, " ")
}

func (i *interpreter) handleInputLine(input string) (string, error) {
	ret := ""
	words := parse(input)
	var err error
	var out string
	for _, word := range words {
		if word == "" {
			continue
		}
		// Is it in the dictionary?
		if f, ok := i.dict[word]; ok {
			i.s, out, err = f(i.s)
			ret += out
			if err != nil {
				return ret, err
			}
			continue
		}
		// Not in the dictionary?  Try to parse it
		// as an int (FIXME: add floats, strings, ...)
		if intVal, err := strconv.Atoi(word); err == nil {
			i.s.push(intElement(intVal))
			continue
		}
		// Out of options:
		return ret, fmt.Errorf("unknown word: %v", word)
	}
	return ret, nil
}

func main() {
	// print a preamble
	fmt.Print("Welcome to Toobeci.\n\n")
	rdr := bufio.NewReader(os.Stdin)
	i := newInterpreter()
	for {
		// print a prompt
		fmt.Print("> ")
		// read a line of input from stdin
		input, err := rdr.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("\nGoodbye.")
				return
			}
			fmt.Println(err)
			continue
		}
		out, err := i.handleInputLine(input)
		if err != nil {
			fmt.Println(err)
		}
		if out != "" {
			fmt.Println(out)
		}
	}
}
