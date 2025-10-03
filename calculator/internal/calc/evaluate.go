package calc

import "errors"

func Evaluate(nums []float64, op string) (float64, error) {
	a, b := nums[0], nums[1]

	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero is forbidden")
		}
		return a / b, nil
	default:
		return 0, errors.New("unsupported operator")
	}
}