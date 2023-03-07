package main

import "fmt"

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

var initialDict = dictionary{
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
