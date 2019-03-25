package dice

import "log"

type associativity int

const (
	left associativity = iota
	right
)

type operator struct {
	associativity
	precedence int
	fn         func(int, int) int
}

var operators = map[string]operator{
	"+": {left, 2, add},
	"-": {left, 2, sub},
}

func getOperator(t token) operator {
	op, ok := operators[t.val]
	if !ok {
		log.Fatalf("Unknown operator: %v", t.val)
	}
	return op
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}
