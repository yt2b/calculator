package buffer

import "testing"

func TestRead(t *testing.T) {
	input := "12345"
	buf := NewBuffer(input)
	tests := []struct {
		expected byte
	}{
		{'1'},
		{'2'},
		{'3'},
		{'4'},
		{'5'},
		{'\x00'},
	}
	for _, tt := range tests {
		if buf.Ch != tt.expected {
			t.Fatalf("not = %q, got = %q", tt.expected, buf.Ch)
		}
		buf.Read()
	}
}
