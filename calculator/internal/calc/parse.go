package calc

import (
	"errors"
	"strconv"
)

func ParseArgs(args []string) ([]float64, string, error) {
	if len(args) != 3 {
		return nil, "", errors.New("need exactly 3 args: <num> <op> <num>")
	}

	// parse first number
	n1, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return nil, "", errors.New("invalid first number")
	}

	// validate operator
	op := args[1]
	if !(op == "+" || op == "-" || op == "*" || op == "/") {
		return nil, "", errors.New("unsupported operator")
	}

	// parse second number
	n2, err := strconv.ParseFloat(args[2], 64)
	if err != nil {
		return nil, "", errors.New("invalid second number")
	}

	return []float64{n1, n2}, op, nil
}