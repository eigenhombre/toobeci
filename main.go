package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type stackElement interface {
	String() string
}

type stack struct {
	elements []stackElement
}

type intElement int32
type floatElement float64
type stringElement string

func (i intElement) String() string {
	return fmt.Sprintf("%d", i)
}

func (f floatElement) String() string {
	return fmt.Sprintf("%f", f)
}

func (s stringElement) String() string {
	return string(s)
}

var builtins = map[string]func(*stack) (*stack, error){
	"drop": func(s *stack) (*stack, error) {
		_, err := s.pop()
		return s, err
	},
	"dup": func(s *stack) (*stack, error) {
		e, err := s.pop()
		if err != nil {
			return s, err
		}
		s.push(e)
		s.push(e)
		return s, nil
	},
	".": func(s *stack) (*stack, error) {
		e, err := s.pop()
		if err != nil {
			return s, err
		}
		fmt.Println(e)
		return s, nil
	},
}

func (s *stack) String() string {
	ret := ""
	for i, e := range s.elements {
		ret += fmt.Sprintf("\t%d %v\n", i, e)
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

func main() {
	// print a preamble
	fmt.Print("Welcome to Toobeci.\n\n")
	rdr := bufio.NewReader(os.Stdin)
	// create a stack
	s := &stack{}
top:
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
		goto top
	}
	trimmed := strings.Trim(input, " \t\n")
	if trimmed != "" {
		// See if input is in dictionary
		if f, ok := builtins[trimmed]; ok {
			// If so, call it
			s, err = f(s)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print(s)
			goto top
		}
		s.push(stringElement(strings.Trim(trimmed, " \t\n")))
		fmt.Print(s)
	}
	goto top
}
