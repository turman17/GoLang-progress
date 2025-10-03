package calc

import (
	"math"
	"testing"
)

func almostEqual(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func TestEvaluate_OK(t *testing.T) {
	tests := []struct {
		name   string
		a, b   float64
		op     string
		want   float64
	}{
		{"add", 2, 3, "+", 5},
		{"sub", 7, 2, "-", 5},
		{"mul", 2.5, 2, "*", 5},
		{"div", 10, 2, "/", 5},
		{"div-float", 7, 2, "/", 3.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Evaluate([]float64{tt.a, tt.b}, tt.op)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !almostEqual(got, tt.want, 1e-9) {
				t.Fatalf("got=%v want=%v", got, tt.want)
			}
		})
	}
}

func TestEvaluate_Errors(t *testing.T) {
	// wrong length
	if _, err := Evaluate([]float64{1}, "+"); err == nil {
		t.Fatal("expected error for wrong nums length")
	}
	// division by zero
	if _, err := Evaluate([]float64{1, 0}, "/"); err == nil {
		t.Fatal("expected division by zero error")
	}
	// unsupported op
	if _, err := Evaluate([]float64{1, 2}, "%"); err == nil {
		t.Fatal("expected unsupported operator error")
	}
}