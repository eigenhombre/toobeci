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

func applyBinOp(s *stack, op func(a, b stackElement) stackElement) (*stack, error) {
	e1, err := s.pop()
	if err != nil {
		return s, err
	}
	e2, err := s.pop()
	if err != nil {
		return s, err
	}
	switch e1.(type) {
	case intElement:
		switch e2.(type) {
		case intElement:
			s.push(op(e1, e2))
			return s, nil
		}
	}
	return s, fmt.Errorf("type error")
}

var builtins = map[string]func(*stack) (*stack, error){
	"+": func(s *stack) (*stack, error) {
		return applyBinOp(s, func(a, b stackElement) stackElement {
			return intElement(a.(intElement) + b.(intElement))
		})
	},
	"-": func(s *stack) (*stack, error) {
		return applyBinOp(s, func(a, b stackElement) stackElement {
			return intElement(a.(intElement) - b.(intElement))
		})
	},
	"*": func(s *stack) (*stack, error) {
		return applyBinOp(s, func(a, b stackElement) stackElement {
			return intElement(a.(intElement) * b.(intElement))
		})
	},
	"/": func(s *stack) (*stack, error) {
		return applyBinOp(s, func(a, b stackElement) stackElement {
			return intElement(a.(intElement) / b.(intElement))
		})
	},
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
	"emit": func(s *stack) (*stack, error) {
		e, err := s.pop()
		if err != nil {
			return s, err
		}
		switch e := e.(type) {
		case intElement:
			fmt.Printf("%c", e)
			return s, nil
		default:
			return s, fmt.Errorf("type error")
		}
	},
	".s": func(s *stack) (*stack, error) {
		fmt.Print(s)
		return s, nil
	},
}

func (s *stack) String() string {
	ret := ""
	for i, e := range s.elements {
		ret += fmt.Sprintf("\t%d %v (%T)\n", i, e, e)
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
	fmt.Print("Welcome to Toobeci.\n")

	rdr := bufio.NewReader(os.Stdin)
	// create a stack
	s := &stack{}
outer:
	for {
		// print a prompt
		fmt.Print("\n> ")
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
		trimmed := strings.Trim(input, " \t\n")
		// Split on whitespace
		words := strings.Split(trimmed, " ")
		// fmt.Println("words:", words)
		for _, word := range words {
			// fmt.Println("handling word", word)
			// See if input is in dictionary
			if f, ok := builtins[word]; ok {
				// If so, call it
				s, err = f(s)
				if err != nil {
					fmt.Println(err)
					continue outer
				}
				continue
			}
			// Not in the dictionary?  Try to parse it
			// as an int
			if i, err := strconv.Atoi(word); err == nil {
				s.push(intElement(i))
				continue
			}
			s.push(stringElement(word))
		}
		fmt.Print(s)
	}
}
