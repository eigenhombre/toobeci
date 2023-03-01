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
type stringElement string

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

func (s stringElement) String() string {
	return string(s)
}

func (s stringElement) Equals(e stackElement) bool {
	switch t := e.(type) {
	case stringElement:
		return s == t
	}
	return false
}

func applyBinOp(s *stack, op func(a, b stackElement) stackElement) (*stack, string, error) {
	e1, err := s.pop()
	if err != nil {
		return s, "", err
	}
	e2, err := s.pop()
	if err != nil {
		return s, "", err
	}
	switch e1.(type) {
	case intElement:
		switch e2.(type) {
		case intElement:
			s.push(op(e1, e2))
			return s, "", nil
		}
	}
	return s, "", fmt.Errorf("type error")
}

const True = intElement(-1)
const False = intElement(0)

var builtins = map[string]func(*stack) (*stack, string, error){
	"+": func(s *stack) (*stack, string, error) {
		return applyBinOp(s, func(a, b stackElement) stackElement {
			return intElement(a.(intElement) + b.(intElement))
		})
	},
	"-": func(s *stack) (*stack, string, error) {
		return applyBinOp(s, func(a, b stackElement) stackElement {
			return intElement(a.(intElement) - b.(intElement))
		})
	},
	"*": func(s *stack) (*stack, string, error) {
		return applyBinOp(s, func(a, b stackElement) stackElement {
			return intElement(a.(intElement) * b.(intElement))
		})
	},
	"/": func(s *stack) (*stack, string, error) {
		return applyBinOp(s, func(a, b stackElement) stackElement {
			return intElement(a.(intElement) / b.(intElement))
		})
	},
	"and": func(s *stack) (*stack, string, error) {
		return applyBinOp(s, func(a, b stackElement) stackElement {
			return intElement(a.(intElement) & b.(intElement))
		})
	},
	"or": func(s *stack) (*stack, string, error) {
		return applyBinOp(s, func(a, b stackElement) stackElement {
			return intElement(a.(intElement) | b.(intElement))
		})
	},
	"=": func(s *stack) (*stack, string, error) {
		return applyBinOp(s, func(a, b stackElement) stackElement {
			if a.Equals(b) {
				return True
			}
			return False
		})
	},
	"not": func(s *stack) (*stack, string, error) {
		e, err := s.pop()
		if err != nil {
			return s, "", err
		}
		switch t := e.(type) {
		case intElement:
			if t == False {
				s.push(True)
			} else {
				s.push(False)
			}
			return s, "", nil
		}
		return s, "", fmt.Errorf("type error")
	},
	"drop": func(s *stack) (*stack, string, error) {
		_, err := s.pop()
		return s, "", err
	},
	"swap": func(s *stack) (*stack, string, error) {
		e1, err := s.pop()
		if err != nil {
			return s, "", err
		}
		e2, err := s.pop()
		if err != nil {
			return s, "", err
		}
		s.push(e1)
		s.push(e2)
		return s, "", nil
	},
	"rot": func(s *stack) (*stack, string, error) {
		e1, err := s.pop()
		if err != nil {
			return s, "", err
		}
		e2, err := s.pop()
		if err != nil {
			return s, "", err
		}
		e3, err := s.pop()
		if err != nil {
			return s, "", err
		}
		s.push(e2)
		s.push(e1)
		s.push(e3)
		return s, "", nil
	},
	"over": func(s *stack) (*stack, string, error) {
		e1, err := s.pop()
		if err != nil {
			return s, "", err
		}
		e2, err := s.pop()
		if err != nil {
			return s, "", err
		}
		s.push(e2)
		s.push(e1)
		s.push(e2)
		return s, "", nil
	},
	"dup": func(s *stack) (*stack, string, error) {
		e, err := s.pop()
		if err != nil {
			return s, "", err
		}
		s.push(e)
		s.push(e)
		return s, "", nil
	},
	".": func(s *stack) (*stack, string, error) {
		e, err := s.pop()
		if err != nil {
			return s, "", err
		}
		return s, e.String() + "\n", nil
	},
	"emit": func(s *stack) (*stack, string, error) {
		e, err := s.pop()
		if err != nil {
			return s, "", err
		}
		switch e := e.(type) {
		case intElement:
			return s, fmt.Sprintf("%c", e), nil
		default:
			return s, "", fmt.Errorf("type error")
		}
	},
	".s": func(s *stack) (*stack, string, error) {
		return s, s.String(), nil
	},
	"clr": func(s *stack) (*stack, string, error) {
		return &stack{}, "", nil
	},
}

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

type interpreter struct {
	s *stack
}

func newInterpreter() *interpreter {
	return &interpreter{
		s: &stack{},
	}
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
		// fmt.Println("handling word", word)
		// See if input is in dictionary
		if word == "" {
			continue
		}
		if f, ok := builtins[word]; ok {
			// If so, call it
			i.s, out, err = f(i.s)
			ret += out
			if err != nil {
				return ret, err
			}
			continue
		}
		// Not in the dictionary?  Try to parse it
		// as an int
		if intVal, err := strconv.Atoi(word); err == nil {
			i.s.push(intElement(intVal))
			continue
		}
		i.s.push(stringElement(word))
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
