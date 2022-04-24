package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"1", 1},
		{"1.23", 1.23},
		{"1 + 2", 3},
		{"53 - 45", 8},
		{"10 - 23 + 34", 21},
		{"2 * 3", 6},
		{"0.5 / 4", 0.125},
		{"120 * 5 / 20", 30},
		{"(34 - 9) * 4", 100},
		{"(15 * (76 - 34)) / ((56 - 36) * 10)", 3.15},
	}
	for _, tt := range tests {
		actual, _ := NewParser(tt.input).Parse()
		if actual != tt.expected {
			t.Fatalf("not = %f, got = %f", tt.expected, actual)
		}
	}
}
